package main

type patterns []string
type responses []string

type intents struct {
	Tag       string    `json:"tag"`
	Patterns  patterns  `json:"patterns"`
	Responses responses `json:"responses"`
}

type Database struct {
	Intents []intents `json:"intents"`
}
type Config struct {
	Database Database `json:"database"`
}
