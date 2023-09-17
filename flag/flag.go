package flag

import (
	"github.com/spf13/pflag"
)

var (
	host string
	port int
)

func Init() {
	pflag.StringVar(&host, "host", "127.0.0.1", "service host address.")
	pflag.IntVarP(&port, "port", "P", 3306, "service host port.")
}
