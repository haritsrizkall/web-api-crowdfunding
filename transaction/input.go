package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type GetUserTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
