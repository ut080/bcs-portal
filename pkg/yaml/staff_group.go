package yaml

type StaffGroup struct {
	Group     string          `yaml:"group"`
	Subgroups []StaffSubGroup `yaml:"subgroups"`
}

type StaffSubGroup struct {
	Subgroup      string           `yaml:"subgroup"`
	Leader        DutyAssignment   `yaml:"leader"`
	DirectReports []DutyAssignment `yaml:"direct_reports"`
}
