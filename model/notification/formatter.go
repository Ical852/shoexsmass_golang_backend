package notification

import "shoexsmass/model/user"

type NotificationFormatter struct {
	ID 		int			`json:"id" binding:"required"`
	UserID 	int			`json:"user_id" binding:"required"`
	Message string		`json:"message" binding:"required"`
	Date 	string		`json:"date" binding:"required"`
	User 	user.UserFormatter 	`json:"user" binding:"required"`
}

func FormatNotification(notification Notification) NotificationFormatter {
	formatted := NotificationFormatter{
		ID:      notification.ID,
		UserID:  notification.UserID,
		Message: notification.Message,
		Date:    notification.Date,
		User:    user.FormatUser(notification.User),
	}

	return formatted
}

func FormatNotifications(notifications []Notification) []NotificationFormatter {
	var formatted []NotificationFormatter
	for _, notification := range notifications {
		formatted = append(formatted, FormatNotification(notification))
	}

	return formatted
}