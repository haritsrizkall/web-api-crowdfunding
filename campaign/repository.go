package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(ID int) ([]Campaign, error)
	FindById(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	SaveCampaignImage(image CampaignImage) (CampaignImage, error)
	MarkAllIsPrimaryFalse(ID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary =  1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserId(ID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", ID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", ID).First(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) SaveCampaignImage(image CampaignImage) (CampaignImage, error) {
	err := r.db.Save(&image).Error
	if err != nil {
		return CampaignImage{}, nil
	}

	return CampaignImage{}, nil
}

func (r *repository) MarkAllIsPrimaryFalse(ID int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", ID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
