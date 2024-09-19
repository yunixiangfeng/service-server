package config

// import (
// 	"git.internal.attains.cn/attains-cloud/service-acs/core/config/app"
// 	"git.internal.attains.cn/attains-cloud/service-acs/core/config/pkg"
// )

type Config struct {
	App *app.Config `json:"app" xml:"app" yaml:"app"`
	Pkg *pkg.Config `json:"pkg" xml:"pkg" yaml:"pkg"`
}
