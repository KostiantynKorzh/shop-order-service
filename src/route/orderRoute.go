package route

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"order-service/src/services"
	"strconv"
)

func Init() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	orders := e.Group("/orders")
	{
		orders.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, services.GetAllOrders())
		})
		orders.GET("/:id", func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			return c.JSON(http.StatusOK, services.GetOrderById(uint(id)))
		})
		orders.POST("/", func(c echo.Context) error {
			jsonMap := make(map[string]interface{})
			json.NewDecoder(c.Request().Body).Decode(&jsonMap)
			services.AddNewItemToCart(
				uint(jsonMap["userId"].(float64)),
				uint(jsonMap["itemId"].(float64)),
				uint(jsonMap["quantity"].(float64)))
			return c.JSON(http.StatusOK, "Adding to cart")
		})
		orders.GET("/user-orders/:id", func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			return c.JSON(http.StatusOK, services.GetLastOrderForUserById(uint(id)))
		})
		orders.POST("/rabbit-test", func(c echo.Context) error {
			msg := c.QueryParam("msg")
			return c.JSON(http.StatusOK, services.PushMessage(msg))
		})
		orders.POST("/user-orders/:id", func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			services.Buy(uint(id))
			return c.JSON(http.StatusOK, "SENDING TO PAYMENTS...")
		})
	}

	e.Logger.Fatal(e.Start(":1313"))

}
