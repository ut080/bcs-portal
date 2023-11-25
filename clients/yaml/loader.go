package yaml

import (
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"

	"github.com/ut080/bcs-portal/internal/logging"
)

func validateYaml(data []byte, schema *jsonschema.Schema) error {
	var out interface{}

	err := yaml.Unmarshal(data, &out)
	if err != nil {
		return err
	}

	return schema.Validate(out)
}

func LoadSchemaFromConfig(schemaName string) (*jsonschema.Schema, error) {
	panic("LoadSchemaFromConfig() function not implemented")
}

func LoadFromFile(file files.File, out interface{}, schema *jsonschema.Schema, logger logging.Logger) error {
	raw, err := file.ReadFile()
	if err != nil {
		logger.Error().Err(err).Str("filename", file.FullPath()).Msg("failed to read yaml file")
		return errors.WithStack(err)
	}

	if schema != nil {
		err = validateYaml(raw, schema)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	err = yaml.Unmarshal(raw, out)
	if err != nil {
		logger.Error().Err(err).Str("filename", file.FullPath()).Msg("failed to unmarshal yaml document")
		return errors.WithStack(err)
	}

	return nil
}
