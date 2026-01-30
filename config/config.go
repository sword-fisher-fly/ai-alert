package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	Server   Server   `json:"Server"`
	Database Database `json:"Database"`
	Jwt      Jwt      `json:"Jwt"`
}

type Server struct {
	Mode           string `json:"mode"`
	Port           string `json:"port"`
	EnableElection bool   `json:"enableElection"`
}

type Database struct {
	Type    string `json:"type"`    // mysql 或 sqlite
	Host    string `json:"host"`    // MySQL 主机地址
	Port    string `json:"port"`    // MySQL 端口
	User    string `json:"user"`    // MySQL 用户名
	Pass    string `json:"pass"`    // MySQL 密码
	DBName  string `json:"dbName"`  // MySQL 数据库名
	Timeout string `json:"timeout"` // MySQL 连接超时
	Path    string `json:"path"`    // SQLite 数据库文件路径
}

type Jwt struct {
	Expire int64 `json:"expire"`
}

var (
	configFile = "config/config.yaml"
)

func InitConfig() App {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Read config failed:", err)
	}
	var config App
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal("Config parse failed:", err)
	}
	return config
}
