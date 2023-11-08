package yaml

import (
	files "github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/ut080/bcs-portal/internal/logging"
)

func LoadFromFile(file files.File, out interface{}, logger logging.Logger) (err error) {
	raw, err := file.ReadFile()
	if err != nil {
		logger.Error().Err(err).Str("filename", file.FullPath()).Msg("failed to read yaml file")
		return errors.WithStack(err)
	}

	err = yaml.Unmarshal(raw, out)
	if err != nil {
		logger.Error().Err(err).Str("filename", file.FullPath()).Msg("failed to unmarshal yaml document")
		return errors.WithStack(err)
	}

	return nil
}
