package server

import (
	"github.com/RobinUS2/tsxdb/rpc"
)

type Opts struct {
	rpc.OptsConnection `yaml:"connection"`
	TelnetPort         int    `yaml:"telnet_port"`
	TelnetHost         string `yaml:"telnet_host"`
}

func NewOpts() *Opts {
	return &Opts{
		OptsConnection: rpc.NewOptsConnection(),
	}
}
