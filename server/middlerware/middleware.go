package middlerware

import (
	"go-jwt-mk1-showcase/gojwt"

	"github.com/gofiber/fiber/v2"
)

// MIDDLEWARE

func Auth(c *fiber.Ctx) error {
	tokenStr := c.Cookies("authToken")

	err := gojwt.VerifyJWT(tokenStr)
	if err != nil {
		return c.Redirect("/", fiber.StatusSeeOther)
	}
	return c.Next()
}
