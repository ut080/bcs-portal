package config

import (
	"os"
	"path/filepath"

	"github.com/ut080/bcs-portal/app/logging"
)

func InitConfig() {
	SetDefault("capwatch.orgid", "")
	SetDefault("capwatch.username", "")
	SetDefault("capwatch.refresh", 7)

	cwd, _ := os.Getwd()
	logging.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		logging.Error().Err(err).Msg("could not find user config dir")
	} else {
		logging.Debug().Str("usrCfgDir", cfgDir).Msg("user config dir found")
	}

	SetConfigType("yaml")
	SetConfigName("config")
	AddConfigPath(filepath.Join(cfgDir, "bcsdp"))
	err = ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "bcsdp", "config.yaml")
		logging.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = os.Mkdir(filepath.Join(cfgDir, "bcsdp"), 0700)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config directory")
		}
		err = WriteConfigAs(path)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
