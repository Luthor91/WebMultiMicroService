package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartWebsite initialise le serveur web avec Fiber
func DefineRoutes() {
	// Initialisation de l'application Fiber
	app := fiber.New()

	app.Static("/", "public")

	// Gestionnaire pour l'URL racine "/"
	app.Get("/index", GetIndexHandler)

	// Route de fallback pour les URL non trouvées, doit se trouver tout à la fin du code
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Page non trouvée")
	})

	log.Fatal(app.Listen("127.0.0.1:8081"))
}