package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlInfo struct {
	Name      string   `yaml:"name"`
	Countries []string `yaml:"countries"`
}

type config struct {
	Path     string
	Yamls    *[]YamlInfo
	BasePath string
}

func Newconfig(path string) *config {
	return &config{
		Path: path,
	}
}

func (c *config) GetConf() error {

	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		return fmt.Errorf("yamlFile.Get err   #%v ", err)
	}

	var configs []YamlInfo
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	c.Yamls = &configs

	return nil
}

func (c *config) Initial() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	c.BasePath = dir

	continentsDir := fmt.Sprintf("%s/%s", dir, "continents")
	os.RemoveAll(continentsDir)

	if err := EnsureDir(continentsDir); err != nil {
		log.Fatal(err)
	}

}

func EnsureDir(name string) error {
	_, err := os.Open(name)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(name, os.FileMode(0760))
		if err != nil {
			return err
		}
	}
	return nil
}
