package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var envURL string
var testURL string

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

func voter(title string, value string, option string) Poll {
	poll, err := getter(title, value)
	if err != nil {
		log.Panic(err)
	}
	t := poll.Title
	o := poll.Options
	o[option] += 1
	p := Poll{t, o}
	filter := bson.D{{Key: title, Value: value}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: t}, {Key: "options", Value: o}}}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Panic(err)
	}
	return p
}

func getter(title string, value string) (*Poll, error) {
	var res *mongo.SingleResult
	result := &Poll{}
	if title == "title" {
		res = collection.FindOne(context.TODO(), bson.M{"title": value})
	} else if title == "id" {
		res = collection.FindOne(context.TODO(), bson.M{"_id": value})
	}
	err := res.Decode(result)
	if err != nil {
		return result, errors.New("Poll not found")
	}
	return result, nil
}

func deleter(title string, value string) *Poll {
	var res *mongo.SingleResult
	result := &Poll{}
	if title == "title" {
		res = collection.FindOneAndDelete(context.TODO(), bson.M{"title": value})
	} else if title == "id" {
		res = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": value})
	}
	err := res.Decode(result)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	db()
	server()
}
