package topup

import "shoexsmass/model/user"

type TopUpFormat struct {
	ID 			int					`json:"id"`
	UserID 		int					`json:"user_id"`
	OrderID 	string				`json:"order_id"`
	GrossAmount int					`json:"gross_amount"`
	Status 		string				`json:"status"`
	PaymentUrl 	string				`json:"payment_url"`
	User 		user.UserFormatter	`json:"user"`
}

func FormatTopUp(topUp TopUp) TopUpFormat {
	topUpFormatted := TopUpFormat{
		ID:          topUp.ID,
		UserID:      topUp.UserID,
		OrderID:     topUp.OrderID,
		GrossAmount: topUp.GrossAmount,
		Status:      topUp.Status,
		PaymentUrl:  topUp.PaymentUrl,
		User:        user.FormatUser(topUp.User),
	}

	return topUpFormatted
}

func FormatTopUps(topUps []TopUp) []TopUpFormat {
	var topUpsFormatted []TopUpFormat
	for _, topUp := range topUps {
		topUpsFormatted = append(topUpsFormatted, FormatTopUp(topUp))
	}

	return topUpsFormatted
}