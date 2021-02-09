package config

import (
	"github.com/tkanos/gonfig"
)
type Config struct {
	DB *Configuration
}

type Configuration struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}


func GetConfig() *Config {
	return &Config{
	configuration := Configuration{}
	err := gonfig.GetConf("/config.json", &configuration)	
	Dialect  =  configuration.Dialect
	Host     =  configuration.Host
	Port     = configuration.Port
	Username = configuration.Username
	Password = configuration.Password
	Name     = configuration.Name
	Charset  = configuration.Charset

}
}

