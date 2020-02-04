package main

import (
	"os"
	"strings"
	"time"

	"github.com/FenixAra/prom-discovery/aws"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
	targetFile = os.Getenv("TARGET_FILE")
	targetIP   = os.Getenv("TARGET_IP")
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
			var t []string
			var err error
			if targetIP == strings.ToLower("private") {
				t, err = awsProvider.GetTargets(tg.Cluster, tg.Name, true)
			} else {
				t, err = awsProvider.GetTargets(tg.Cluster, tg.Name, false)
			}
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
