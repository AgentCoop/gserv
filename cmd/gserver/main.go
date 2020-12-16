package main

import (
	"fmt"
	g "github.com/AgentCoop/gserv/internal/gserver"
	m "github.com/AgentCoop/gserv/internal/middleware"
	pm "github.com/AgentCoop/gserv/internal/perfmetrics"
	"github.com/jessevdk/go-flags"
	"os"
	"syscall"
	"os/signal"
)

func showVersion() {
	fmt.Printf("version\n")
	os.Exit(0)
}

func mainEventLoop() {
	osSigs := make(chan os.Signal, 1)
	signal.Notify(osSigs, syscall.SIGINT, syscall.SIGTERM)

	runtimeInfo := g.NewRuntimeInfo()
	m.OnInit()

	perfMetrics := pm.NewPerfMetrics()
	go perfMetrics.Run(runtimeInfo)

	shutdown := func() {
		perfMetrics.Shutdown()
	}

	for {
		select {
		case <-osSigs:
			fmt.Printf("Got system exit singal\n")
			shutdown()
			return
		case <-runtimeInfo.ExitChannel:
			fmt.Printf("Got exit singal\n")
			shutdown()
			return
		case <-runtimeInfo.ReloadCfgChannel:
		}
	}
}

func main() {
	g.SetDefaultCliOptions()

	parser := flags.NewParser(&g.CliOptions, flags.PassDoubleDash | flags.PrintErrors)
	parser.ParseArgs(os.Args)

	switch {
	case g.CliOptions.Version:
		showVersion()
	default:
		mainEventLoop()
	}
}
