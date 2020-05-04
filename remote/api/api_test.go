package api

import (
	"testing"
)

func TestGetNodeConfig(t *testing.T) {
	v2node, err := GetNodeConfig()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v2node)
}

func TestGetUser(t *testing.T) {
	u, err := GetUser("xxxxxxxxx")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)
}

func TestPushStats(t *testing.T) {
	stats := [][]interface{}{
		{"t@tt", 0, 1},
	}
	err := PushStats(stats)
	if err != nil {
		t.Error(err)
		return
	}
}
