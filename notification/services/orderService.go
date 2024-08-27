package services

import (
	"fmt"
	"notification/entities"
	"notification/externalServices"
)


type OrderService struct {
	Notifier externalServices.Notifier

}

func (orderService *OrderService) CreateOrder(order *entities.Order)(*entities.Order){
	fmt.Printf("Order created: %v\n", order)
	// send sms or email
	orderService.Notifier = externalServices.NewNotifier(order.User.NotificationType)
	orderService.Notifier.SendNotify(order.PhoneNumber, "message send")
	
	return order
}



func NewOrderService() *OrderService{
	return &OrderService{}
}