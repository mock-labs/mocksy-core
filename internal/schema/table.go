package schemas

import "github.com/mock-labs/mocksy-core/internal/schema/column"

type Table struct {
	Name    string
	Columns []column.Column[any]
}
