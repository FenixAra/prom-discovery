package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Targets []Target `json:"targets"`
}

type Target struct {
	Name     string `json:"name"`
	Cluster  string `json:"cluster"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
	Port     string `json:"port"`
}

var conf *Config

func init() {
	conf = &Config{}
}

func ReadConfig(fileName string) {
	if fileName == "" {
		fileName = "config.json"
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Unable to read config file. Err:", err)
		return
	}

	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Println("Unable to unmarshal config file. Err:", err)
		return
	}
}
