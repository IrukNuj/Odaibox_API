package db

import (
	"os"
	//"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iruknuj/odaibox_API/model"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	db.AutoMigrate(&model.Post{})
}
