package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-service/src/services"
	"strconv"
)

func Init() {
	e := echo.New()

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
			return c.JSON(http.StatusOK, services.CreateNewOrder())
		})
		orders.POST("/rabbit-test", func(c echo.Context) error {
			msg := c.QueryParam("msg")
			return c.JSON(http.StatusOK, services.PushMessage(msg))
		})
		orders.POST("/users/:id", func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			return c.JSON(http.StatusOK, services.GetUserById(uint(id)))
			//return c.JSON(http.StatusOK, services.PushMessage(msg))
		})
	}

	e.Logger.Fatal(e.Start(":1313"))

}
