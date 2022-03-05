package get

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func (t Tools) Len() int { return len(t) }

func (t Tools) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t Tools) Less(i, j int) bool {
	var ti = t[i]
	var tj = t[j]
	var tiNameLower = strings.ToLower(ti.Name)
	var tjNameLower = strings.ToLower(tj.Name)
	if tiNameLower == tjNameLower {
		return ti.Name < tj.Name
	}
	return tiNameLower < tjNameLower
}

// Config struct for webapp config
type Config struct {
	Tools []Tool `yaml:"tools"`
}

func ReadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

type Tools []Tool

func MakeTools() Tools {

	tools, err := ReadConfig("tools.yml")
	if err != nil {
		fmt.Println("Error reading tools config - " + err.Error())
	}

	return tools.Tools
}
