// routes/routes.go
package routes

import (
    "github.com/ADEMOLA200/waitlist.git/controller"
    "github.com/ADEMOLA200/waitlist.git/services"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// SetupRoutes defines the application routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
    // Initialize email service
    emailService := services.NewEmailService(db)

    // Email validation route
    app.Post("/email-validation", controller.EmailValidationHandler(emailService))

    // Waitlist route
    app.Post("/waitlist", controller.WaitlistHandler(db, emailService))

    // Campaigns route
    app.Post("/campaigns", controller.CampaignHandler(db))

    // Assign credit route
    app.Post("/assign-credit", controller.AssignCreditHandler(db))
}
