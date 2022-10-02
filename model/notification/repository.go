package notification

import "gorm.io/gorm"

type Repository interface {
	Create(notification Notification) (Notification, error)
	GetById(ID int) (Notification, error)
	Delete(ID int) (Notification, error)
	GetByUserID(userID int) ([]Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(notification Notification) (Notification, error) {
	err := r.db.Create(&notification).Error
	if err != nil {
		return notification, err
	}

	return notification, nil
}

func (r *repository) GetById(ID int) (Notification, error) {
	var notification Notification
	err := r.db.Where("id = ?", ID).Preload("User").Find(&notification).Error
	if err != nil {
		return notification, err
	}

	return notification, nil
}

func (r *repository) Delete(ID int) (Notification, error) {
	notif, err := r.GetById(ID)
	if err != nil {
		return notif, err
	}

	err = r.db.Delete(&notif, ID).Error
	if err != nil {
		return notif, err
	}

	return notif, nil
}

func (r *repository) GetByUserID(userID int) ([]Notification, error) {
	var notifications []Notification
	err := r.db.Where("user_id = ?", userID).Preload("User").Find(&notifications).Error
	if err != nil {
		return notifications, err
	}

	return notifications, nil
}

