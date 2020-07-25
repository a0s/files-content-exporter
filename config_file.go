package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	Entities           []ConfigFileEntity `yaml:"entities,flow"`
	PathAsLabelEnabled bool               `yaml:"path_as_label_enabled,flow"`
}

type ConfigFileEntity struct {
	File   string           `yaml:"file,flow"`
	Name   string           `yaml:"name,flow"`
	Labels ConfigFileLabels `yaml:"labels,flow"`
}

type ConfigFileLabels map[string]string

func ReadConfigFile(configFilePath string) *ConfigFile {
	if strings.TrimSpace(configFilePath) == "" {
		log.Fatalf("Empty yaml config path")
	}

	filePath, err := filepath.Abs(configFilePath)
	if err != nil {
		log.Fatalf("Cannot normalize yaml config path")
	}

	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Cannot read yaml config: %v", err)
	}

	var configFile ConfigFile
	err = yaml.Unmarshal(yamlContent, &configFile)
	if err != nil {
		log.Fatalf("Cannot parse yaml config: %v", err)
	}

	return &configFile
}
