package yaml

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/ut080/bcs-portal/internal/logging"
)

func LoadYamlDocFromFile(filename string, out interface{}, logger logging.Logger) (err error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		logger.Error().Err(err).Str("filename", filename).Msg("failed to read yaml file")
		return errors.WithStack(err)
	}

	err = yaml.Unmarshal(raw, out)
	if err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal yaml document")
		return errors.WithStack(err)
	}

	return nil
}
