package xyaml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type config struct {
		App struct {
			Name    string   `yaml:"name" json:"name"`
			Version string   `yaml:"version" json:"version"`
			Address []string `yaml:"address" json:"address"`
		}
	}
	var cfg config
	err := LoadFromPath("testdata/a.yaml", &cfg)
	assert.Nil(t, err)
	fmt.Printf("%+v\n", cfg)
}
