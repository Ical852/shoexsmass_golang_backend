package user

type RegisterUserInput struct {
	FullName 	string	`json:"full_name" binding:"required"`
	Email 		string	`json:"email" binding:"required"`
	PhoneNumber int		`json:"phone_number" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
}

type LoginInput struct {
	Email 		string `json:"email" form:"email" binding:"required,email"`
	Password 	string `json:"password" form:"password" binding:"required"`
}

type CheckEmail struct {
	Email 		string `json:"email" binding:"required"`
}

type FormUpdateUserInput struct {
	ID         	int		`json:"id" binding:"required"`
	FullName 	string	`json:"full_name" binding:"required"`
	Email 		string	`json:"email" binding:"required"`
	PhoneNumber int		`json:"phone_number" binding:"required"`
	Password 	string	`json:"password"`
	Address 	string 	`json:"address"`
	Balance		int 	`json:"balance"`
}

type UploadAvatar struct {
	ID int `uri:"id" binding:"required"`
}