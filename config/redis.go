package config

type Redis struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Db           int    `mapstructure:"db" json:"db" yaml:"db"`
	PoolSize     int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
	MinIdleConns int    `mapstructure:"minIdleConns" json:"minIdleConns" yaml:"minIdleConns"`
}
