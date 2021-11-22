package conf

/**
 * 配置文件读取与保存模板以及默认配置设置
 *
 */

import "github.com/spf13/viper"

var (
	NodeName             string // 节点名称
	NodeCookie           string // 密钥
	NodeListenRangeBegin uint16 // 节点动态监听端口开始
	NodeListenRangeEnd   uint16 // 节点动态监听端口结束
	NodeEpmdPort         uint16 // 节点所在 epmd 监听端口
	NodeHidden           bool   // 节点是否隐藏
	ProcessName          string // 进程名称
	GmPort               uint16 // 后端管理接口所在端口
	NodeHandshakeVersion int    // 节点握手版本
)

func init() {
	viper.SetDefault("node.name", "gm@127.0.0.1")
	viper.SetDefault("node.cookie", "123")
	viper.SetDefault("node.listen_range_begin", uint16(40001))
	viper.SetDefault("node.listen_range_end", uint16(44000))
	viper.SetDefault("node.epmd_port", uint16(4369))
	viper.SetDefault("node.hidden", false)
	viper.SetDefault("node.handshake_version", 6)

	viper.SetDefault("process.name", "gm_process")

	viper.SetDefault("gm.port", uint16(80))

	viper.SetConfigFile("conf/config.toml")
	viper.ReadInConfig()

	viper.WriteConfigAs("conf/config.default.toml")

	NodeName = viper.GetString("node.name")
	NodeCookie = viper.GetString("node.cookie")
	NodeListenRangeBegin = uint16(viper.GetUint("node.listen_port_begin"))
	NodeListenRangeEnd = uint16(viper.GetUint("node.listen_port_end"))
	NodeEpmdPort = uint16(viper.GetUint("node.epmd_port"))
	NodeHidden = viper.GetBool("node.hidden")
	NodeHandshakeVersion = viper.GetInt("node.handshake_version")

	ProcessName = viper.GetString("process.name")

	GmPort = uint16(viper.GetUint("gm.port"))

}
