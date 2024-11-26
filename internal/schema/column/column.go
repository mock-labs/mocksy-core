package column

type Column[T any] struct {
	Name                  string
	DataType              T
	ReferentialConstraint *ReferentialConstraint
	Constraints           *Constraint
	Default               T
}
