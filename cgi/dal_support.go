package cgi

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/dal"
	"murphyl.com/lego/pkg/shared"
)

func DefaultDataAccessLayer(c fiber.Ctx) dal.DataAccessLayer {
	return GetDataAccessLayer(c, shared.DeafultKey)
}

func GetDataAccessLayer(c fiber.Ctx, key string) dal.DataAccessLayer {
	if dao, ok := fiber.GetService[dal.DataAccessLayer](c.App().State(), dal.RefKey(key)); ok {
		return dao
	}
	return nil
}
