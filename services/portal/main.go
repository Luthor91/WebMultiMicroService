package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	// Initialisation de l'application Fiber
	app := fiber.New()

	// Route pour le service1
	app.Use("/service1/*", func(c *fiber.Ctx) error {
		target := "http://127.0.0.1:8081" + strings.TrimPrefix(c.OriginalURL(), "/service1")
		return proxy.Do(c, target)
	})

	// Route de fallback pour les URL non trouvées
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Page non trouvée")
	})

	// Démarrer le serveur principal
	log.Fatal(app.Listen("127.0.0.1:3000"))
}
