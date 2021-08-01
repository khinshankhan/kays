package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the main configuration structure
type Config struct {
	// Web is the configuration for the web listener
	Web *WebConfig `yaml:"web"`

	// DB is the configuration for the db connection
	DB *DBConfig `yaml:"db"`

	// Meta contains meta information about the config
	// it isn't specified in the user config
	Meta *MetaConfig
}

// Load loads a custom configuration file
func Load(commitHash, buildDate, configFile string) (*Config, error) {
	log.Printf("Reading configuration from configFile=%s\n", configFile)

	cfg, err := readConfigurationFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("configuration file not found")
		}

		return nil, err
	}

	metaInfo := &MetaConfig{
		CommitHash: commitHash,
		BuildDate:  buildDate,
		filePath:   configFile,
	}
	cfg.Meta = metaInfo

	return cfg, nil
}

func readConfigurationFile(filename string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// expand environment variables before parsing them
	bytes = []byte(os.ExpandEnv(string(bytes)))

	return parseAndValidateConfig(bytes, filename)
}

func parseAndValidateConfig(bytes []byte, filename string) (*Config, error) {
	config := &Config{}
	yaml.Unmarshal(bytes, config)

	if err := validateWebConfig(config); err != nil {
		return nil, err
	}

	if err := validateDBConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}
