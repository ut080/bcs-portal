package postgres

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

type PlanType int

const (
	OPLAN PlanType = iota
	CONPLAN
	MeetingPlan
)

func ParsePlanType(planType string) (PlanType, error) {
	switch planType {
	case "OPLAN":
		return OPLAN, nil
	case "CONPLAN":
		return CONPLAN, nil
	case "Meeting Plan":
		return MeetingPlan, nil
	default:
		return -1, errors.Errorf("invalid plan type: %s", planType)
	}
}

func (pt *PlanType) Scan(src any) error {
	planType, ok := src.(string)
	if !ok {
		return errors.New("failed to scan PlanType")
	}

	p, err := ParsePlanType(planType)
	if err != nil {
		return errors.WithStack(err)
	}

	*pt = p
	return nil
}

func (pt *PlanType) String() string {
	switch *pt {
	case OPLAN:
		return "OPLAN"
	case CONPLAN:
		return "CONPLAN"
	case MeetingPlan:
		return "Meeting Plan"
	default:
		panic(errors.Errorf("invalid PlanType enum value: %d", *pt))
	}
}

func (pt *PlanType) Value() (driver.Value, error) {
	return pt.String(), nil
}
