package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/sys/unix"
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

	clean := path.Clean(cfg.Storage.Path)
	if strings.HasPrefix(clean, "~/") {
		return fmt.Errorf("use $HOME instead of ~ in your path")
	}

	info, err := os.Stat(cfg.Storage.Path)
	if os.IsNotExist(err) {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not retrieve current working directory for error checking")
		}
		absValid := path.IsAbs(clean) && strings.HasPrefix(clean, cwd)
		relValid := !path.IsAbs(clean) && !strings.HasPrefix(clean, "..")

		// if it is a local path, create it
		if absValid || relValid {
			err = os.MkdirAll(clean, 0700)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid path: \"%s\" does not exist and cannot be created automatically", cfg.Storage.Path)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("invalid path: \"%s\" is not a valid directory", cfg.Storage.Path)
	} else if unix.Access(cfg.Storage.Path, unix.W_OK) != nil {
		return fmt.Errorf("invalid path: user does not have write permission on directory \"%s\"", cfg.Storage.Path)
	}

	return nil
}
