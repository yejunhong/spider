package model

import (
	// "fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct{
	Db *gorm.DB
}

func InitDb() *gorm.DB {
  var err error
	db, err := gorm.Open("mysql", "caiji:caijishiwo7788dd@tcp(202.43.91.26)/caiji?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true) // 不进行转换表名
	return db
}