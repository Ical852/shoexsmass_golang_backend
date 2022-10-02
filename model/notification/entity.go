package notification

import "shoexsmass/model/user"

type Notification struct {
	ID int
	UserID int
	Message string
	Date string
	User user.User
}