package gserver

var CliOptions struct {
	RootDir string
	Version bool `long:"version" short:"v" description:""`
}

func SetDefaultCliOptions() {
	CliOptions.RootDir = "./assets/gserv"
}
