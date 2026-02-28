package cgi

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/dal"
	"murphyl.com/lego/udf"
)

func DefaultDataAccessLayer(c fiber.Ctx) dal.DataAccessLayer {
	return GetDataAccessLayer(c, udf.DeafultKey)
}

func GetDataAccessLayer(c fiber.Ctx, key string) dal.DataAccessLayer {
	if dao, ok := fiber.GetService[dal.DataAccessLayer](c.App().State(), dal.RefKey(key)); ok {
		return dao
	}
	return nil
}
