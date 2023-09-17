// Package main @Author: youngalone [2023/9/17]
package main

import (
	"consistence/config"
	"consistence/database/mysql"
)

func main() {
	config.Init()
	mysql.Init()
}
