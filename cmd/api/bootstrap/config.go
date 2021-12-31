package bootstrap

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"products/model"
)

func newConfiguration(path string) model.Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := model.Configuration{}
	if err := json.Unmarshal(file, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
