package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ping(c *fiber.Ctx) error {
	msg, err := json.MarshalIndent("Voting Service. Version 0.0.2", "\t", "\t")
	if err != nil {
		log.Panic(err)
	}
	return c.SendString(string(msg))
}

func createPoll(c *fiber.Ctx) error {
	title := c.Query("title")
	options := strings.Split(c.Query("options"), ",")
	poll := creator(title, options)
	res, err := json.MarshalIndent(poll, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	return c.SendString(string(res))
}

func setupRoutes(app *fiber.App) {
	app.Get("/ping", ping)
	app.Post("/createPoll", createPoll)
}

func server() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
