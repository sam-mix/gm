package erl

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/halturin/ergo"
)

func Call() {

	// 节点选项
	nodeOps := ergo.NodeOptions{
		ListenRangeBegin: global.GVA_CONFIG.Erl.ListenRangeBegin,
		ListenRangeEnd:   global.GVA_CONFIG.Erl.ListenRangeBegin,
		Hidden:           global.GVA_CONFIG.Erl.Hidden,
		EPMDPort:         global.GVA_CONFIG.Erl.EpmdPort,
	}
	ergo.CreateNode(global.GVA_CONFIG.Erl.NodeName, global.GVA_CONFIG.Erl.Cookie, nodeOps)

}
