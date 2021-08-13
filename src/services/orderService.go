package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"order-service/db"
	"order-service/model"
	"strconv"
)

var FullOrder []model.FullOrderInfo

func GetAllOrders() []model.Order {
	var orders []model.Order
	db.Db.Find(&orders)
	return orders
}

func GetOrderById(id uint) model.Order {
	var order model.Order
	db.Db.Find(&order, id)
	return order
}

func AddNewItemToCart(userId uint, itemId uint, quantity uint) {
	if isAlreadyShopping(userId) {
		print("Already shopping...")
		var order = model.Order{}
		db.Db.Where("user_id = ? AND status = ?", userId, "SHOPPING").First(&order)
		currentId := order.ID
		newOrderItem := []model.OrderItem{{OrderId: currentId, ItemId: itemId, Quantity: quantity}}
		db.Db.Model(model.Order{}).Updates(model.Order{OrderItems: newOrderItem})
		db.Db.Create(&newOrderItem)
	} else {
		print("New cart...")
		order := model.Order{
			UserId:     userId,
			OrderItems: []model.OrderItem{{ItemId: itemId, Quantity: quantity}},
			Status:     "SHOPPING",
		}
		db.Db.Create(&order)
	}
}

func isAlreadyShopping(userId uint) bool {
	var order model.Order
	r := db.Db.Where("user_id = ? AND status = ?", userId, "SHOPPING").Find(&order)
	exists := r.RowsAffected > 0
	if exists {
		return true
	}
	return false
}

func findActiveShoppingCartForUserById(userId uint) model.Order {
	var order model.Order
	db.Db.Where("user_id = ? AND status = ?", userId, "SHOPPING").Find(&order)
	return order
}

func GetLastOrderForUserById(userId uint) []model.FullOrderInfo {
	var order = findActiveShoppingCartForUserById(userId)
	if order.ID != 0 {
		var orderItems []model.OrderItem
		//db.Db.Where("user_id = ? AND status = ?", userId, "SHOPPING").Preload("OrderItems").First(&order)
		//db.Db.Model(&order).Where("order_id = ?", order.ID).Association("OrderItems").Find(&orderItems)
		db.Db.Where("order_id = ? AND status = ?", order.ID, "SHOPPING").Find(&orderItems)
		return fetchOrderItemsInfo([]uint{4, 5, 2})
	}
	return nil
}

// TODO Add cache for prices

func fetchOrderItemsInfo(itemsIds []uint) []model.FullOrderInfo {
	itemServiceUrl := "http://localhost:8081/items/orders?ids="
	var idsFormatted string
	for id := range itemsIds {
		idsFormatted += strconv.Itoa(id) + ","
	}
	resp, _ := http.Get(itemServiceUrl + idsFormatted)
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(data)
	var fullOrder []model.FullOrderInfo
	err := json.Unmarshal([]byte(bodyString), &fullOrder)
	if err != nil {
		println(err.Error())
	}

	for i, order := range fullOrder {
		order.Quantity = uint(i + 1)
		fullOrder[i] = order
	}

	FullOrder = fullOrder

	return fullOrder
}

func Buy(userId uint) {
	var order = findActiveShoppingCartForUserById(userId)
	db.Db.Model(&order).Updates(model.Order{Status: "PENDING PAYMENT"})

	SendOrderToPaymentService(order)
}
