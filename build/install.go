package build

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
)

func CreateConfigDirectories() error {
	cfgDir, err := config.CfgDir()
	if err != nil {
		return errors.WithStack(err)
	}

	logging.Info().Str("dir", cfgDir).Msg("creating config directory")
	err = os.Mkdir(cfgDir, 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir).Msg("failed to create config directory")
	}

	logging.Info().Str("subidr", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "cfg/defs").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg", "defs"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg/defs").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "cfg/schemas").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg", "schemas"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg/schemas").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "assets"), 0700)
	err = ClearFileExistsError(err)
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
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cacheDir).Msg("failed to create cache directory")
	}

	logging.Info().Str("subidr", "build").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cacheDir, "build"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "build").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "capwatch").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cacheDir, "capwatch"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "capwatch").Msg("failed to create subdirectory")
	}

	return nil
}

func ClearFileExistsError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "file exists") {
		return nil
	}

	return err
}
