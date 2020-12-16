package gserver

import (
	n "github.com/AgentCoop/gserv/internal/common/net"
	c "github.com/AgentCoop/gserv/internal/common"
)

type RuntimeInfo struct {
	ReloadCfgChannel chan c.ReloadCfgSignal
	ExitChannel chan c.ExitSignal
}

type PortForwardingSchema int

const (
	Local PortForwardingSchema = iota
	Remote
)

//const (
//	Active ServerState = iota
//)

type ServerGroup struct {
	Name string
	Items []*n.Server
}

type PortForwarding struct {
	Schema PortForwardingSchema
}
