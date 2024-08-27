package entities

type User struct {
	Id int
	Name string
	Email string
	PhoneNumber string
	NotificationType NotificationType
	
}


type SendNotificationOption struct{
	Id int
	Title string
}


type UserNotificationOption struct {
	Id int
	User
	SendNotificationOption
}



type Order struct {
	Id int
	User
	Date string
	Price float64

}