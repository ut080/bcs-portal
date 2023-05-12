package domain

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType string

const (
	SeniorMember MemberType = "SENIOR"
	CadetMember  MemberType = "CADET"
)

func ParseMemberType(memberTypeStr string) (mt MemberType, err error) {
	switch strings.ToUpper(memberTypeStr) {
	case "SENIOR":
		mt = SeniorMember
	case "CADET":
		mt = CadetMember
	default:
		err = errors.Errorf("invalid member type: %s", memberTypeStr)
		return "", err
	}

	return mt, nil
}
