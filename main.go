package main

import (
	"log"
	"os"
	"time"

	"github.com/FenixAra/prom-discovery/aws"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
	targetFile = os.Getenv("TARGET_FILE")
)

func main() {
	for {
		updatePromTargets()
		time.Sleep(15 * time.Second)
	}
}

func updatePromTargets() {
	ReadConfig(configFile)

	for _, target := range conf.Targets {
		switch target.Provider {
		case ProviderAWS:
			awsProvider := aws.New(target.Type)
			targets, err := awsProvider.GetTargets(target.Cluster, target.Name)
			if err != nil {
				continue
			}

			log.Println(targets)
		default:
			continue
		}
	}
}
