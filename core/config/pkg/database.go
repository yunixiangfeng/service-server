package pkg

type Database struct {
	Name      string `json:"name" xml:"name" yaml:"name"`
	Driver    string `json:"driver" xml:"driver" yaml:"driver"`
	Dsn       string `json:"dsn" xml:"dsn" yaml:"dsn"` // dsn连接字符串，不为空时使用该字符串连接，以下配置均无效
	Host      string `json:"host" xml:"host" yaml:"host"`
	Port      int    `json:"port" xml:"port" yaml:"port"`
	User      string `json:"user" xml:"user" yaml:"user"`
	Pwd       string `json:"pwd" xml:"pwd" yaml:"pwd"`
	DbName    string `json:"dbname" xml:"dbname" yaml:"dbname"`
	Charset   string `json:"charset" xml:"charset" yaml:"charset"`
	Collation string `json:"collation" xml:"collation" yaml:"collation"`
}
