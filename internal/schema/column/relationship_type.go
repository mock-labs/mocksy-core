package column

type RelationType string

const (
	OneToOne   RelationType = "One-to-One"
	OneToMany  RelationType = "One-to-Many"
	ManyToMany RelationType = "Many-to-Many"
)

type ReferentialConstraint struct {
	RelationType RelationType
	PrimaryKey   bool
	ForeignKey   *Column[any]
}
