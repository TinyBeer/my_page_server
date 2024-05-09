package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		Username string
		Password string
		Host     string
		Port     int
		DbName   string
		Timeout  int
	}
}
