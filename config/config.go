package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config structure for database
type Config struct {
	Database struct {
		Server   string
		Port     int
		User     string
		Password string
		Name     string
	}
}

// LoadConfig loads configuration from config.yaml
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")                                            // Config file name (without extension)
	viper.SetConfigType("yaml")                                              // YAML format
	viper.AddConfigPath("C:/Users/kahwai.voon/Desktop/Go_API/GO_API/config") // Absolute path
	viper.AddConfigPath("./config")                                          // Relative path (if running from project root)

	// Read in environment variables, allowing overrides
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	log.Println("✅ Config loaded successfully")
	return &config, nil
}
