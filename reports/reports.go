package reports

import (
	"path/filepath"

	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
)

func ConfigureLaTeXCompiler(logger logging.Logger) (*latex.Compiler, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find config directory")
	}

	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find cache directory")
	}

	assetDir := filepath.Join(cfgDir, "assets")
	buildDir := filepath.Join(cacheDir, "build")

	compiler := latex.NewCompiler(assetDir, buildDir, logger.DefaultLogger())

	return &compiler, nil
}
