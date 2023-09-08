package airtable

import (
	"strconv"
)

type MemberType string

const (
	Senior MemberType = "Senior"
	Cadet  MemberType = "Cadet"
)

type MemberFields struct {
	CAPID          uint       `json:"CAPID"`
	MembershipType MemberType `json:"Membership Type"`
	LastName       string     `json:"Last Name"`
	FirstName      string     `json:"First Name"`
	Grade          []string   `json:"Grade"`
}

func (mf MemberFields) Key() string {
	return strconv.Itoa(int(mf.CAPID))
}
