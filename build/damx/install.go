package main

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
)

func CreateConfigDirectories() error {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	logging.Info().Str("dir", cfgDir).Msg("creating config directory")
	err = os.Mkdir(cfgDir, 0700)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir).Msg("failed to create config directory")
	}

	logging.Info().Str("subidr", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg"), 0700)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "assets"), 0700)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "assets").Msg("failed to create subdirectory")
	}

	return nil
}

func CreateCacheDirectories() error {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}
	logging.Info().Str("dir", cacheDir).Msg("creating cache directory")
	err = os.Mkdir(cacheDir, 0700)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cacheDir).Msg("failed to create cache directory")
	}

	logging.Info().Str("subidr", "build").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cacheDir, "build"), 0700)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "build").Msg("failed to create subdirectory")
	}

	return nil
}

func CopyAssets(projectPath string) error {
	defaultCfgDir := filepath.Join(projectPath, "config")

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destCfgDir := filepath.Join(cfgDir, "cfg")

	logging.Info().Str("file", "disposition_instructions.yaml").Msg("copying config")
	err = files.Copy(filepath.Join(defaultCfgDir, "disposition_instructions.yaml"), filepath.Join(destCfgDir, "disposition_instructions.yaml"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}

	return nil
}

func main() {
	logging.InitLogging("info", true)

	err := CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = CreateCacheDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create cache directories")
		os.Exit(1)
	}

	err = CopyAssets(os.Args[1])
	if err != nil {
		logging.Error().Err(err).Msg("failed to copy assets")
		os.Exit(1)
	}

	config.InitConfig()
}
