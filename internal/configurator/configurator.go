package configurator

var (
	ProxyServerCfg *ProxyServerConfig
)

func OnInit() {
	ProxyServerCfg = &ProxyServerConfig{
		Middleware:   Middleware{},
		VirtualHosts: nil,
	}
}