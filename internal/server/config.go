package server

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gitea.programmerfamily.com/go/product/internal/pkg/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port   int            `json:"port" yaml:"port"`
	Db     *db.Config     `json:"db" yaml:"db"`
	Logger *logger.Config `json:"logger" yaml:"logger"`
}

func newConfigFromFile(filename string) (*Config, error) {
	var cfg *Config

	// 读取配置文件
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解析 yaml
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}