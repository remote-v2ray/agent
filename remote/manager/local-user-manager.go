package manager

import (
	"context"
	"sync"
	"time"

	"github.com/remote-v2ray/agent/remote/api"
	"v2ray.com/core/common/protocol"
	feature_stats "v2ray.com/core/features/stats"
	"v2ray.com/core/proxy"
	"v2ray.com/core/proxy/vmess"
)

type LocalUserManager struct {
	users  map[string]int64
	access sync.RWMutex
	// 用户过期时间, 超过这个时间的用户将会被删除
	// 如果用户继续保持流量的话, 将保留用户但会上报并重置流量
	ExpireDuration int64
	// 定时任务执行间隔
	ScheduledTaskInterval time.Duration
	MemoryUserManager     proxy.UserManager
	StatsManager          feature_stats.Manager
}

func NewLocalUserManager() *LocalUserManager {
	return &LocalUserManager{
		users:                 map[string]int64{},
		access:                sync.RWMutex{},
		ExpireDuration:        int64(5 * time.Minute / time.Second),
		ScheduledTaskInterval: time.Minute,
	}
}

func toLocalMemoryUser(u api.V2rayNodeUser) (user *protocol.MemoryUser, err error) {
	va := &vmess.Account{
		Id:      u.ID,
		AlterId: uint32(u.AlterID),
	}
	ma, err := va.AsAccount()
	if err != nil {
		return nil, err
	}
	user = &protocol.MemoryUser{
		Email:   u.Email,
		Level:   uint32(0),
		Account: ma,
	}
	return
}

func (um *LocalUserManager) AddUser(user api.V2rayNodeUser) (err error) {
	um.access.Lock()
	defer um.access.Unlock()
	if _, found := um.users[user.Email]; found {
		um.users[user.Email] = time.Now().Unix()
		return
	}
	newError("AddUser:", user.Email).AtDebug().WriteToLog()
	mu, err := toLocalMemoryUser(user)
	if err != nil {
		return
	}
	if isExist := um.MemoryUserManager.AddUser(context.Background(), mu); isExist != nil {
		// 用户已存在再添加的话会报错, 无视这个错误
	}
	um.users[user.Email] = time.Now().Unix()
	return
}

// 流量统计
type stat struct {
	uplink   int64
	downlink int64
}

func (um *LocalUserManager) RemoveUser(email string) stat {
	um.access.Lock()
	defer um.access.Unlock()

	// get stats counter
	up := "user>>>" + email + ">>>traffic>>>uplink"
	down := "user>>>" + email + ">>>traffic>>>downlink"
	upC := um.StatsManager.GetCounter(up)
	downC := um.StatsManager.GetCounter(down)
	uplink := upC.Value()
	downlink := downC.Value()

	// 如果仍然有流量变动的话, 就只更新流量, 不做删除.
	// 首次的流量肯定是不为 0 的, 所以最少需要等待两个检查周期才能删除
	if downlink != 0 {
		newError("重置用户流量:", email).AtDebug().WriteToLog()
		upC.Set(0)
		downC.Set(0)
		return stat{uplink, downlink}
	}
	newError("回收不活跃用户:", email).AtDebug().WriteToLog()

	// clear stats counter
	um.StatsManager.UnregisterCounter(up)
	um.StatsManager.UnregisterCounter(down)
	// clear memory user
	um.MemoryUserManager.RemoveUser(context.Background(), email)
	delete(um.users, email)
	return stat{uplink, downlink}
}

func (um *LocalUserManager) ClearInactiveUser() {
	now := time.Now().Unix()
	stats := [][]interface{}{}
	for email, lastActiveTime := range um.users {
		inactiveDuration := now - lastActiveTime
		if inactiveDuration > um.ExpireDuration {
			if t := um.RemoveUser(email); t.downlink != 0 || t.uplink != 0 {
				stat := []interface{}{email, t.uplink, t.downlink}
				stats = append(stats, stat)
			}
		}
	}
	if len(stats) > 0 {
		newError("发送流量统计").AtDebug().WriteToLog()
		go api.PushStats(stats)
	}
}

func (um *LocalUserManager) StartScheduledTask() {
	ticker := time.NewTicker(um.ScheduledTaskInterval)
	for {
		select {
		case <-ticker.C:
			go um.ClearInactiveUser()
		}
	}
}
