package config

import (
	"github.com/jinzhu/gorm"  // mysql
	_ "github.com/jinzhu/gorm/dialects/mysql" // orm 
)

var (
	db *gorm.DB
)

func Connect() {

	// d, err := gorm.Open("mysql", "root:Vijay@123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:Vijay@123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	//database name simplerest

	if err != nil {

		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {

	return db
}
