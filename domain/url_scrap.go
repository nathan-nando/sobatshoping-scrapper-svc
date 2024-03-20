package domain

type UrlScrap struct {
	Name   string `json:"name" yaml:"name"`
	Domain string `json:"domain" yaml:"domain"`
	Param  string `json:"param" yaml:"param"`
}
