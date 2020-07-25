package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type yamlConfig struct {
	Entities           []entity `yaml:"entities,flow"`
	PathAsLabelEnabled bool     `yaml:"path_as_label_enabled,flow"`
}

type entity struct {
	File   string            `yaml:"file,flow"`
	Name   string            `yaml:"name,flow"`
	Labels map[string]string `yaml:"labels,flow"`
	Help   string            `yaml:"help,flow"`
}

func readYamlConfig(filePath string) *yamlConfig {
	if strings.TrimSpace(filePath) == "" {
		log.Fatalf("empty yaml config path")
	}

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatalf("can't normalize yaml config path")
	}

	yamlContent, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		log.Fatalf("can't read yaml config: %v", err)
	}

	var configFile yamlConfig
	err = yaml.Unmarshal(yamlContent, &configFile)
	if err != nil {
		log.Fatalf("can't parse yaml config: %v", err)
	}

	return &configFile
}
