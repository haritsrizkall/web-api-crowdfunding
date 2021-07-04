package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(CampaignID int) ([]Transaction, error)
	GetByUserID(UserID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(CampaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("campaign_id = ?", CampaignID).Preload("Campaign").Preload("User").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(UserID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("user_id = ?", UserID).Preload("Campaign.CampaignImages", "campaign_images.is_primary = ?", 1).Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
