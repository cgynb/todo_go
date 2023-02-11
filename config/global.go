package config

import "time"

type Config struct {
	PageSize    int
	TokenConfig tokenConfig
	MysqlConfig mysqlConfig
}
type tokenConfig struct {
	SecretKey  []byte
	EffectTime time.Duration
}

type mysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DNS      string
}

var Conf Config
