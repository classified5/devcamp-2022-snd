package productmodule

import (
	"database/sql"
)

type Module struct {
	Storage *storage
}

func NewProductModule(db *sql.DB) *Module {
	return &Module{
		Storage: newStorage(db),
	}
}
