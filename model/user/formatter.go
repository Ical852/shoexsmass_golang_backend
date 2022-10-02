package user

type UserFormatter struct {
	ID 			int		`json:"id"`
	FullName 	string	`json:"full_name"`
	Email 		string	`json:"email"`
	PhoneNumber int		`json:"phone_number"`
	Address 	string	`json:"address"`
	Password 	string	`json:"password"`
	Image 		string	`json:"image"`
	Balance 	int		`json:"balance"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Password:    user.Password,
		Image:       user.Image,
		Balance:     user.Balance,
	}

	return formatter
}