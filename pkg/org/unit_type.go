package org

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type UnitType int

const (
	GroupUnit UnitType = iota
	CompositeSquadronUnit
	CadetSquadronUnit
	SeniorSquadronUnit
	FlightUnit
	ActivityUnit
)

func ParseUnitType(unitTypeStr string) (UnitType, error) {
	switch strings.ToLower(unitTypeStr) {
	case "group":
		return GroupUnit, nil
	case "composite squadron":
		return CompositeSquadronUnit, nil
	case "cadet squadron":
		return CadetSquadronUnit, nil
	case "senior squadron":
		return SeniorSquadronUnit, nil
	case "flight":
		return FlightUnit, nil
	case "activity":
		return ActivityUnit, nil
	default:
		return -1, errors.Errorf("invalid unit type: %s", unitTypeStr)
	}
}

func (ut *UnitType) MarshalYAML() (interface{}, error) {
	return ut.String(), nil
}

func (ut *UnitType) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.New("failed to scan MemberType")
	}

	t, err := ParseUnitType(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*ut = t
	return nil
}

func (ut UnitType) String() string {
	switch ut {
	case GroupUnit:
		return "Group"
	case CompositeSquadronUnit:
		return "Composite Squadron"
	case CadetSquadronUnit:
		return "Cadet Squadron"
	case SeniorSquadronUnit:
		return "Senior Squadron"
	case FlightUnit:
		return "Flight"
	case ActivityUnit:
		return "Activity"
	default:
		panic(fmt.Errorf("invalid UnitType enum value: %d", ut))
	}
}

func (ut *UnitType) UnmarshalYAML(value *yaml.Node) error {
	v, err := ParseUnitType(value.Value)
	if err != nil {
		return errors.WithStack(err)
	}

	*ut = v
	return nil
}

func (ut UnitType) Value() (driver.Value, error) {
	return ut.String(), nil
}
