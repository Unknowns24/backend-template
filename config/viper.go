package config

import "github.com/spf13/viper"

var ENV Config

// All app config is stored in this structure
// The values are read by viper from a config file or enviroment variables

type Config struct {
	// General Config
	APP_PORT string `mapstructure:"APP_PORT"`

	// Database Config
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASS"`
	DB_CHARSET  string `mapstructure:"DB_CHAR"`
}

// LoadConfig reads configuration from file or enviroment variables
func Load() {
	cfg := Config{}

	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	ENV = cfg
}
