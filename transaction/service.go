package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(input GetUserTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	var transactions []Transaction
	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return transactions, err
	}
	if campaign.UserID != input.User.ID {
		return transactions, errors.New("Unauthorized")
	}
	transactions, err = s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	if len(transactions) == 0 {
		return transactions, errors.New("Transaction not found by that campaign id")
	}
	return transactions, nil
}

func (s *service) GetTransactionsByUserID(input GetUserTransactionsInput) ([]Transaction, error) {
	if input.User.ID != input.ID {
		return []Transaction{}, errors.New("Unauthorized")
	}

	transactions, err := s.repository.GetByUserID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
