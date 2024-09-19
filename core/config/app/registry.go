package app

type Registry struct {
	Type       string `json:"type" xml:"type" yaml:"type"`
	Connection string `json:"connection" xml:"connection" yaml:"connection"`
}
