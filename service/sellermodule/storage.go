package sellermodule

import (
	"context"
	"database/sql"
	"log"

	m "github.com/classified5/devcamp-2022-snd/service/model"
)

type storage struct {
	SellerDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		SellerDB: db,
	}
}

func (s *storage) AddSeller(ctx context.Context, data m.SellerRequest) (result m.SellerResponse, err error) {
	var id int64
	if err := s.SellerDB.QueryRowContext(ctx, addSellerQuery,
		data.Name,
		data.Password,
		data.CreatedAt,
		data.CreatedBy,
		data.UpdatedAt,
		data.UpdatedBy,
	).Scan(&id); err != nil {
		log.Println("[SellerModule][AddSeller][Storage] problem querying to db, err: ", err.Error())
		return result, err
	}

	result.ID = id

	return
}

func (s *storage) GetSeller(ctx context.Context, id int64) (result m.SellerResponse, err error) {
	if err := s.SellerDB.QueryRowContext(ctx, getSellerAllQuery, id).Scan(
		&result.Name,
		&result.Password,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	); err != nil {
		log.Println("[SellerModule][GetSeller] problem querying to db, err: ", err.Error())
		return result, err
	}
	result.ID = id

	return
}

func (s *storage) GetSellerAll(ctx context.Context) (result []m.SellerResponse, err error) {
	result = make([]m.SellerResponse, 0)

	rows, err := s.SellerDB.QueryContext(ctx, getSellerAllQuery)
	if err != nil {
		log.Println("[SellerModule][GetSellerAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowData m.SellerResponse
		if err = rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.Password,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.UpdatedAt,
			&rowData.UpdatedBy,
		); err != nil {
			log.Println("[SellerModule][GetSellerAll] problem with scanning db row, err: ", err.Error())
			return
		}
		result = append(result, rowData)
	}

	return
}
