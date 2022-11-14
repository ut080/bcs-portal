package config

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("capwatch.orgid", "")
	viper.SetDefault("capwatch.username", "")
	viper.SetDefault("capwatch.refresh", 7)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cwd, _ := os.Getwd()
	log.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		log.Error().Err(err).Msg("could not find user config dir")
	} else {
		log.Debug().Str("usrCfgDir", cfgDir).Msg("user config dir found")
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(cfgDir, "bcsdp"))
	err = viper.ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "bcsdp", "config.yaml")
		log.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = os.Mkdir(filepath.Join(cfgDir, "bcsdp"), 0700)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config directory")
		}
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
