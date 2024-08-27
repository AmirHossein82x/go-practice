package main

import (
	"notification/entities"
	"notification/services"
)

func main() {
	user := entities.User{
		Id:               1,
		Name:             "ali",
		Email:            "sjlfjs@",
		PhoneNumber:      "1234566",
		NotificationType: "Sms",
	}
	order := entities.Order{
		Id:    1,
		User:  user,
		Date:  "1403-02-02",
		Price: 12.000,
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&order)

}
