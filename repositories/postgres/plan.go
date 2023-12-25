package postgres

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/pkg/planning"
)

type Plan struct {
	ID                    uuid.UUID
	PlanType              PlanType
	PlanNumber            string
	Title                 string
	PlanningStart         time.Time
	PlanningDue           time.Time
	EventStart            time.Time
	EventEnd              time.Time
	ProjectOfficerID      uuid.UUID
	CadetProjectOfficerID uuid.UUID
	Coordination          []Coordination
	ProjectOfficer        Member
	CadetProjectOfficer   Member
	PlanSections          []PlanSection
	TrainingBlocks        []TrainingBlock
}

func (p *Plan) planSectionsFromDomain(sections []planning.PlanSection) ([]PlanSection, error) {
	var sec []PlanSection
	for i, v := range sections {
		s := PlanSection{SectionNumber: i}
		err := s.FromDomainObject(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		sec = append(sec, s)
	}

	return sec, nil
}

func (p *Plan) coordinationFromDomain(coordination []planning.Coordination) ([]Coordination, error) {
	var coord []Coordination
	for i, v := range coordination {
		c := Coordination{PlanID: p.ID, CoordOrder: i + 1}
		err := c.FromDomainObject(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		coord = append(coord, c)
	}

	return coord, nil
}

func (p *Plan) fromCONPLAN(conplan planning.CONPLAN) error {
	coordination, err := p.coordinationFromDomain(conplan.GetCoordination())
	if err != nil {
		return errors.WithStack(err)
	}

	sections, err := p.planSectionsFromDomain(conplan.Sections)
	if err != nil {
		return errors.WithStack(err)
	}

	p.ID = conplan.ID()
	p.PlanType = CONPLAN
	p.Coordination = coordination
	p.PlanNumber = conplan.PlanNumber
	p.Title = conplan.Title
	p.PlanSections = sections

	return nil
}

func (p *Plan) fromOPLAN(oplan planning.OPLAN) error {
	coordination, err := p.coordinationFromDomain(oplan.GetCoordination())
	if err != nil {
		return errors.WithStack(err)
	}

	sections, err := p.planSectionsFromDomain(oplan.Sections)
	if err != nil {
		return errors.WithStack(err)
	}

	projectOfficer := Member{}
	err = projectOfficer.FromDomainObject(oplan.ProjectOfficer)
	if err != nil {
		return errors.WithStack(err)
	}

	cadetProjectOfficer := Member{}
	err = cadetProjectOfficer.FromDomainObject(oplan.CadetProjectOfficer)
	if err != nil {
		return errors.WithStack(err)
	}

	p.ID = oplan.ID()
	p.PlanType = OPLAN
	p.Coordination = coordination
	p.PlanNumber = oplan.PlanNumber
	p.Title = oplan.Title
	p.ProjectOfficerID = oplan.ProjectOfficer.ID()
	p.ProjectOfficer = projectOfficer
	p.CadetProjectOfficerID = oplan.CadetProjectOfficer.ID()
	p.CadetProjectOfficer = cadetProjectOfficer
	p.PlanSections = sections

	return nil
}

func (p *Plan) fromMeetingPlan(meetingPlan planning.MeetingPlan) error {
	coordination, err := p.coordinationFromDomain(meetingPlan.GetCoordination())
	if err != nil {
		return errors.WithStack(err)
	}

	var trainingBlocks []TrainingBlock
	for i, v := range meetingPlan.TrainingBlocks {
		trainingBlock := TrainingBlock{PlanID: p.ID, BlockNumber: i + 1}
		err = trainingBlock.FromDomainObject(v)
		if err != nil {
			return errors.WithStack(err)
		}

		trainingBlocks = append(trainingBlocks, trainingBlock)
	}

	p.ID = meetingPlan.ID()
	p.PlanType = MeetingPlan
	p.Coordination = coordination
	p.PlanningStart = meetingPlan.PlanningStart
	p.PlanningDue = meetingPlan.PlanDue
	p.TrainingBlocks = trainingBlocks

	return nil
}

func (p *Plan) planSectionsToDomain() []planning.PlanSection {
	var sections []planning.PlanSection
	for _, v := range p.PlanSections {
		section := planning.NewPlanSection(v.ID, v.Title, v.Body)
		sections = append(sections, section)
	}

	return sections
}

func (p *Plan) coordinationToDomain() []planning.Coordination {
	var coordination []planning.Coordination
	for _, v := range p.Coordination {
		office := v.Office.ToDomainObject().(org.DutyAssignment)
		coord := planning.NewCoordination(v.ID, office, v.Action, v.Completed, v.Outcome)
		coordination = append(coordination, coord)
	}

	return coordination
}

func (p *Plan) toCONPLAN() planning.CONPLAN {
	coordination := p.coordinationToDomain()
	sections := p.planSectionsToDomain()

	return planning.NewCONPLAN(p.ID, coordination, p.PlanNumber, p.Title, sections)
}

func (p *Plan) toOPLAN() planning.OPLAN {
	coordination := p.coordinationToDomain()
	sections := p.planSectionsToDomain()
	projectOfficer := p.ProjectOfficer.ToDomainObject().(org.Member)
	cadetProjectOfficer := p.CadetProjectOfficer.ToDomainObject().(org.Member)

	return planning.NewOPLAN(p.ID, coordination, p.PlanNumber, p.Title, projectOfficer, cadetProjectOfficer, sections)
}

func (p *Plan) toMeetingPlan() planning.MeetingPlan {
	panic(errors.New("Plan.toMeetingPlan() not implemented"))
}

func (p *Plan) FromDomainObject(object pkg.DomainObject) error {
	switch reflect.TypeOf(object).Name() {
	case "CONPLAN":
		return p.fromCONPLAN(*object.(*planning.CONPLAN))
	case "OPLAN":
		return p.fromOPLAN(*object.(*planning.OPLAN))
	case "MeetingPlan":
		return p.fromMeetingPlan(*object.(*planning.MeetingPlan))
	default:
		return errors.New("not a valid domain Plan object.")
	}
}

func (p *Plan) ToDomainObject() pkg.DomainObject {
	switch p.PlanType {
	case CONPLAN:
		plan := p.toCONPLAN()
		return &plan
	case OPLAN:
		plan := p.toOPLAN()
		return &plan
	case MeetingPlan:
		plan := p.toMeetingPlan()
		return &plan
	default:
		return nil
	}
}
