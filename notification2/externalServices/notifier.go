package externalServices

import "notification/models"


type Notifier interface{
	SendNotify(receiver string, message string)
}


func NewNotifier(notificationType []models.NotificationType) []Notifier{
	notifires := []Notifier{}
	for _, item:= range notificationType{
		if item == "email"{
			notifires = append(notifires, NewEmailService())
		}else if item == "Sms"{
			notifires = append(notifires, NewSmsSmsService())

		}
	}
	
	return notifires
}