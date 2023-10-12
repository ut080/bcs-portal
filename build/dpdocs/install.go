package main

import (
	"os"
	"path/filepath"
	"strings"

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
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir).Msg("failed to create config directory")
	}

	logging.Info().Str("subidr", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg"), 0700)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "assets"), 0700)
	err = clearFileExistsError(err)
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
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cacheDir).Msg("failed to create cache directory")
	}

	logging.Info().Str("subidr", "capwatch").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cacheDir, "capwatch"), 0700)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "capwatch").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "build").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cacheDir, "build"), 0700)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "build").Msg("failed to create subdirectory")
	}

	return nil
}

func CopyAssets(projectPath string, logger logging.Logger) error {
	defaultCfgDir := filepath.Join(projectPath, "config")
	imgDir := filepath.Join(projectPath, "assets", "img")

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destImgDir := filepath.Join(cfgDir, "assets")
	destCfgDir := filepath.Join(cfgDir, "cfg")

	logging.Info().Str("file", "cap_command_emblem.jpg").Msg("copying asset")
	capEmblem, err := files.NewFile(filepath.Join(imgDir, "cap_command_emblem.jpg"), logger)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = capEmblem.Copy(destImgDir)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "cap_command_emblem.jpg").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "ut080_color.png").Msg("copying asset")
	bcsPatch, err := files.NewFile(filepath.Join(imgDir, "ut080_color.png"), logger)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = bcsPatch.Copy(destImgDir)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "ut080_color.png").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "duty_assignments.yaml").Msg("copying config")
	daCfg, err := files.NewFile(filepath.Join(defaultCfgDir, "duty_assignments.yaml"), logger)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = daCfg.Copy(destCfgDir)
	err = clearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}

	return nil
}

func clearFileExistsError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "file exists") {
		return nil
	}

	return err
}

func main() {
	logging.InitLogging("info", true)

	logger := logging.Logger{}

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

	err = CopyAssets(os.Args[1], logger)
	if err != nil {
		logging.Error().Err(err).Msg("failed to copy assets")
		os.Exit(1)
	}

	config.InitConfig()
}
