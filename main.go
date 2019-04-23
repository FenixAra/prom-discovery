package main

import (
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
	var targets []target
	for _, tg := range conf.Targets {
		switch tg.Provider {
		case ProviderAWS:
			awsProvider := aws.New(tg.Type)
			t, err := awsProvider.GetTargets(tg.Cluster, tg.Name)
			if err != nil {
				continue
			}

			labels := make(map[string]string)
			labels["job"] = tg.Name
			targets = append(targets, target{
				Targets: t,
				Labels:  labels,
			})
		default:
			continue
		}
	}

	writeToTargetFile(targets)
}
