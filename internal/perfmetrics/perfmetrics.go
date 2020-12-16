package perfmetrics

import (
	"fmt"
	c "github.com/AgentCoop/gserv/internal/common"
	g "github.com/AgentCoop/gserv/internal/gserver"
	"time"
)

type NetUsageKeyVal struct {
	Hostname string
	Val interface{}
}

type perfMetrics struct {
	Network struct {
		BytesSent <-chan NetUsageKeyVal
		BytesReceived <-chan NetUsageKeyVal
	}
}

func NewPerfMetrics() *perfMetrics {
	p := &perfMetrics{}
	p.Network.BytesSent = make(chan NetUsageKeyVal)
	p.Network.BytesReceived = make(chan NetUsageKeyVal)
	return p
}

func (p *perfMetrics) Run(runtimeInfo *g.RuntimeInfo) {
	//var bytesReceivedAccum int
	//var bytesSentAccum int
	oneSecTick := time.Tick(time.Second)
	_10SecTick := time.Tick(10 * time.Second)
	oneMinTick := time.Tick(time.Minute)
	var i int
	for {
		select {
		case <-oneSecTick:
			fmt.Printf("One second tick %d\n", i + 1)
			i++
		case <-_10SecTick:
			fmt.Printf("Every 10 seconds tick\n")
			runtimeInfo.ExitChannel <- c.ExitSignal{}
		case <-oneMinTick:
		}
	}
}

func (p *perfMetrics) Shutdown() {
	fmt.Printf("Run shutdown routine\n")
}

