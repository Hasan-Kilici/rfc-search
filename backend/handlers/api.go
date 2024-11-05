package handlers

import (
	"rfc/utils"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	name := c.Query("title")
	rfcs := utils.Search(name)

	return c.JSON(fiber.Map{
		"data": rfcs,
	})
}
