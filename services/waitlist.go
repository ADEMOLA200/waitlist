// services/waitlist.go

package services

import (
    "github.com/ADEMOLA200/waitlist.git/models"
    "gorm.io/gorm"
    "time"
)

type WaitlistService struct {
    DB *gorm.DB
}

func NewWaitlistService(db *gorm.DB) *WaitlistService {
    return &WaitlistService{DB: db}
}


// The key steps are:
// 1. Create a new Waitlist model struct with the email and timestamps
// 2. Save the waitlist entry to the database using ws.DB. Create(&waitlist)
// 3. Handle any errors from the database insert
// AddToWaitlist adds a new entry to the waitlist
func (ws *WaitlistService) AddToWaitlist(email, fullname string) error {
    // Create a new Waitlist entry
    waitlist := models.Waitlist{
        Email:     email,
        FullName:  fullname,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Save to DB
    result := ws.DB.Create(&waitlist)
    if result.Error != nil {
        return result.Error
    }

    return nil
}



// Get all users in the waitlist
func (s *WaitlistService) GetWaitlist() ([]models.Waitlist, error) {
    var waitlists []models.Waitlist

    result := s.DB.Find(&waitlists)
    if result.Error != nil {
        return nil, result.Error
    }

    return waitlists, nil
}

// GetWaitlistByEmail retrieves a waitlist entry by email
func (ws *WaitlistService) GetWaitlistByEmail(email string) (*models.Waitlist, error) {
    var waitlist models.Waitlist
    result := ws.DB.Where("email = ?", email).First(&waitlist)
    if result.Error != nil {
        return nil, result.Error
    }
    return &waitlist, nil
}


// Remove a user from the waitlist by their ID
func (s *WaitlistService) RemoveFromWaitlist(id uint) error {
    waitlist := models.Waitlist{
        ID: id,
    }

    result := s.DB.Delete(&waitlist)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

