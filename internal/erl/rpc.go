package erl

import (
	"gm/internal/node"
	"gm/internal/process"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

/**
 * rpc 调用处理
 */

type GmErlangServer struct {
	Node    *ergo.Node    // Erlang 节点
	Process *ergo.Process // Erlang 执行进程
}

func NewServer() (*GmErlangServer, error) {
	n := node.Create()
	p, err := process.Spawn(n)
	return &GmErlangServer{
		Node:    n,
		Process: p,
	}, err
}

// rpc call
func (s *GmErlangServer) Call(nodeName, m, f string, a ...etf.Term) (etf.Term, error) {
	return s.Process.CallRPCWithTimeout(60, nodeName, m, f, a)
}
