package externalServices

import (
	"fmt"
)


type SmsService struct {

}


// func (e *SmsService) SendSms(order *entities.Order) {
// 	fmt.Printf("Sms sent: %v\n", order)
// }

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("Sms sent: receiver: %s,   message: %s\n",receiver , message)
}

func NewSmsService()*SmsService{
	return &SmsService{}
}