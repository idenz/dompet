package config

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/jinzhu/configor"
)

type Configuration struct {
	Server   Server   `env:"server"`
	JWT      JWT      `env:"jwt"`
	Database Database `env:"database"`
}

type Server struct {
	ENV     string `env:"env"`
	AppName string `env:"app_name"`
	AppKey  string `env:"app_key"`
	Port    string `env:"port"`
	Version string `env:"version"`
}

type JWT struct {
	Secret string `env:"secret"`
}

type Database struct {
	DBHost        string `env:"db_host"`
	DBUser        string `env:"db_user"`
	DBPass        string `env:"db_pass"`
	DBPort        string `env:"db_port"`
	DBName        string `env:"db_name"`
	DBProvider    string `env:"db_provider"`
	DBSSL         string `env:"db_ssl"`
	DBTZ          string `env:"db_tz"`
	DBAutomigrate bool   `env:"db_automigrate"`
}

var Config *Configuration = &Configuration{}

func Load(path string) error {

	if path == "" {
		path = "config.yml"
	}

	err := configor.New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).Load(Config, path)
	if err != nil {
		return err
	}

	return nil
}

func (Configuration) String() string {
	sb := strings.Builder{}
	return sb.String()
}

func (c *Configuration) Raw() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}
