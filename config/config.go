package config

import (
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
	//rootPath := GetAbsRootPath() // TODO: Check and remove it completely. As I understand it was necessary for tests
	//absProfilePath := rootPath + "/" + confFileName
	absProfilePath := confFileName

	if err := godotenv.Load(absProfilePath); err != nil {
		log.Fatalf("Error loading %s file: %s", confFileName, err)
	}
	loaded = true
}
