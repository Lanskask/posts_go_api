package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type SystemConfig struct {
	DB     DBType     `yaml:"db"`
	Router RouterType `yaml:"router"`
}

func GetSystemConfig(file string) (SystemConfig, error) {
	conf, err := GetConfigFromFile[SystemConfig](file)
	return *conf, err
}

func GetConfigFromFile[T any](filename string) (*T, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading the file %q: %s", filename, err)
	}

	c := new(T)

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("error marshalling a config file %q: %v", filename, err)
	}

	return c, nil
}

type DBType string

const (
	SQLITE   DBType = "sqlite"
	FIREBASE        = "firebase"
	MEM             = "mem"
)

type RouterType string

const (
	CHI   RouterType = "chi"
	FIBER            = "fiber"
	GIN              = "gin"
	MUX              = "mux"
)

var Databases = []DBType{SQLITE, FIREBASE, MEM}
var Routers = []RouterType{CHI, FIBER, GIN, MUX}
