package org

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	InvalidMemberType MemberType = iota
	SeniorMember
	CadetMember
	CadetSponsorMember
)

func ParseMemberType(memberTypeStr string) (MemberType, error) {
	switch strings.ToUpper(memberTypeStr) {
	case "SENIOR":
		return SeniorMember, nil
	case "CADET":
		return CadetMember, nil
	case "CADET SPONSOR":
		return CadetSponsorMember, nil
	default:
		return InvalidMemberType, errors.Errorf("invalid member type: %s", memberTypeStr)
	}
}

func (mt MemberType) String() string {
	switch mt {
	case SeniorMember:
		return "SENIOR"
	case CadetMember:
		return "CADET"
	case CadetSponsorMember:
		return "CADET SPONSOR"
	default:
		panic(fmt.Errorf("invalid MemberType enum value: %d", mt))
	}
}
