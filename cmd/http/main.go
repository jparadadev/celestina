package main

import (
	"celestina/cmd/http/bootstrap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	rawConfigData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var configData map[string]interface{}
	err = yaml.Unmarshal(rawConfigData, &configData)
	if err != nil {
		panic(err)
	}

	rawSubsData, err := ioutil.ReadFile("subscriptions.yaml")
	if err != nil {
		panic(err)
	}
	var subsData map[string][]string
	err = yaml.Unmarshal(rawSubsData, &subsData)
	if err != nil {
		panic(err)
	}

	if err := bootstrap.Run(configData, subsData); err != nil {
		log.Fatal(err)
	}
}
