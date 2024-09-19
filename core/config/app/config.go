package app

// Config app配置
type Config struct {
	Name       string      `json:"name" xml:"name" yaml:"name"`
	Env        string      `json:"env" xml:"env" yaml:"env"`                      // 环境
	Debug      bool        `json:"debug" xml:"debug" yaml:"debug"`                // debug模式
	Registries []*Registry `json:"registries" xml:"registries" yaml:"registries"` // 注册中心
	Server     *Server     `json:"server" xml:"server" yaml:"server"`
}
