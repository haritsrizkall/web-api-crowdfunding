package campaign

import (
	"bwastartup/user"
)

type CreateCampaignInput struct {
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	Perks            string
	User             user.User
}
