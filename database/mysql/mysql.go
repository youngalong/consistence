// Package database @Author: youngalone [2023/9/17]
package mysql

import (
	"consistence/config"
	"consistence/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DSN string
var DB *gorm.DB
var err error

func Init() {
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlCfg.Username,
		config.MysqlCfg.Password,
		config.MysqlCfg.Host,
		config.MysqlCfg.Port,
		config.MysqlCfg.Schema,
	)
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("MySQL连接失败 %v", err)
		return
	}
	DB.AutoMigrate(&model.User{})
}
