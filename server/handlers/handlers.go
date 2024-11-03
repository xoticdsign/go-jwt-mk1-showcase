package handlers

import (
	"go-jwt-mk1-showcase/gojwt"

	"time"

	"github.com/gofiber/fiber/v2"
)

// VARIABLES

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

// SUPPORTING FUNCTIONS

func imaginaryDB(username string, password string) bool {
	val, ok := Users[username]
	if !ok {
		return false
	}

	if val != password {
		return false
	}
	return true
}

// HANDLERS

func Error(c *fiber.Ctx, err error) error {
	switch err {
	case fiber.ErrNotFound:
		return c.SendString("Sorry, this page doesn't exist: 404")

	default:
		return c.SendString("Something bad happened on our side: 500")
	}
}

func Root(c *fiber.Ctx) error {
	err := c.Render("index", fiber.Map{})
	if err != nil {
		return fiber.ErrNotFound
	}
	return nil
}

func Submit(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	ok := imaginaryDB(username, password)
	if !ok {
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	tokenStr, err := gojwt.ConfigJWT(username)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "authToken"
	cookie.Value = tokenStr
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour)
	cookie.Secure = true
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.Redirect("/secret-page", fiber.StatusSeeOther)
}

func SecretPage(c *fiber.Ctx) error {
	tokenStr := c.Cookies("authToken")

	return c.SendString("Your token string: " + tokenStr)
}
