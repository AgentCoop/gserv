package middleware

import (
	"fmt"
	"github.com/AgentCoop/gserv/internal/common/fs"
	g "github.com/AgentCoop/gserv/internal/gserver"
	"path"
	"plugin"
)

var (
	middlewareMap MiddlewareMap
)

func loadMiddlewares(rootDir string) {
	scanner := fs.NewScanner(rootDir)
	scanner.IncludePattern = "*.so|*.dll"
	fColl := scanner.Run()
	for _, v := range fColl.Items {
		p, err := plugin.Open(v.AbsPath)
		if err != nil { panic(err) }
		mid := &Middleware{}
		if err != nil {
			panic(err)
		}

		f, err := p.Lookup(EntryPointFuncName)
		if err != nil { panic(err) }
		mid.Entrypoint = f.(func(api g.ProxyInternalApi) interface{})

		f, _ = p.Lookup(GetNameFuncName)
		mid.Name = f.(func() string)()

		fmt.Printf("Loaded plugin: %s\n", mid.Name)
		fmt.Printf("Entry point plugin: %v\n", mid.Entrypoint(nil))

		middlewareMap[mid.Name] = mid
	}
}

func OnInit() {
	middlewareMap = make(MiddlewareMap)
	midRootDir := path.Join(g.CliOptions.RootDir, "middleware")
	loadMiddlewares(midRootDir)
}
