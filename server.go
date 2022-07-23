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
	v := fromDB(title)
	if v == "id not found" {
		poll := creator(title, options)
		res := toDB(poll)
		return c.JSONPretty(http.StatusOK, res, "\t")
	} else {
		v = "The poll duplicates one already created. Id of an existing poll: " + v
		return c.JSONPretty(http.StatusOK, v, "\t")
	}
}

func poll(c echo.Context) error {
	title := c.QueryParam("title")
	option := c.QueryParam("options")
	poll := voter(title, option)
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func getPoll(c echo.Context) error {
	var poll ResultPoll
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	if title != "" {
		poll = getter("title", title)
	} else {
		poll = getter("id", id)
	}
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func deletePoll(c echo.Context) error {
	var poll ResultPoll
	// TODO: Should return a remote poll in JSON format
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	if title != "" {
		poll = deleter("title", title)
	} else {
		poll = deleter("id", id)
	}
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func server() {
	e := echo.New()
	e.GET("/ping", ping)
	e.POST("/createpoll", createPoll)
	e.PATCH("/poll", poll)
	e.GET("/getpoll", getPoll)
	e.DELETE("/deletepoll", deletePoll)
	log.Fatal(e.Start(envURL))
}
