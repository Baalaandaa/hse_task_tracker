package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	_errConfig = fmt.Errorf("config error")
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"postgres"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func LoadAppConfig() (*Config, error) {
	return LoadConfig("/app/config/config.yml")
}

func LoadConfig(path string) (*Config, error) {
	// #nosec
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("file not found %s %w", path, _errConfig)
	}
	yamlDecoder := yaml.NewDecoder(file)

	cfg := &Config{}
	err = yamlDecoder.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err.Error(), _errConfig)
	}

	return cfg, nil
}
