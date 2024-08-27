package services

import (
	"fmt"
	"notification/externalServices"
	"notification/models"
)


type OrderService struct {
	Notifier [] externalServices.Notifier

}


func (orderService *OrderService) CreateOrder(order models.Order, options []models.NotificationType){
	fmt.Printf("order created")

	orderService.Notifier = externalServices.NewNotifier(options)
	for _, service := range orderService.Notifier{
		service.SendNotify(order.PhoneNumber, "sucess")
	}
}


func NewOrder()(*OrderService){
	return &OrderService{}
}