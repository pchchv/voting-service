package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Poll struct {
	Title   string         `json:"title"`
	Options map[string]int `json:"options"`
}

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func creator(title string, options []string) Poll {
	o := make(map[string]int)
	for _, v := range options {
		o[v] = 0
	}
	return Poll{Title: title, Options: o}
}

func main() {
	server()
}
