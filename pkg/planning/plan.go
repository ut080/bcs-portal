package planning

type Plan interface {
	GetCoordination() []Coordination
	UpdateCoordination(index int, coord Coordination)
}
