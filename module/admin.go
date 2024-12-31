package module

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Proyek-Three/bp-promosi-umkm/controller" 
)

func main() {
    r := fiber.New()
    r.Use(logger.New())

    // Route login
    r.Post("/login", controller.LoginAdmin)

    r.Listen(":8080")
}
