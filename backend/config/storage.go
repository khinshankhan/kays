package config

import (
	"fmt"
	"os"
	"strings"
)

type StorageConfig struct {
	Path string `yaml:"path"`
}

func validateStorageConfig(cfg *Config) error {
	if cfg.Storage == nil {
		cfg.Storage = &StorageConfig{
			Path: DefaultStoragePath,
		}
	}

	if len(cfg.Storage.Path) == 0 {
		cfg.Storage.Path = DefaultStoragePath
	}

	info, err := os.Stat(cfg.Storage.Path)
	if os.IsNotExist(err) {
		// if it is a local path, create it
		if !strings.HasPrefix(cfg.Storage.Path, "/") &&
		!strings.HasPrefix(cfg.Storage.Path, "../") &&
		!strings.HasPrefix(cfg.Storage.Path, "~/") {
			err = os.MkdirAll(cfg.Storage.Path, 0700)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid path: \"%s\" does not exist and cannot be created automatically", cfg.Storage.Path)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("invalid path: \"%s\" is not a valid directory", cfg.Storage.Path)
	}

	return nil
}
