package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
)


func UnslashURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.HasSuffix(c.OriginalURL(), "/") && len(c.OriginalURL()) > 1 {
			return c.Redirect(strings.TrimSuffix(c.OriginalURL(), "/"), 301)
		}
		return c.Next()
	}
}