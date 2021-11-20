package process

import (
	"gm/internal/conf"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

type gmServerGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

/**
 * Erlang 处理进程相关
 *
 */
func (trpc *gmServerGenServer) Init(p *ergo.Process, args ...interface{}) (state interface{}) {
	trpc.process = p
	return nil
}
func (trpc *gmServerGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	// fmt.Printf("gmServerGenServer ({%s, %s}): HandleCast: %#v\n", trpc.process.name, trpc.process.Node.FullName, message)
	// trpc.v <- message
	return "noreply", state
}
func (trpc *gmServerGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	// fmt.Printf("gmServerGenServer ({%s, %s}): HandleCall: %#v, From: %#v\n", trpc.process.name, trpc.process.Node.FullName, message, from)
	return "reply", message, state
}
func (trpc *gmServerGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	// fmt.Printf("gmServerGenServer ({%s, %s}): HandleInfo: %#v\n", trpc.process.name, trpc.process.Node.FullName, message)
	// trpc.v <- message
	return "noreply", state
}
func (trpc *gmServerGenServer) Terminate(reason string, state interface{}) {
	// fmt.Printf("\ngmServerGenServer ({%s, %s}): Terminate: %#v\n", trpc.process.name, trpc.process.Node.FullName, reason)
}

// 分裂 Erlang 进程
func Spawn(n *ergo.Node) (*ergo.Process, error) {
	return n.Spawn(conf.ProcessName, ergo.ProcessOptions{}, &gmServerGenServer{}, nil)
}
