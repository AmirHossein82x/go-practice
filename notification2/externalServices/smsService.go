package externalServices

import "fmt"


type SmsService struct{

}


func (smsService *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("Sms:::: receiver: %s, message: %s", receiver, message)
}

func NewSmsSmsService() *SmsService{
	return &SmsService{}
}