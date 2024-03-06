package controller

import (
	_ "errors"
	"fmt"
	"net/url"
	"time"

	"github.com/ADEMOLA200/waitlist.git/models"
	"github.com/ADEMOLA200/waitlist.git/services"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ValidateEmail verifies the email using the provided email service
func ValidateEmail(emailService *services.EmailService, email string) error {
	// Perform email validation steps
	if err := emailService.ValidateEmail(email); err != nil {
		// If validation fails, return error response
		return err
	}
	return nil
}

// EmailValidationHandler handles email validation requests
func EmailValidationHandler(emailService *services.EmailService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var req struct {
			Email string `json:"email"`
		}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		// Validate email format
		if err := ValidateEmail(emailService, req.Email); err != nil {
			// If validation fails, return error response
			res := fiber.Map{
				"status":  "error",
				"message": err.Error(),
			}
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		// If validation succeeds, return success response
		res := fiber.Map{
			"status":  "success",
			"message": "Validation successful",
		}
		return c.JSON(res)
	}
}

// WaitlistHandler handles waitlist signup requests
func WaitlistHandler(db *gorm.DB, emailService *services.EmailService) fiber.Handler {


	// Initialize waitlist service
	waitlistService := services.NewWaitlistService(db)

	return func(c *fiber.Ctx) error {
		// Parse request body as JSON
		var req struct {
			Email    string `json:"email"`
			Fullname string `json:"fullname"`
		}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		// Validate email format
		if err := ValidateEmail(emailService, req.Email); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid email format",
			})
		}

		// Check if user is already on waitlist
		_, err := waitlistService.GetWaitlistByEmail(req.Email)
		if err == nil {
			// If the user is already on the waitlist, redirect them to the opt-in link
			optInLink := fmt.Sprintf("https://console.pipeops.io/auth/signup?email=%s", url.QueryEscape(req.Email))
			return c.Redirect(optInLink, fiber.StatusSeeOther)
		}

		// Save new waitlist entry
		if err := waitlistService.AddToWaitlist(req.Email, req.Fullname); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Return 201 created response if there are no errors
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "Your email has been successfully added to the waitlist. We will notify you when space becomes available.",
		})
	}
}

// CampaignHandler handles campaign creation requests
func CampaignHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var req struct {
			UID        string `json:"uid"`
			WaitlistID uint   `json:"waitlist_id"`
			IsRedeemed bool   `json:"is_redeemed"`
		}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		// Fetch Waitlist object
		var waitlist models.Waitlist
		if err := db.First(&waitlist, req.WaitlistID).Error; err != nil {
			return err
		}

		// Create campaign
		campaign := models.Campaign{
			UID:        req.UID,
			Waitlist:   waitlist,
			IsRedeemed: req.IsRedeemed,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		result := db.Create(&campaign)
		if result.Error != nil {
			return result.Error
		}

		// Construct response
		res := fiber.Map{
			"status":  "success",
			"message": "Campaign created successfully",
		}
		return c.JSON(res)
	}
}

// AssignCreditHandler assigns credit to eligible users
func AssignCreditHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Query waitlist for eligible users

		// Assign credit to users

		// Update campaign status

		// Send email notifications to users

		// Construct response
		res := fiber.Map{
			"status":  "success",
			"message": "Credit assigned successfully",
		}
		return c.JSON(res)
	}
}
