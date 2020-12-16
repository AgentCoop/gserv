package gserver

import (
	c "github.com/AgentCoop/gserv/internal/common"
	"net/http"
)

type ProxyInternalApi interface {

}

type HttpRequest struct {
	req *http.Request
}

func HandleRequest(r http.Request) {
	//for middleware := range m.Middleware{}
}


func StartServer() {

}

func NewRuntimeInfo() *RuntimeInfo {
	r := &RuntimeInfo{}
	r.ExitChannel = make(chan c.ExitSignal, 1)
	r.ReloadCfgChannel = make(chan c.ReloadCfgSignal, 1)
	return r
}
