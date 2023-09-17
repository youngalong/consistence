// Package config @Author: youngalone [2023/9/17]
package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	mysqlViper *viper.Viper
	MysqlCfg   mysql
)

func Init() {
	viper.SetConfigFile("./config/settings.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("配置文件加载成功")
	}

	// 初始化MySQL
	mysqlViper = viper.Sub("settings.mysql")
	if mysqlViper == nil {
		log.Fatal("配置文件缺失MySQL项")
	}
	MysqlCfg = InitMysql(mysqlViper)
}

type mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Schema   string
}

func InitMysql(mysqlViper *viper.Viper) mysql {
	return mysql{
		Host:     mysqlViper.GetString("host"),
		Port:     mysqlViper.GetString("port"),
		Username: mysqlViper.GetString("username"),
		Password: mysqlViper.GetString("password"),
		Schema:   mysqlViper.GetString("schema"),
	}
}
