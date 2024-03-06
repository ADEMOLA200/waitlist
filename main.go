package main

import (
	"github.com/ADEMOLA200/waitlist.git/database"
	"github.com/ADEMOLA200/waitlist.git/routes"
	
	"github.com/ADEMOLA200/waitlist.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Connect to database
    db := database.ConnectDB()
	
    // Initialize Fiber
    app := fiber.New()

	 // Register middleware
	 app.Use(middlewares.ErrorHandler())
	 app.Use(middlewares.Logger())

    
    // Migrate the database models
    database.MigrateDB(db)

    // Define routes
    routes.SetupRoutes(app, db)

    // Start server
    app.Listen(":5000")
}
