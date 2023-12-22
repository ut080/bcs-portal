package planning

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

type CoordinationAction int

const (
	CoordAction CoordinationAction = iota
	ApproveAction
	SignAction
	ActionAction
)

func ParseCoordinationAction(action string) (CoordinationAction, error) {
	switch action {
	case "COORD":
		return CoordAction, nil
	case "APPROVE":
		return ApproveAction, nil
	case "SIGN":
		return SignAction, nil
	case "ACTION":
		return ActionAction, nil
	default:
		return -1, errors.Errorf("invalid coordination action: %s", action)
	}
}

func (ca *CoordinationAction) Scan(src any) error {
	coord, ok := src.(string)
	if !ok {
		return errors.New("failed to scan CoordinationAction")
	}

	c, err := ParseCoordinationAction(coord)
	if err != nil {
		return errors.WithStack(err)
	}

	*ca = c
	return nil
}

func (ca *CoordinationAction) String() string {
	switch *ca {
	case CoordAction:
		return "COORD"
	case ApproveAction:
		return "APPROVE"
	case SignAction:
		return "SIGN"
	case ActionAction:
		return "ACTION"
	default:
		panic(errors.Errorf("invalid value for CoordinationAction enum: %d", *ca))
	}
}

func (ca *CoordinationAction) Value() (driver.Value, error) {
	return ca.String(), nil
}
