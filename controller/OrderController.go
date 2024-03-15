package controller

import (
	"assignment2/model"
	"assignment2/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (oc *orderController) GetAllOrders(ctx *gin.Context) {
	orders, err := oc.orderRepository.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfull get data",
		"data":    orders,
	})
}

func (oc *orderController) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	order, err := oc.orderRepository.CreateOrder(newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success create order",
		"data":    order,
	})
}

func (oc *orderController) UpdateOrder(ctx *gin.Context) {
	newOrder := model.Order{}
	orderID := ctx.Param("orderID")

	err := ctx.ShouldBindJSON(&newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	newOrder.OrderID = orderID

	order, err := oc.orderRepository.UpdateOrder(orderID, newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "success updating",
		"new_data:": order,
	})

}

func (oc *orderController) DeletingOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

	err := oc.orderRepository.DeleteOrder(orderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})
}
