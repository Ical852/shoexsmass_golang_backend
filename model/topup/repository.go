package topup

import (
	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

type Repository interface {
	TopUP(topUp TopUp) (TopUp, error)
	GetByUserID(userID int) ([]TopUp, error)
	GetByID(ID int) (TopUp, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) TopUP(topUp TopUp) (TopUp, error) {
	midclient :=midtrans.NewClient()
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: topUp.User.FullName,
			Email: topUp.User.Email,
			Phone: string(topUp.User.PhoneNumber),
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: topUp.OrderID,
			GrossAmt: int64(topUp.GrossAmount),
		},
	}

	topUp.Status = "order"

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return topUp, err
	}

	topUp.PaymentUrl = snapTokenResp.RedirectURL

	err = r.db.Create(&topUp).Error
	if err != nil {
		return topUp, err
	}

	return topUp, err
}

func (r *repository) GetByUserID(userID int) ([]TopUp, error) {
	var topUps []TopUp
	err := r.db.Where("user_id", userID).Preload("User").Find(&topUps).Error
	if err != nil {
		return topUps, err
	}

	return topUps, nil
}

func (r *repository) GetByID(ID int) (TopUp, error) {
	var topUp TopUp
	err := r.db.Where("id", ID).Preload("User").Find(&topUp).Error
	if err != nil {
		return topUp, err
	}

	return topUp, nil
}
