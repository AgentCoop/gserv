package configurator

type VirtualHost struct {
	Pattern string
	Middlewares []string
}

type Middleware struct {
	RootDir	string
}

type ProxyServerConfig struct {
	Middleware Middleware
	VirtualHosts []VirtualHost
}

