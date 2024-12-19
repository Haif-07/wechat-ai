package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // or viper.SetConfigType("toml")
	viper.AddConfigPath(filepath.Dir(ex)) // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}
