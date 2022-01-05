package config

type Server struct {
	HttpPort      int `mapstructure:"httpPort" json:"httpPort" yaml:"httpPort"`                // http端口
	WebSocketPort int `mapstructure:"webSocketPort" json:"webSocketPort" yaml:"webSocketPort"` // websocket端口
	RpcPort       int `mapstructure:"rpcPort" json:"rpcPort" yaml:"rpcPort"`                   // rpc端口
}
