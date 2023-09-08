package airtable

type GradeFields struct {
	Abbreviation string `json:"Abbreviation"`
	Name         string `json:"Name"`
	Phase        string `json:"Phase"`
}

func (gf GradeFields) Key() string {
	return gf.Abbreviation
}
