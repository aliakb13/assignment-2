package main

import (
	"assignment2/controller"
	"assignment2/lib"
	"assignment2/model"
	"assignment2/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.InitDatabase()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Order{}, &model.Item{})
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/orders", orderController.GetAllOrders)
	ginEngine.POST("/orders", orderController.CreateOrder)
	ginEngine.PUT("/orders/:orderID", orderController.UpdateOrder)
	ginEngine.DELETE("/orders/:orderID", orderController.DeletingOrder)

	err = ginEngine.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
