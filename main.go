package main

import (
	"fmt"

	"github.com/halturin/ergo"
	"github.com/spf13/viper"
)

var (
	nodeName             string // 节点名称
	nodeCookie           string // 密钥
	nodeListenRangeBegin uint16 // 节点动态监听端口开始
	nodeListenRangeEnd   uint16 // 节点动态监听端口结束
	nodeEpmdPort         uint16 // 节点所在 epmd 监听端口
	nodeHidden           bool   // 节点是否隐藏
)

func init() {
	viper.SetDefault("node.name", "gm@127.0.0.1")
	viper.SetDefault("node.cookie", "123")
	viper.SetDefault("node.listen_range_begin", uint16(40001))
	viper.SetDefault("node.listen_range_end", uint16(44000))
	viper.SetDefault("node.epmd_port", uint16(4369))
	viper.SetDefault("node.hidden", false)
	viper.SetConfigFile("conf/config.toml")
	viper.WriteConfigAs("conf/config.default.toml")
	nodeName = viper.GetString("node.name")
	nodeCookie = viper.GetString("node.cookie")
	nodeListenRangeBegin = uint16(viper.GetUint("node.listen_port_begin"))
	nodeListenRangeEnd = uint16(viper.GetUint("node.listen_port_end"))
	nodeEpmdPort = uint16(viper.GetUint("node.epmd_port"))
	nodeHidden = viper.GetBool("node.hidden")

}

func main() {
	// 节点选项
	nodeOps := ergo.NodeOptions{
		ListenRangeBegin: nodeListenRangeBegin,
		ListenRangeEnd:   nodeListenRangeEnd,
		Hidden:           nodeHidden,
		EPMDPort:         nodeEpmdPort,
	}
	node := ergo.CreateNode(nodeName, nodeCookie, nodeOps)
	fmt.Println(node)
}
