package config

import "github.com/spf13/viper"

type Config struct {
	PORT string
	NEW_DB bool
	Data struct {
		LISTS       []string
	}
}

func initViper(fileName string) (Config, error) {
	viper.SetConfigName(fileName)	// config fileName without extension
	viper.AddConfigPath(".")		// search root dir for the config file
	err := viper.ReadInConfig()   	// Find and read the config file
	if err != nil {
		return Config{}, err
	}
	viper.SetDefault("PORT", "80")
	viper.SetDefault("LISTS", []string{"groceries"})
	viper.SetDefault("NEW_DB", false)

	var conf Config
	err = viper.Unmarshal(&conf)
	return conf, err
}

func New(fileName string) (*Config, error) {
	config, err := initViper(fileName)
	if err != nil {
		return nil, err
	}
	return &config, err
}
