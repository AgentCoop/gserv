package main

import (
	//g "github.com/AgentCoop/gserv/internal/gserver"
	//m "github.com/AgentCoop/gserv/internal/middleware"
	g "github.com/AgentCoop/gserv/internal/gserver"
)

//import (
//	g "github.com/AgentCoop/gserv/internal/gserver"
//	m "github.com/AgentCoop/gserv/internal/middleware"
//)

func OnInit() {

}

func GetName() string {
	return "static"
}

func GetVersion() string {
	return "0.0.1"
}

func GetBuildDateTime() string {
	return "YYYY-MM-DD 00:00"
}

func EntryPoint(internalApi g.ProxyInternalApi) interface{} {
	return "Hello"
}
