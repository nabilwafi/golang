package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Minute * 5,
		ReadTimeout:  time.Minute * 5,
		WriteTimeout: time.Minute * 5,
		Prefork: true,
	})

	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("Middleware before processing request")
		err := ctx.Next()
		fmt.Println("Middleware after processing request")
		return err
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Wolrd")
	})

	err := app.Listen(":8132")
	if err != nil {
		panic(err)
	}
}
