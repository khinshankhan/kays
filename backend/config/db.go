package config

import (
	"fmt"
	"math"
)

// DBConfig is the structure which supports the configuration of the database
type DBConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Name     string `yaml:"dbname" json:"dbname"`
}

func validateDBConfig(cfg *Config) error {
	if cfg.DB == nil {
		cfg.DB = &DBConfig{
			Host:     DefaultDBHost,
			Port:     DefaultDBPort,
			User:     DefaultDBUser,
			Password: DefaultDBPassword,
			Name:     DefaultDBName,
		}
		return nil
	}

	if len(cfg.DB.Host) == 0 {
		cfg.DB.Host = DefaultDBHost
	}

	if cfg.DB.Port == 0 {
		cfg.DB.Port = DefaultDBPort
	} else if cfg.DB.Port < 0 || cfg.DB.Port > math.MaxUint16 {
		return fmt.Errorf("invalid port: value should be between %d and %d", 0, math.MaxUint16)
	}

	if len(cfg.DB.User) == 0 {
		cfg.DB.User = DefaultDBUser
	}

	if len(cfg.DB.Password) == 0 {
		cfg.DB.Password = DefaultDBPassword
	}

	if len(cfg.DB.Name) == 0 {
		cfg.DB.Name = DefaultDBName
	}

	return nil
}

func (db *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.Name,
	)
}
