package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var db Config

func dbInit() error {
	content, err := os.ReadFile("./db.json")
	if err != nil {
		fmt.Println("Error while reading JSON file:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(content, &db)
	if err != nil {
		fmt.Println("Error while decoding JSON:", err)
		os.Exit(1)
	}

	fmt.Printf("data from db : %v", db)

	return nil

}

func dbhandler(data interface{}, from string) {

	if from == "message" {

		smessage = &sMessage{
			Message: "not found",
		}

		q := data.(rMessage)
		fmt.Println(q.Query)

		var len_res int
		var randIndx int

		for _, i := range db.Database.Intents {
			for _, p := range i.Patterns {
				if p == q.Query {
					len_res = len(i.Responses)
					randIndx = rand.Intn(len_res)

					r := i.Responses[randIndx]
					smessage = &sMessage{
						Message: r,
					}
				}

			}
		}
	}

	if from == "train" {
		d := data.(train)

		var t_data intents
		p := strings.Split(d.Patterns, "||")
		r := strings.Split(d.Responses, "||")
		t_data = intents{
			Tag:       d.Tag,
			Patterns:  p,
			Responses: r,
		}

		if len(d.Tag) > 0 && len(d.Patterns) > 0 && len(d.Responses) > 0 {

			newIn := []intents{t_data}
			newIn = append(newIn, db.Database.Intents...)
			db.Database.Intents = newIn

			db, _ := json.Marshal(db)
			os.WriteFile("./db.json", db, 0064)
		}
		return
	}

	if from == "botTrain" {
		d := data.(sbt)

		bt_data := intents{
			Tag:       "ai",
			Patterns:  []string{d.Query},
			Responses: []string{d.Mesg},
		}

		db.Database.Intents = append(db.Database.Intents, bt_data)
		db, _ := json.Marshal(db)
		os.WriteFile("./db.json", db, 0064)

	}

}
