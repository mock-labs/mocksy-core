package schemas


type Schema struct {
	Tables map[string]*Table
}

func (s Schema) NewSchema() *Schema {
	return &Schema{
		Tables: make(map[string]*Table),
	}
}
