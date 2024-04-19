package config

import (
	"github.com/jihnzu/gorm"
	_ "github.com/jihnzu/gorm/dialects/mysql"   //see in documentation
)

var {
	db * gorm.DB
}

func Connect(){
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err := nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}