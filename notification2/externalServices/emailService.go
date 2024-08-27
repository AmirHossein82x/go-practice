package externalServices

import "fmt"


type EmailService struct {

}


func (emailService *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email:::: receiver: %s, message: %s", receiver, message)
}

func NewEmailService() *EmailService{
	return &EmailService{}
}