package calendar

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

type Uniform int

const (
	SemiFormalUniform Uniform = iota
	ServiceDressUniform
	ServiceUniform
	UtilityUniform
	FieldUniform
	PTUniform
	CivilianAttire
)

func ParseUniform(uniform string) (Uniform, error) {
	switch uniform {
	case "Semi-formal":
		return SemiFormalUniform, nil
	case "Service Dress":
		return ServiceDressUniform, nil
	case "Service":
		return ServiceUniform, nil
	case "Utility":
		return UtilityUniform, nil
	case "Field":
		return FieldUniform, nil
	case "PT":
		return PTUniform, nil
	case "Civilian":
		return CivilianAttire, nil
	default:
		return -1, errors.Errorf("invalid uniform type: %s", uniform)
	}
}

func (u *Uniform) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.New("failed to scan uniform")
	}

	v, err := ParseUniform(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*u = v
	return nil
}

func (u *Uniform) String() string {
	switch *u {
	case SemiFormalUniform:
		return "Semi-formal"
	case ServiceDressUniform:
		return "Service Dress"
	case ServiceUniform:
		return "Service"
	case UtilityUniform:
		return "Utility"
	case FieldUniform:
		return "Field"
	case PTUniform:
		return "PT"
	case CivilianAttire:
		return "Civilian"
	default:
		panic(errors.Errorf("invalid Uniform enum value: %d", u))
	}
}

func (u *Uniform) Value() (driver.Value, error) {
	return u.String(), nil
}
