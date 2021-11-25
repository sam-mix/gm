package erl

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

type gmGenSrv struct {
	ergo.GenServer
}

func (dgs *gmGenSrv) Init(p *ergo.Process, args ...interface{}) interface{} {
	return nil
}

func (dgs *gmGenSrv) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	return "noreply", state
}

func (dgs *gmGenSrv) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	return "reply", message, state
}

func (dgs *gmGenSrv) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	return "noreply", state
}

func (dgs *gmGenSrv) Terminate(reason string, state interface{}) {
}

func Call(peerNode, mod, fun string, args ...etf.Term) (etf.Term, error) {
	// 节点选项
	nodeOps := ergo.NodeOptions{
		ListenRangeBegin: global.GVA_CONFIG.Erl.ListenRangeBegin,
		ListenRangeEnd:   global.GVA_CONFIG.Erl.ListenRangeBegin,
		EPMDPort:         global.GVA_CONFIG.Erl.EpmdPort,
	}
	node := ergo.CreateNode(global.GVA_CONFIG.Erl.NodeName, global.GVA_CONFIG.Erl.Cookie, nodeOps)
	p, err := node.Spawn(global.GVA_CONFIG.Erl.ProcessName, ergo.ProcessOptions{}, new(gmGenSrv), nil)
	if err != nil {
		return nil, err
	}
	return p.CallRPC(peerNode, mod, fun, args...)
}
