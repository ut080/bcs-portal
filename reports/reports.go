package reports

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
)

type LaTeXable interface {
	LaTeX() string
}

type DocLaTeXable interface {
	DocLaTeX() string
}

func generateLaTeX(report DocLaTeXable, buildDir, filename string, logger logging.Logger) error {

	latexFilePath := filepath.Join(buildDir, fmt.Sprintf("%s.tex", filename))
	err := files.Write(latexFilePath, report.DocLaTeX())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func copyAssets(assets []string, assetDir, buildDir string) error {
	for _, asset := range assets {
		err := files.Copy(filepath.Join(assetDir, asset), filepath.Join(buildDir, asset))
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func removeAssets(assets []string, buildDir string) error {
	for _, asset := range assets {
		err := files.Remove(filepath.Join(buildDir, asset))
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func CompileLaTeX(report DocLaTeXable, assetDir, outputPath, filename string, assets []string, logger logging.Logger) error {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	buildDir := filepath.Join(cacheDir, "build")

	err = generateLaTeX(report, buildDir, filename, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = copyAssets(assets, assetDir, buildDir)
	if err != nil {
		return errors.WithStack(err)
	}

	// pdflatex usually needs to be run a couple of times to get formatting right. The magic number seems to be 3.
	for i := 0; i < 3; i++ {
		cmd := exec.Command("pdflatex", "-halt-on-error", fmt.Sprintf("%s.tex", filename))
		cmd.Dir = filepath.Join(cacheDir, "build")

		err = cmd.Run()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	err = files.Move(
		filepath.Join(cacheDir, "build", fmt.Sprintf("%s.pdf", filename)),
		filepath.Join(outputPath, fmt.Sprintf("%s.pdf", filename)),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	/*
		_ = files.Remove(fmt.Sprintf("%s.*", filename))
		_ = removeAssets(assets, assetDir)
	*/

	return nil
}
