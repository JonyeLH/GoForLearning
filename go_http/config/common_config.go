package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)
var BathPath string

func InitConfigs(){
	BathPath = "conf/"
}

func initConf(path string, config interface{}) error{
	content, err := ioutil.ReadFile(path)
	defer func() {
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
	}()
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, config)
	if err != nil {
		return err
	}
	return nil
}
