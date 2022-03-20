package fields

type Requirements []Requirement

type Requirement struct {
	Operator string
	Filed    string
	Value    string
}
