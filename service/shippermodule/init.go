package shippermodule

import (
	"database/sql"
)

type Module struct {
	Storage *storage
}

func NewShipperModule(db *sql.DB) *Module {
	return &Module{
		Storage: newStorage(db),
	}
}
