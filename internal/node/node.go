package node

/**
 * Erlang node 节点相关
 *
 */

import (
	"gm/internal/conf"

	"github.com/halturin/ergo"
)

var (
	nodeOps ergo.NodeOptions
)

func init() {
	// 节点选项
	nodeOps = ergo.NodeOptions{
		ListenRangeBegin: conf.NodeListenRangeBegin,
		ListenRangeEnd:   conf.NodeListenRangeEnd,
		Hidden:           conf.NodeHidden,
		EPMDPort:         conf.NodeEpmdPort,
	}

}

// 创建 erlang 节点
func Create() *ergo.Node {
	return ergo.CreateNode(conf.NodeName, conf.NodeCookie, nodeOps)
}
