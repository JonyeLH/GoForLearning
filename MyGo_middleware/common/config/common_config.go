package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var BathPath string

func InitConfigs() {
	BathPath = "conf/"
}

func initConf(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	fmt.Println(string(content))
	defer func() {
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
	}()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return err
	}
	return nil
}
