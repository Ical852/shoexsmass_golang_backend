package transaction

import (
	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
	"shoexsmass/model/transactiondetail"
	"shoexsmass/model/user"
)

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	GetByID(ID int) (Transaction, error)
}

type repository struct {
	db *gorm.DB
	repository transactiondetail.Repository
	userRepository user.Repository
}

func NewRepository(db *gorm.DB, transactiondetail transactiondetail.Repository, userRepository user.Repository) *repository {
	return &repository{db, transactiondetail, userRepository}
}

func (r *repository) Create(transaction Transaction) (Transaction, error) {
	user, err := r.userRepository.FindById(transaction.UserID)
	if err != nil {
		return transaction, err
	}

	transaction.User = user

	if transaction.PayMethod == "midtrans" {
		midclient :=midtrans.NewClient()
		midclient.APIEnvType = midtrans.Sandbox

		snapGateway := midtrans.SnapGateway{
			Client: midclient,
		}

		snapReq := &midtrans.SnapReq{
			CustomerDetail: &midtrans.CustDetail{
				FName: transaction.User.FullName,
				Email: transaction.User.Email,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID: transaction.OrderID,
				GrossAmt: int64(transaction.GrossAmount),
			},
		}

		snapTokenResp, err := snapGateway.GetToken(snapReq)
		if err != nil {
			return transaction, err
		}

		transaction.PaymentUrl = snapTokenResp.RedirectURL
	} else {
		transaction.PaymentUrl = "ssmPay"
		user.Balance = user.Balance - transaction.GrossAmount
		_, err := r.userRepository.Update(user)
		if err != nil {
			return transaction, err
		}
	}

	transaction.Status = "order"

	err = r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	for _, tDetail := range transaction.TransactionDetail {
		tDetail.TransactionID = transaction.ID
		_, _ = r.repository.Create(tDetail)
	}

	return transaction, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("User").Preload(	"TransactionDetail.Product.Category").Preload(	"TransactionDetail.Product.ProductImage").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Preload("User").Preload(	"TransactionDetail.Product.Category").Preload(	"TransactionDetail.Product.ProductImage").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
