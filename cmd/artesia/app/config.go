package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config contains config data for the server
type Config struct {
	ListenAddress string `mapstructure:"listen_address"`
	Database      DatabaseConfig
}

// DatabaseConfig has the data needed to connect to a Postgres DB
type DatabaseConfig struct {
	Host     string
	DBName   string
	User     string
	Password string
	Port     int
}

// LoadConfig tries to load a config file from the ARTESIA_CONFIG_FILE environment variable. It should be an absolute path
func LoadConfig() (*Config, error) {
	v := viper.New()
	configFilePath := os.Getenv("ARTESIA_CONFIG_FILE")
	if !filepath.IsAbs(configFilePath) {
		return nil, fmt.Errorf("Expected absolute path to config file but got: %s", configFilePath)
	}

	v.SetConfigFile(configFilePath)

	v.AutomaticEnv()

	v.SetDefault("ListenAddress", "0.0.0.0:8080")
	v.SetDefault("Database.Port", 5432)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
