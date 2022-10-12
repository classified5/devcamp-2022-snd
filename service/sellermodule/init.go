package sellermodule

import (
	"database/sql"
)

type Module struct {
	Storage *storage
}

func NewSellerModule(db *sql.DB) *Module {
	return &Module{
		Storage: newStorage(db),
	}
}
