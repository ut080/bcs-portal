package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Directory functions

// CfgDir returns the directory where configuration filing and assets are stored.
// If the environment variable BCSPORTAL_CONFIG is set, then that configuration directory will be used.
// Otherwise, it defaults to ${UserConfigDir}/bcs-portal
func CfgDir() (string, error) {
	cfgDir := os.Getenv("BCSPORTAL_CONFIG")
	if cfgDir != "" {
		return cfgDir, nil
	}

	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user home directory")
	}

	return filepath.Join(hd, "bcs-portal"), nil
}

// CacheDir returns the directory where CAPWATCH and other cache filing are stored.
// If the environment variable BCSPORTAL_CACHE is set, then that configuration directory will be used.
// Otherwise, it defaults to ${UserCacheDir}/bcs-portal
func CacheDir() (string, error) {
	cacheDir := os.Getenv("BCSPORTAL_CACHE")
	if cacheDir != "" {
		return cacheDir, nil
	}

	cd, err := os.UserCacheDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user cache directory")
	}

	return filepath.Join(cd, "bcs-portal"), nil
}

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

// GetInt wraps Viper's GetInt function.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString wraps Viper's GetString function.
func GetString(key string) string {
	return viper.GetString(key)
}
