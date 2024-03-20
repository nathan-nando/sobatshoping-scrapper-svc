package configs

import "github.com/nathan-nando/sobatshoping-scrapper-svc/domain"

type AppMode string

const (
	Development AppMode = "development"
	Production  AppMode = "production"
)

type Config struct {
	Server  Server  `json:"server" yaml:"server"`
	MongoDB MongoDB `json:"mongo_db" yaml:"mongo_db"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Grpc    Grpc    `json:"grpc" yaml:"grpc"`
	Scrap   Scrap   `json:"scrap" yaml:"scrap"`
}

type Server struct {
	Host    string `json:"host" yaml:"host"`
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	Mode    string `json:"mode" yaml:"mode"`
}

type MongoDB struct {
	Product1 string `json:"product_1" yaml:"product_1"`
	Product2 string `json:"product_2" yaml:"product_2"`
}

type Redis struct {
	Host string `json:"host" yaml:"host"`
}

type Grpc struct {
	Host string `json:"host" yaml:"host"`
}

type Scrap struct {
	Urls       []domain.UrlScrap  `json:"urls" yaml:"urls"`
	UserAgents []domain.UserAgent `json:"user_agents" yaml:"user_agents"`
}
