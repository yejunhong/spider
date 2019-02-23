package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() *gorm.DB {
  	var err error
	db, err := gorm.Open("mysql", "root:root@/paimei?charset=utf8")
	if err == nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	return db
}