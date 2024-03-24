package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Dev  bool   `mapstructure:"DEV"`
	Port string `mapstructure:"PORT" required:"true"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config file or maybe it was not present: ", err.Error())
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Unable to decode into struct, ", err.Error())
		return nil, err
	}

	return &config, nil
}
