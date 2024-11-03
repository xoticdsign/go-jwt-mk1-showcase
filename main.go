package main

import (
	"go-jwt-mk1-showcase/server/handlers"
	"go-jwt-mk1-showcase/server/middlerware"

	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// ENTRY POINT

func main() {
	fmt.Println("Test users:")
	for key, val := range handlers.Users {
		fmt.Println(key, val)
	}

	app := fiber.New(fiber.Config{
		Views:        html.New("./server/templates", ".html"),
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		ErrorHandler: handlers.Error,
		AppName:      "go-jwt-mk1-showcase",
	})

	app.Use("/secret-page", middlerware.Auth)

	app.Get("/", handlers.Root)
	app.Post("/submit", handlers.Submit)
	app.Get("/secret-page", handlers.SecretPage)

	err := app.Listen("0.0.0.0:4032")
	if err != nil {
		log.Fatal(err)
	}
}
