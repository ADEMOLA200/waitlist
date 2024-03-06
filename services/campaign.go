// services/campaign.go

package services

import (
	"time"

	"github.com/ADEMOLA200/waitlist.git/models"
	"gorm.io/gorm"
)

type CampaignService struct {
    DB *gorm.DB
}

func NewCampaignService(db *gorm.DB) *CampaignService {
    return &CampaignService{DB: db}
}

// CreateCampaign creates a new campaign
func (cs *CampaignService) CreateCampaign(uid string, waitlistID uint, isRedeemed bool) error {
    // Fetch Waitlist object
    var waitlist models.Waitlist
    if err := cs.DB.First(&waitlist, waitlistID).Error; err != nil {
        return err
    }

    // Create a new campaign
    campaign := models.Campaign{
        UID:        uid,
        Waitlist:   waitlist,
        IsRedeemed: isRedeemed,
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
    }

    // Save to DB
    result := cs.DB.Create(&campaign)
    if result.Error != nil {
        return result.Error
    }

    return nil
}