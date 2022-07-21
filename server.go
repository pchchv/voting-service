package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func server() {
	app := fiber.New()
	log.Println("Server started!")
	app.Get("/", func(c *fiber.Ctx) error {
		msg, err := json.MarshalIndent("Voting Service. Version 0.0.1", "\t", "\t")
		if err != nil {
			log.Panic(err)
		}
		return c.SendString(string(msg))
	})
	app.Get("/createPoll", func(c *fiber.Ctx) error {
		title := c.Query("title")
		options := strings.Split(c.Query("options"), ",")
		poll := creator(title, options)
		res, err := json.MarshalIndent(poll, "", "\t")
		if err != nil {
			log.Panic(err)
		}
		return c.SendString(string(res))
	})

	log.Fatal(app.Listen(":8000"))
}
