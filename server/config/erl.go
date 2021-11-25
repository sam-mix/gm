package config

// Erlang 相关配置
type Erl struct {
	Cookie           string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`                                   // 密钥
	EpmdPort         uint16 `mapstructure:"epmd-port" json:"epmdPort" yaml:"epmd-port"`                           // 节点所在 epmd 监听端口
	Hidden           bool   `mapstructure:"hidden" json:"hidden" yaml:"hidden"`                                   // 节点是否隐藏
	ListenRangeBegin uint16 `mapstructure:"listen-range-begin" json:"listenRangeBegin" yaml:"listen-range-begin"` // 节点动态监听端口开始
	ListenRangeEnd   uint16 `mapstructure:"listen-range-end" json:"listenRangeEnd" yaml:"listen-range-end"`       // 节点动态监听端口结束
	NodeName         string `mapstructure:"node-name" json:"nodeName" yaml:"node-name"`                           // 节点名称
	ProcessName      string `mapstructure:"process-name" json:"processName" yaml:"process-name"`                  // 进程名称
}
