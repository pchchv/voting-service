package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var envURL string

type Poll struct {
	Title   string         `json:"title"`
	Options map[string]int `json:"options"`
}

type ResultPoll struct {
	Id   string `json:"id"`
	Poll Poll
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

func toDB(poll Poll) ResultPoll {
	v, err := bson.Marshal(poll)
	if err != nil {
		log.Panic(err)
	}
	result, err := collection.InsertOne(context.TODO(), v)
	if err != nil {
		log.Panic(err)
	}
	return ResultPoll{fmt.Sprint(result.InsertedID), poll}
}

func deleter(t string, v string) {
	// TODO: Should delete the poll from mongo and return it
	if t == "title" {
		fmt.Println(v)
	} else if t == "id" {
		fmt.Println(v)
	}
}
func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	db()
	server()
}
