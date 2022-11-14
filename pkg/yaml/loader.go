package yaml

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func LoadYamlDocFromFile(filename string, out interface{}) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to read yaml file")
		return errors.WithStack(err)
	}

	err = yaml.Unmarshal(raw, out)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal yaml document")
		return errors.WithStack(err)
	}

	return nil
}
