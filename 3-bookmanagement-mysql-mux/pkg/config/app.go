package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:password@/indexer?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	db = d
}

func GetDB() *gorm.DB {
	return db
}
