// middlewares/error_handler.go

package middlewares

import (
    "github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
    return func(c *fiber.Ctx) error {
        if err := c.Next(); err != nil {
            // Handle errors
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return nil
    }
}
