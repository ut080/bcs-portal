package yaml

type Flight struct {
	Name           string         `yaml:"name"`
	Commander      DutyAssignment `yaml:"cc"`
	FlightSergeant DutyAssignment `yaml:"ccf"`
	Elements       []Element      `yaml:"elements"`
}

type Element struct {
	Name              string         `yaml:"name"`
	ElementLeader     DutyAssignment `yaml:"leader"`
	AsstElementLeader DutyAssignment `yaml:"assistant"`
	Members           []uint         `yaml:"members"`
}
