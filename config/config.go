package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	Environment  string `mapstructure:"ENV"`
	ReadTimeout  string `mapstructure:"READ_TIMEOUT"`
	WriteTimeout string `mapstructure:"WRITE_TIMEOUT"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigType("env")

	// Menentukan file konfigurasi berdasarkan environment
	env := viper.GetString("ENV")
	if env == "" {
		env = "dev" // Default ke development
	}
	viper.SetConfigName(env)
	viper.AddConfigPath("./config")

	// Membaca file konfigurasi
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Parsing konfigurasi ke struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}
}
