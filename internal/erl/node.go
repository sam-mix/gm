package erl

import (
	"gm/internal/conf"

	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/node"
)

var (
	nodeOps node.Options
)

func init() {
	// 节点选项
	nodeOps = node.Options{
		ListenRangeBegin: conf.NodeListenRangeBegin,
		ListenRangeEnd:   conf.NodeListenRangeEnd,
		Hidden:           conf.NodeHidden,
		EPMDPort:         conf.NodeEpmdPort,
		HandshakeVersion: conf.NodeHandshakeVersion,
	}

}

// 创建 erlang 节点
func Create() (node.Node, error) {
	return ergo.StartNode(conf.NodeName, conf.NodeCookie, nodeOps)
}
