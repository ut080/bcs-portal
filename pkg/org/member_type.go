package org

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type MemberType int

const (
	InvalidMemberType MemberType = iota
	SeniorMember
	CadetMember
	CadetSponsorMember
	PatronMember
)

func ParseMemberType(memberTypeStr string) (MemberType, error) {
	switch strings.ToUpper(memberTypeStr) {
	case "SENIOR":
		return SeniorMember, nil
	case "CADET":
		return CadetMember, nil
	case "CADET SPONSOR":
		return CadetSponsorMember, nil
	case "PATRON":
		return PatronMember, nil
	default:
		return InvalidMemberType, errors.Errorf("invalid member type: %s", memberTypeStr)
	}
}

func (mt *MemberType) MarshalYAML() (interface{}, error) {
	return mt.String(), nil
}

func (mt *MemberType) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.New("failed to scan MemberType")
	}

	t, err := ParseMemberType(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*mt = t
	return nil
}

func (mt *MemberType) String() string {
	switch *mt {
	case SeniorMember:
		return "SENIOR"
	case CadetMember:
		return "CADET"
	case CadetSponsorMember:
		return "CADET SPONSOR"
	case PatronMember:
		return "PATRON"
	default:
		panic(fmt.Errorf("invalid MemberType enum value: %d", *mt))
	}
}

func (mt *MemberType) UnmarshalYAML(value *yaml.Node) error {
	v, err := ParseMemberType(value.Value)
	if err != nil {
		return errors.WithStack(err)
	}

	*mt = v
	return nil
}

func (mt *MemberType) Value() (driver.Value, error) {
	return mt.String(), nil
}
