package config

import (
	"log"

	"github.com/joho/godotenv"
)

var loaded = false

func LoadConfig(fileName string) {
	if loaded {
		return
	}

	if err := godotenv.Load(fileName); err != nil {
		log.Fatalf("Error loading %s file: %s", fileName, err)
	}
	loaded = true
}
