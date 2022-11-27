package config

import (
	"github.com/spf13/viper"
)

// Config setup functions

// SetConfigType wraps Viper's SetConfigType function
func SetConfigType(t string) {
	viper.SetConfigType(t)
}

// SetConfigName wraps Viper's SetConfigName function
func SetConfigName(name string) {
	viper.SetConfigName(name)
}

// AddConfigPath wraps Viper's AddConfigPath function
func AddConfigPath(path string) {
	viper.AddConfigPath(path)
}

// SetDefault wraps Viper's SetDefault function.
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// ReadInConfig wraps Viper's ReadInConfig function.
func ReadInConfig() error {
	return viper.ReadInConfig()
}

// WriteConfigAs wraps Viper's WriteConfigAs function.
func WriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

// Configuration access functions

// GetString wraps Viper's GetString function.
func GetString(key string) string {
	return viper.GetString(key)
}
