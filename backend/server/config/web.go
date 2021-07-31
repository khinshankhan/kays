package config

import (
	"fmt"
	"math"
)

// WebConfig is the structure which supports the configuration of the endpoint
// which provides access to the web frontend
type WebConfig struct {
	// Address to listen on (defaults to 0.0.0.0 specified by DefaultWebAddress)
	Address string `yaml:"address" json:"address"`

	// Port to listen on (default to 8080 specified by DefaultWebPort)
	Port int `yaml:"port" json:"port"`
}

func validateWebConfig(cfg *Config) error {
	if cfg.Web == nil {
		cfg.Web = &WebConfig{Address: DefaultWebAddress, Port: DefaultWebPort}
		return nil
	}

	if len(cfg.Web.Address) == 0 {
		cfg.Web.Address = DefaultWebAddress
	}

	if cfg.Web.Port == 0 {
		cfg.Web.Port = DefaultWebPort
	} else if cfg.Web.Port < 0 || cfg.Web.Port > math.MaxUint16 {
		return fmt.Errorf("invalid port: value should be between %d and %d", 0, math.MaxUint16)
	}

	return nil
}

// SocketAddress returns the combination of the Address and the Port
func (web *WebConfig) SocketAddress() string {
	return fmt.Sprintf("%s:%d", web.Address, web.Port)
}
