// Package model @Author: youngalone [2023/9/17]
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}
