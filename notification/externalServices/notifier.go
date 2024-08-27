package externalServices

import "notification/entities"

type Notifier interface {
	SendNotify(receiver string, message string)
}

func NewNotifier(notificationType entities.NotificationType) Notifier {
	switch notificationType {
	case "Email":
		return NewEmailService()
	case "Sms":
		return NewSmsService()
	default:
		return nil

	}
}
