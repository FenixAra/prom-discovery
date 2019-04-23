package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type target struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func writeToTargetFile(targets []target) {
	data, err := json.Marshal(targets)
	if err != nil {
		log.Println("Unable to marshal targets. Err: ", err)
		return
	}

	if targetFile == "" {
		targetFile = "targets.json"
	}

	err = ioutil.WriteFile(targetFile, data, 0644)
	if err != nil {
		log.Println("Unable to save file. Err: ", err)
	}
}
