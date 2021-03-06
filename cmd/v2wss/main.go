package main

//go:generate errorgen

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/remote-v2ray/agent/remote"
	"v2ray.com/core"

	_ "github.com/remote-v2ray/agent/cmd/distro/v2wss"
)

var (
	version = flag.Bool("version", false, "Show current version of V2Ray.")
	test    = flag.Bool("test", false, "Test config file only, without launching V2Ray server.")
)

func startV2Ray() (core.Server, error) {
	configFile := "stdin:"
	pbconfig, err := remote.GetNodeConfig()
	if err != nil || len(pbconfig) == 0 {
		return nil, newError("failed to load config: ", configFile).Base(err)
	}
	configInput := bytes.NewReader(pbconfig)

	config, err := core.LoadConfig("protobuf", configFile, configInput)
	if err != nil {
		return nil, newError("failed to read config file: ", configFile).Base(err)
	}

	server, err := core.New(config)
	if err != nil {
		return nil, newError("failed to create server").Base(err)
	}
	if err := remote.Init(server); err != nil {
		return nil, newError("failed to init remote server").Base(err)
	}
	return server, nil
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}

func main() {
	flag.Parse()

	printVersion()

	if *version {
		return
	}

	server, err := startV2Ray()
	if err != nil {
		fmt.Println(err.Error())
		// Configuration error. Exit with a special value to prevent systemd from restarting.
		os.Exit(23)
	}

	if *test {
		fmt.Println("Configuration OK.")
		os.Exit(0)
	}

	if err := server.Start(); err != nil {
		fmt.Println("Failed to start", err)
		os.Exit(-1)
	}
	defer server.Close()

	// Explicitly triggering GC to remove garbage from config loading.
	runtime.GC()

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-osSignals
	}
}
