package main

import (
	"order-service/db"
	"order-service/model"
	"order-service/route"
)

func initDb() {
	database := db.InitDb()
	database.AutoMigrate(&model.Order{}, &model.OrderItem{})
}

func main() {
	initDb()
	route.Init()
}
