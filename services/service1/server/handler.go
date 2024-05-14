package server

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
)

func GetIndexHandler(c *fiber.Ctx) error {
	tmpl := template.Must(template.ParseFiles("public/html/index.html"))
	err := tmpl.Execute(c.Response().BodyWriter(), nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}
