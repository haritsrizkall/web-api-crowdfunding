package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func CampaignFormatTransaction(transaction Transaction) CampaignTransactionFormatter {
	transactionFormatter := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return transactionFormatter
}

func CampaignFormatTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		transactionFormatter := CampaignFormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}
	return transactionsFormatter
}

func UserFormatTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign: CampaignFormatter{
			Name:     transaction.Campaign.Name,
			ImageURL: transaction.Campaign.CampaignImages[0].FileName,
		},
	}

	return userTransactionFormatter
}

func UserFormatTransactions(transactions []Transaction) []UserTransactionFormatter {
	var userTransactionsFormatter []UserTransactionFormatter
	for _, transaction := range transactions {
		userTransactionsFormatter = append(userTransactionsFormatter, UserFormatTransaction(transaction))
	}
	return userTransactionsFormatter
}
