package transaction

import "time"

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return transactionFormatter
}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {
	var transactionsFormatter []TransactionFormatter

	for _, transaction := range transactions {
		transactionFormatter := FormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}
	return transactionsFormatter
}
