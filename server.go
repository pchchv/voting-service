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
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func server() {
	e := echo.New()
	e.GET("/ping", ping)
	e.POST("/createPoll", createPoll)
	log.Fatal(e.Start(":8000"))
}
