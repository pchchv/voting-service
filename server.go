package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	msg := "Voting Service. Version 0.0.2"
	return c.String(http.StatusOK, msg)
}

func createPoll(c echo.Context) error {
	title := c.QueryParam("title")
	options := strings.Split(c.QueryParam("options"), ",")
	poll := creator(title, options)
	id := toDB(poll)
	log.Panicln(id)
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func deletePoll(c echo.Context) error {
	// TODO: Should return a remote poll in JSON format
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	if title != "" {
		deleter("title", title)
	} else {
		deleter("id", id)
	}
	return c.String(http.StatusOK, "Poll deleted")
}

func server() {
	e := echo.New()
	e.GET("/ping", ping)
	e.POST("/createPoll", createPoll)
	e.DELETE("/deletePoll", deletePoll)
	log.Fatal(e.Start(envURL))
}
