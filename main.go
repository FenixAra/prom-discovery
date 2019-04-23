package main

import (
	"os"
	"time"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
	targetFile = os.Getenv("TARGET_FILE")
)

func main() {
	for {
		updatePromTargets()
		time.Sleep(1 * time.Second)
	}
}

func updatePromTargets() {
	ReadConfig(configFile)
}
