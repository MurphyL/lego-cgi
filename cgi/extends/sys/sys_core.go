package sys

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/app/sys/handlers"
)

func UseSystemDictManager(router fiber.Router) {
	router.Get("/dict/type", handlers.GetDictTypeHandler)
	router.Get("/dict/items", handlers.GetDictTypeHandler)
}
