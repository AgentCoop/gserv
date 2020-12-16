package middleware

import (
	g "github.com/AgentCoop/gserv/internal/gserver"
)

const OnInitFuncName 				= "OntInit"
const EntryPointFuncName 			= "EntryPoint"
const GetNameFuncName 				= "GetName"
const GetVersionFuncName 			= "GetVersion"
const GetBuildDateTime 				= "GetBuildDateTime"

type EntryPointFunc func(api g.ProxyInternalApi) interface{}

type Middleware struct {
	Name string
	Version string
	BuildDateTime string
	Entrypoint EntryPointFunc
}

type MiddlewareMap map[string]*Middleware
