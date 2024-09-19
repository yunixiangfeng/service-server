package pkg

type Config struct {
	Database []*Database `json:"database" xml:"database" yaml:"database"`
	Etcd     []*Etcd     `json:"etcd" xml:"etcd" yaml:"etcd"`
	Logger   []*Logger   `json:"logger" xml:"logger" yaml:"logger"`
	Redis    []*Redis    `json:"redis" xml:"redis" yaml:"redis"`
}
