package manifest

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Manifest represets a service manifest
type Manifest struct {
	Name     string            `yaml:"name"`
	Replicas int               `yaml:"replicas"`
	Env      map[string]string `yaml:"env"`
}

// LoadFromFile load manifest file from file path
func LoadFromFile(path string) (*Manifest, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't read yaml file at path=%s, %s", path, err)
	}

	m := Manifest{}
	if err := yaml.Unmarshal(yamlFile, &m); err != nil {
		return nil, fmt.Errorf("can't unmarshal yaml file at path=%s, %s", path, err)
	}

	return &m, nil
}
