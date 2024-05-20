package flag_vars

import (
	"flag"
	"net"
)

var (
	host = flag.String("host", "0.0.0.0", "listen host.")
	port = flag.String("port", "80", "Port to run the server.")

	root = flag.String("root", "", "specify the root.")

	runmode = flag.String("runmode", "", "runmode default is empty")

	logfile  = flag.String("logfile", "", "logfile default os stdout")
	loglevel = flag.String("loglevel", "error", "log level")
)

func GetAddress() string {
	return net.JoinHostPort(GetHost(), GetPort())
}

func GetHost() string {
	return *host
}

func GetPort() string {
	return *port
}

func GetRoot() string {
	return *root
}

func GetRunMode() string {
	return *runmode
}

func GetLogfile() string {
	return *logfile
}

func GetLoglevel() string {
	return *loglevel
}

func init() {
	flag.Parse()
}
