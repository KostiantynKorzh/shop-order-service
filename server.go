package main

import (
	"order-service/src/db"
	"order-service/src/model"
	"order-service/src/route"
)

func initDb() {
	database := db.InitDb()
	database.AutoMigrate(&model.Order{}, &model.OrderItem{})
}

func main() {

	//rabbit.Init()
	initDb()
	route.Init()

}
