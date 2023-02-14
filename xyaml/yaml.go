package xyaml

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// load file to yaml
func LoadFromPath(path string, t interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(b, t); err != nil {
		return err
	}
	return nil
}
