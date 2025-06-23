package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ResendAPIKey string `mapstructure:"RESEND_API_KEY"`
	FromName     string `mapstructure:"FROM_NAME"`
	FromEmail    string `mapstructure:"FROM_EMAIL"`
	Region       string `mapstructure:"AWS_REGION"`    
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUser     string `mapstructure:"SMTP_USER"`
	SMTPPass     string `mapstructure:"SMTP_PASSWORD"`
	 
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
   
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}


	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config, nil
}