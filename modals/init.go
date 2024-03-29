package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	dsn := `root:P@ssword123@tcp(127.0.0.1:3306)/indexexample?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		panic(err)
	}

	DB = db
}
