package latex

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
)

type LaTeX interface {
	LaTeX() string
}

func GenerateLaTeX(latex LaTeX, outputFile files.File, assets []string, logger logging.Logger) error {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	for _, asset := range assets {
		assetFile, err := files.NewFile(filepath.Join(cfgDir, "assets", asset), logger)
		if err != nil {
			logger.Warn().Err(err).Str("file", asset).Msg("failed to acquire reference to asset")
		}

		_, err = assetFile.Copy(filepath.Join(cacheDir, "build", asset))
		if err != nil {
			// TODO: React to whether this build asset has already been copied
			logger.Warn().Err(err).Str("file", asset).Msg("failed to copy build asset")
		}
	}

	texSourceFileName := fmt.Sprintf("%s.tex", outputFile.Base())
	texSource, err := files.NewFile(filepath.Join(cacheDir, "build", texSourceFileName), logger)
	if err != nil {
		return errors.WithStack(err)
	}
	err = texSource.WriteString(latex.LaTeX())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func CompileLaTeX(outputFile files.File, logger logging.Logger) error {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	texSourceFile, err := files.NewFile(filepath.Join(cacheDir, "build", fmt.Sprintf("%s.tex", outputFile.Base())), logger)

	// First run
	cmd := exec.Command("pdflatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Second run (pdflatex usually needs two runs to get formatting right)
	cmd = exec.Command("pdflatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	cachedBuildfile, err := files.NewFile(filepath.Join(cacheDir, "build", fmt.Sprintf("%s.pdf", outputFile.Base())), logger)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = cachedBuildfile.Move(outputFile.Dir())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
