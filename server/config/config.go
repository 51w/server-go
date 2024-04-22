package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// oss
	Local    Local    `mapstructure:"local" json:"local" yaml:"local"`

	// 跨域配置
	Cors     CORS     `mapstructure:"cors" json:"cors" yaml:"cors"`
}
