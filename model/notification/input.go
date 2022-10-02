package notification

type NotificationCreateInput struct {
	UserID 	int		`json:"user_id" binding:"required"`
	Message string	`json:"message" binding:"required"`
	Date 	string	`json:"date" binding:"required"`
}

type SomethingWithID struct {
	ID int `uri:"id" binding:"required"`
}