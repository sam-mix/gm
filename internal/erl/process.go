package erl

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

type ErlGenServer struct {
	gen.Server
	res chan interface{}
}

func (trpc *ErlGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	return message, gen.ServerStatusOK
}

func (trpc *ErlGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	switch m := message.(type) {
	case RPC:
		if v, e := process.CallRPC(m.Node, m.Mod, m.Fun, m.Args...); e != nil {
			trpc.res <- e
		} else {
			trpc.res <- v
		}
		return gen.ServerStatusOK

	}

	return gen.ServerStatusStop
}
