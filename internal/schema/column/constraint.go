package column

type Constraint struct {
	Null   bool
	Check  bool
	Unique bool
}
