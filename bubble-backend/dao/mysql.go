package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnMySQL() {
	dst := "root:xxx@tcp(xxx:3306)/bubble?charset=utf8&parseTime=True"
	var err error
	DB, err = gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		panic("mysql open failed")
	}
}
