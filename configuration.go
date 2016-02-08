package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type configuration struct {
	TelegramKey string
	Host        string
	DBName      string
	User        string
	Password    string
	SSLMode     string
}

var cfg_instance *configuration
var cfg_once sync.Once

func GetConfigurationInstance() *configuration {
	cfg_once.Do(func() {
		file, _ := os.Open("config.json")

		decoder := json.NewDecoder(file)

		cfg := &configuration{}

		err := decoder.Decode(&cfg)

		if err != nil {
			log.Fatal(err)
		}

		cfg_instance = cfg
	})

	return cfg_instance
}
