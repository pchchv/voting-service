package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ping(c echo.Context) error {
	msg := "Voting Service. Version 0.0.2"
	return c.String(http.StatusOK, msg)
}

func createPoll(c echo.Context) error {
	title := c.QueryParam("title")
	options := strings.Split(c.QueryParam("options"), ",")
	v, err := getter("title", title)
	if err != nil {
		poll := creator(title, options)
		res := toDB(poll)
		return c.JSONPretty(http.StatusOK, res, "\t")
	} else {
		res := fmt.Sprintf("The poll duplicates one already created. Existing poll: %v", v)
		return c.JSONPretty(http.StatusOK, res, "\t")
	}
}

func poll(c echo.Context) error {
	var poll Poll
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	option := c.QueryParam("option")
	if title != "" {
		poll = voter("title", title, option)
	} else {
		poll = voter("id", id, option)
	}
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func getPoll(c echo.Context) error {
	var poll *Poll
	var err error
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	if title != "" {
		poll, err = getter("title", title)
		if err != nil {
			log.Panic(err)
		}
	} else {
		poll, err = getter("id", id)
		if err != nil {
			log.Panic(err)
		}
	}
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func deletePoll(c echo.Context) error {
	var poll *Poll
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	if title != "" {
		poll = deleter("title", title)
	} else {
		poll = deleter("id", id)
	}
	return c.JSONPretty(http.StatusOK, poll, "\t")
}

func routes(e *echo.Echo) {
	e.GET("/ping", ping)
	e.POST("/poll", createPoll)
	e.PATCH("/poll", poll)
	e.GET("/poll", getPoll)
	e.DELETE("/poll", deletePoll)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(envURL))
}
