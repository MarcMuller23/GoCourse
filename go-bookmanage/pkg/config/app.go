package config

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "marc:MarcMul@23@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
