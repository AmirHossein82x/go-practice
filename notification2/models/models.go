package models


type User struct {
	Id int
	FirstName string
	PhoneNumber string
	Email string
}



type UserNotificationOptions struct{
	Id int
	User
	NotificationType
}


type Order struct {
	Id int
	User 
	Date string
	Price float32

}

