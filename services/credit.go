// services/credit.go

package services

import (
	_ "errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/ADEMOLA200/waitlist.git/models"
	"github.com/ADEMOLA200/waitlist.git/templates"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditService struct {
    DB *gorm.DB
}

func NewCreditService(db *gorm.DB) *CreditService {
    return &CreditService{DB: db}
}

// AssignCredit assigns credit to eligible users
func (cs *CreditService) AssignCredit() error {
    // Query waitlist for eligible users
    var waitlistUsers []models.Waitlist
    if err := cs.DB.Find(&waitlistUsers).Error; err != nil {
        return err
    }

    // Assign credit to users
    for _, user := range waitlistUsers {
        // Assuming some criteria for assigning credit (e.g., every user gets $150 credits)
        creditAmount := 150

        // Create a new campaign for each user
        campaign := models.Campaign{
            UID:        generateUniqueID(), // Assuming there's a function to generate a unique ID
            Waitlist:   models.Waitlist{},              // Assign the Waitlist struct directly
            IsRedeemed: false,
            CreatedAt:  time.Now(),
            UpdatedAt:  time.Now(),
        }

        // Save the campaign to the database
        if err := cs.DB.Create(&campaign).Error; err != nil {
            return err
        }

        // Send email notification to the user
        // This could be a separate function to handle sending emails
        if err := sendEmailNotification(user, creditAmount); err != nil {
            return err
        }

        // Optionally, update the waitlist status or remove the user from the waitlist
        // Assuming there's a function to update the waitlist status or remove the user
        if err := updateUserStatus(user.ID, "credited"); err != nil {
            return err
        }
    }

    return nil
}

// generateUniqueID generates a unique ID for the campaign
func generateUniqueID() string {

	// Generate a UUID
	uuid := uuid.New()
  
	// Optionally add timestamp for ordering
	timestamp := time.Now().Format("20060102150405")
  
	return fmt.Sprintf("campaign_%s_%s", uuid, timestamp)
}

// sendEmailNotification sends email to user
func sendEmailNotification(waitlist models.Waitlist, amount int) error {
    // Render email template
    body, err := templates.RenderCreditNotification(waitlist.Email, amount)
    if err != nil {
        return err
    }

    // Your email sending mechanism here
    // For example, if using SMTP to send email
    smtpAddress := "your_smtp_address:port"
    smtpAuth := smtp.PlainAuth("", "username", "password", "your_smtp_host")
    to := []string{waitlist.Email}
    msg := []byte("To: " + waitlist.Email + "\r\n" +
        "Subject: Your Credits\r\n" +
        "\r\n" +
        body)

    err = smtp.SendMail(smtpAddress, smtpAuth, "sender_email", to, msg)
    if err != nil {
        return err
    }

    return nil
}


// updateUserStatus updates the status of the user in the waitlist
func updateUserStatus(userID uint, status string) error {
    // Implementation to update the status of the user
    // This is just a placeholder implementation
    fmt.Printf("Updating status of user with ID %d to %s\n", userID, status)
    return nil
}