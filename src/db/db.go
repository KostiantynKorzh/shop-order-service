package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db := connectDb()
	return Db
}

func connectDb() *gorm.DB {
	dsn := "root:" + os.Getenv("MYSQL_ROOT_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_URL") + ":3306)/orders_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db

}
