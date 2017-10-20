package config

import (
	"fmt"
	"github.com/koding/multiconfig"
)

var C = new(ServerConfig)

func init() {
	multiconfig.MustLoadWithPath("config.json", C)
}

type GlobalConfig struct {
	Listen string `default:"0.0.0.0:8080"`
	Debug  bool   `default:"true"`
}

type PostgresConfig struct {
	Address        string `required:"true"`
	Port           int    `default:"5432"`
	Db             string `required:"true"`
	User           string `required:"true"`
	Password       string `required:"true"`
	MaxConnections int    `default:"10"`
	MaxIdles       int    `default:"3"`
}

func (p PostgresConfig) String() string {
	fmtString := "user='%s' dbname='%s' password='%s' host='%s' port=%d"
	return fmt.Sprintf(fmtString, p.User, p.Db, p.Password, p.Address, p.Port)
}

type AliOssConfig struct {
	Endpoint  string `required:"true"`
	AccessID  string `required:"true"`
	AccessKey string `required:"true"`
}

type ServerConfig struct {
	Global   GlobalConfig
	Postgres PostgresConfig
	AliOss   AliOssConfig
}
