package tests

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetUpTestConfig(testDataDir string) {
	_ = os.MkdirAll(filepath.Join(testDataDir, "cache"), os.ModePerm)
	_ = os.MkdirAll(filepath.Join(testDataDir, "config"), os.ModePerm)

	viper.SetDefault("capwatch.orgid", "")
	viper.SetDefault("capwatch.username", "")
	viper.SetDefault("capwatch.password", "")
	viper.SetDefault("capwatch.refresh", 7)
	viper.SetDefault("test_member.capid", 0)
	viper.SetDefault("test_member.last_name", "")
	viper.SetDefault("test_member.first_name", "")
	viper.SetDefault("test_member.grade", "")
	viper.SetDefault("test_member.member_type", "")

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(testDataDir, "config"))
	err := viper.ReadInConfig()
	if err != nil {
		path := filepath.Join(testDataDir, "config", "config.yaml")
		log.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
