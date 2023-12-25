package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg/org"
)

type OPLAN struct {
	id                  uuid.UUID
	coordination        []Coordination
	planNumber          string
	title               string
	projectOfficer      org.Member
	cadetProjectOfficer org.Member
	sections            []PlanSection
}

func NewOPLAN(
	id uuid.UUID,
	coordination []Coordination,
	planNumber string,
	title string,
	projectOfficer org.Member,
	cadetProjectOfficer org.Member,
	sections []PlanSection,
) OPLAN {
	return OPLAN{
		id:                  id,
		coordination:        coordination,
		planNumber:          planNumber,
		title:               title,
		projectOfficer:      projectOfficer,
		cadetProjectOfficer: cadetProjectOfficer,
		sections:            sections,
	}
}

func (o OPLAN) ID() uuid.UUID {
	return o.id
}

func (o OPLAN) Coordination() []Coordination {
	return o.coordination
}

func (o OPLAN) PlanNumber() string {
	return o.planNumber
}

func (o OPLAN) Title() string {
	return o.title
}

func (o OPLAN) ProjectOfficer() org.Member {
	return o.projectOfficer
}

func (o OPLAN) CadetProjectOfficer() org.Member {
	return o.cadetProjectOfficer
}

func (o OPLAN) Sections() []PlanSection {
	return o.sections
}

func (o OPLAN) LaTeX() string {
	panic(errors.New("OPLAN.LaTeX() not implemented"))
}
