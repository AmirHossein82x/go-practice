package main

import (
	"fmt"
	"notification/models"
	"notification/services"
)


func main(){
	user := models.User{Id: 1, FirstName: "amirhossein", Email: "amirhosseinG@gmail.com", PhoneNumber: "0912228282"}
	userOption1 := models.UserNotificationOptions{Id:1, User:  user, NotificationType: models.Sms}
	// userOption2 := models.UserNotificationOptions{Id:2, User:  user, NotificationType: models.Email}
	fmt.Println(userOption1)
	options := []models.NotificationType{userOption1.NotificationType}
	order := models.Order{
		User: user,
		Id: 1,
		Price: 12.000,
	}


	orderService := services.NewOrder()
	orderService.CreateOrder(order, options)

}