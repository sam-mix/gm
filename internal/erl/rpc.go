package erl

import (
	"gm/internal/conf"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

type RPC struct {
	Node string
	Mod  string
	Fun  string
	Args []etf.Term
}

func Call(node, mod, fun string, args ...etf.Term) etf.Term {
	rpc := RPC{
		Node: node,
		Mod:  mod,
		Fun:  fun,
		Args: args,
	}
	n, err := Create()
	if err != nil {
		panic(err)
	}
	gs := &ErlGenServer{res: make(chan interface{}, 2)}
	p, err := n.Spawn(conf.ProcessName, gen.ProcessOptions{}, gs)
	if err != nil {
		panic(err)
	}
	if p.Send(p.Self(), rpc); err != nil {
		panic(err)
	}
	return waitForResultWithValue(gs.res)
}

func waitForResultWithValue(w chan interface{}) etf.Term {
	select {
	case v := <-w:
		return v
	case <-time.After(time.Second * time.Duration(2)):
		return "time out"
	}
}
