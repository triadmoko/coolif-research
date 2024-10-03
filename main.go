package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Membuat instance Fiber baru
	app := fiber.New()

	// Route dasar
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Selamat datang di homepage")
	})

	// Route dengan parameter
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Halo, " + name)
	})

	// Route dengan middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "Nilai Header Kustom")
		return c.Next()
	})

	// Grouping routes
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Daftar pengguna dari API v1",
		})
	})

	v2 := api.Group("/v2")
	v2.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Daftar pengguna dari API v2",
		})
	})

	// Menjalankan server pada port 3000
	app.Listen(":3000")
}
