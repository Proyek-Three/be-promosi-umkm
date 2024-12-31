package module

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger" 
)

func main() {
    r := fiber.New()
    r.Use(logger.New())
    r.Listen(":8080")
}
