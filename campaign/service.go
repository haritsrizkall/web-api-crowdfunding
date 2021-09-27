package campaign

import (
	"errors"
	"fmt"
	"strings"
)

type Service interface {
	GetCampaigns(ID int) ([]Campaign, error)
	GetCampaignByID(ID int) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(input CreateCampaignInput, ID int) (Campaign, error)
	CreateImageCampaign(input CreateImageCampaignInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

// tambah user
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(ID int) ([]Campaign, error) {
	if ID == 0 {
		campaigns, err := s.repository.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindByUserId(ID)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(ID int) (Campaign, error) {
	campaign, err := s.repository.FindById(ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugElement := strings.Replace(input.Name, " ", "-", -1)

	slug := fmt.Sprintf("%d-%s", campaign.UserID, slugElement)
	campaign.Slug = slug

	NewCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return NewCampaign, err
	}
	return NewCampaign, nil
}

func (s *service) UpdateCampaign(input CreateCampaignInput, ID int) (Campaign, error) {
	campaign, err := s.repository.FindById(ID)
	if err != nil {
		return campaign, err
	}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks

	UpdatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return UpdatedCampaign, err
	}

	return campaign, nil
}

func (s *service) CreateImageCampaign(input CreateImageCampaignInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindById(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}
	if input.User.ID != campaign.UserID {
		return CampaignImage{}, errors.New("Unauthorized")
	}
	isPrimary := 0
	if input.IsPrimary {
		_, err := s.repository.MarkAllIsPrimaryFalse(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
		isPrimary = 1
	}

	campaignImage := CampaignImage{
		CampaignID: input.CampaignID,
		FileName:   fileLocation,
		IsPrimary:  isPrimary,
	}

	NewcampaignImage, err := s.repository.SaveCampaignImage(campaignImage)

	if err != nil {
		return NewcampaignImage, err
	}

	return NewcampaignImage, nil

}
