package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

const (
	confFileName = ".profile"
)

var loaded = false

func LoadConfig() {
	if loaded {
		return
	}
	err := godotenv.Load(confFileName)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error loading %s file", confFileName))
	}
	loaded = true
}
