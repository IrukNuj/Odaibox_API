package db

import (
	//"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("mysql", "root:@/odaibox_db")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
