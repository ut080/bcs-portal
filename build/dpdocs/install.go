package main

import (
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/build"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
)

func CopyAssets(projectPath string, logger logging.Logger) error {
	defaultCfgDir := filepath.Join(projectPath, "config")
	defaultSchemaDir := filepath.Join(projectPath, "docs", "schemas")
	imgDir := filepath.Join(projectPath, "assets", "img")

	cfgDir, err := config.CfgDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destImgDir := filepath.Join(cfgDir, "assets")
	destCfgDir := filepath.Join(cfgDir, "cfg")

	logging.Info().Str("file", "cap_command_emblem.jpg").Msg("copying asset")
	capEmblem, err := files.NewFile(filepath.Join(imgDir, "cap_command_emblem.jpg"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = capEmblem.Copy(destImgDir)
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "cap_command_emblem.jpg").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "ut080_color.png").Msg("copying asset")
	bcsPatch, err := files.NewFile(filepath.Join(imgDir, "ut080_color.png"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = bcsPatch.Copy(destImgDir)
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "ut080_color.png").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "duty_assignments.yaml").Msg("copying config")
	daCfg, err := files.NewFile(filepath.Join(defaultCfgDir, "duty_assignments.yaml"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = daCfg.Copy(filepath.Join(destCfgDir, "defs"))
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}

	logging.Info().Str("file", "schemas/duty_assignments.json").Msg("copying schema")
	dutyAssignmentsSchema, err := files.NewFile(filepath.Join(defaultSchemaDir, "duty_assignments.json"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = dutyAssignmentsSchema.Copy(filepath.Join(destCfgDir, "schemas"))
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "schemas/duty_assignments.json").Msg("failed to copy schema")
	}

	logging.Info().Str("file", "schemas/duty_assignments.json").Msg("copying schema")
	toSchema, err := files.NewFile(filepath.Join(defaultSchemaDir, "table_of_organization.json"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = toSchema.Copy(filepath.Join(destCfgDir, "schemas"))
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "schemas/table_of_organization.json").Msg("failed to copy schema")
	}

	return nil
}

func main() {
	logging.InitLogging("info", true)

	logger := logging.Logger{}

	err := build.CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = build.CreateCacheDirectories()
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
