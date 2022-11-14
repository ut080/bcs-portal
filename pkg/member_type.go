package pkg

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType string

const (
	SeniorMember MemberType = "SENIOR"
	CadetMember  MemberType = "CADET"
)

func ParseMemberType(memberType string) (MemberType, error) {
	switch strings.ToUpper(memberType) {
	case "SENIOR":
		return SeniorMember, nil
	case "CADET":
		return CadetMember, nil
	}

	return "", errors.Errorf("invalid member type: %s", memberType)
}
