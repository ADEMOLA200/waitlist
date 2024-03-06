// services/email.go

package services

import (
    "errors"
    "regexp"
    "strings"

    "github.com/ADEMOLA200/waitlist.git/models"
    "gorm.io/gorm"
)

const maxWaitlistCapacity = 1000 // Adjusted to 1000

type EmailService struct {
    DB *gorm.DB
}

func NewEmailService(db *gorm.DB) *EmailService {
    return &EmailService{DB: db}
}

func (es *EmailService) ValidateEmail(email string) error {
    // Check if email already exists in Waitlist or Campaign
    var (
        waitlistCount int64
        campaignCount int64
    )

    es.DB.Model(&models.Waitlist{}).Where("email = ?", email).Count(&waitlistCount)
    es.DB.Model(&models.Campaign{}).Where("email = ?", email).Count(&campaignCount)

    if waitlistCount+campaignCount > maxWaitlistCapacity {
        return errors.New("waitlist is currently full")
    }

    // Validate email format
    match, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
    if !match {
        return errors.New("invalid email format")
    }

    // Check against disposable email domains
    disposableDomains := []string{"mailinator.com", "guerrillamail.com"}
    for _, domain := range disposableDomains {
        if strings.Contains(email, domain) {
            return errors.New("disposable email not allowed")
        }
    }

    // Check if the email domain is allowed
    allowedDomains := []string{"gmail.com", "email.com", "yahoo.com", "outlook.com", "hotmail.com", "aol.com", "mail.com"}
    domain := strings.Split(email, "@")
    if len(domain) != 2 {
        return errors.New("invalid email format")
    }
    domainAllowed := false
    for _, allowedDomain := range allowedDomains {
        if domain[1] == allowedDomain {
            domainAllowed = true
            break
        }
    }
    if !domainAllowed {
        return errors.New("email address must be from " + strings.Join(allowedDomains, ", "))
    }

    return nil
}