package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func GetTransfers(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	transfers, err := service.GetUserTransfers(claims)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.TransfersNotFound,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"transfers": transfers,
	})
}
