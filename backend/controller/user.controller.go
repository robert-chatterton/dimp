package controller

import (
	"main/database"
	"main/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user model.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}
