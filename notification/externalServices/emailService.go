package externalServices

import (
	"fmt"
)


type EmailService struct {

}


// func (e *EmailService) SendEmail(order *entities.Order) {
// 	fmt.Printf("Email sent: %v\n", order)
// }

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email sent: receiver: %s,   message: %s\n",receiver , message)
}

func NewEmailService()*EmailService{
	return &EmailService{}
}