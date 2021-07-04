package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service, campaignService campaign.Service) *transactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) GetTransactionsByCampaignID(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse("Get Campaign Transactions Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)
	input.User = user

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)

	if err != nil {
		response := helper.APIResponse("Get Transaction by Campaign ID Failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormatter := transaction.FormatTransactions(transactions)
	response := helper.APIResponse("Success to Get Transactions by Campaign ID", http.StatusBadRequest, "success", transactionsFormatter)
	c.JSON(http.StatusBadRequest, response)

}
