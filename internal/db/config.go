package db

import "fmt"

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     int `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (c *Config) getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc-Local",
		c.Username, c.Password, c.Host, c.Port, c.Database)
}