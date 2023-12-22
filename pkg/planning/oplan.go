package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg/org"
)

type OPLAN struct {
	id                  uuid.UUID
	coordination        []Coordination
	PlanNumber          string
	Title               string
	ProjectOfficer      org.Member
	CadetProjectOfficer org.Member
	Sections            []PlanSection
}

func (o *OPLAN) ID() uuid.UUID {
	return (*o).id
}

func (o *OPLAN) GetCoordination() []Coordination {
	return (*o).coordination
}

func (o *OPLAN) UpdateCoordination(idx int, coord Coordination) {
	if idx >= len((*o).coordination) {
		(*o).coordination = append((*o).coordination, coord)
	} else {
		(*o).coordination[idx] = coord
	}
}

func (o *OPLAN) LaTeX() string {
	panic(errors.New("OPLAN.LaTeX() not implemented"))
}
