package yaml

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
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
