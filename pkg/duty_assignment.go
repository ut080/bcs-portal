package pkg

type DutyAssignment struct {
	OfficeSymbol string
	DutyTitle    string
	Assistant    bool
	Acting       bool
	MinGrade     *Grade
	MaxGrade     *Grade
	Assignee     *Member
}
