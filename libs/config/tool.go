package config

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig(filename string, config interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &config); err != nil {
		return err
	}

	return nil
}
