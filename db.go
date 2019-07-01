package db

import "github.com/jinzhu/gorm"

var(
	db *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}