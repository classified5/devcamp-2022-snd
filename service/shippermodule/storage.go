package shippermodule

import (
	"context"
	"database/sql"
	"log"

	m "github.com/classified5/devcamp-2022-snd/service/model"
)

type storage struct {
	ShipperDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ShipperDB: db,
	}
}

func (s *storage) AddShipper(ctx context.Context, data m.ShipperRequest) (result m.ShipperResponse, err error) {
	var id int64
	if err := s.ShipperDB.QueryRowContext(ctx, addShipperQuery,
		data.Name,
		data.ImageURL,
		data.Description,
		data.MaxWeight,
		data.CreatedAt,
		data.CreatedBy,
		data.UpdatedAt,
		data.UpdatedBy,
	).Scan(&id); err != nil {
		log.Println("[ShipperModule][AddShipper][Storage] problem querying to db, err: ", err.Error())
		return result, err
	}

	result.ID = id

	return
}

func (s *storage) GetShipper(ctx context.Context, id int64) (result m.ShipperResponse, err error) {
	if err := s.ShipperDB.QueryRowContext(ctx, getShipperQuery, id).Scan(
		&result.Name,
		&result.ImageURL,
		&result.Description,
		&result.MaxWeight,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	); err != nil {
		log.Println("[ShipperModule][GetShipper] problem querying to db, err: ", err.Error())
		return result, err
	}
	result.ID = id

	return
}

func (s *storage) DeleteShipper(ctx context.Context, id int64) (result m.ShipperResponse, err error) {
	res, err := s.ShipperDB.ExecContext(ctx, deleteShipperQuery, id)
	if err != nil {
		log.Println("[ShipperModule][DeleteShipper][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ShipperModule][DeleteShipper] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[ShipperModule][DeleteShipper] no rows affected in db")
		return
	}

	result.ID = id

	return
}

func (s *storage) GetShipperAll(ctx context.Context) (result []m.ShipperResponse, err error) {
	result = make([]m.ShipperResponse, 0)

	rows, err := s.ShipperDB.QueryContext(ctx, getShipperAllQuery)
	if err != nil {
		log.Println("[ShipperModule][GetShipperAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowData m.ShipperResponse
		if err = rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.ImageURL,
			&rowData.Description,
			&rowData.MaxWeight,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.UpdatedAt,
			&rowData.UpdatedBy,
		); err != nil {
			log.Println("[ShipperModule][GetShipperAll] problem with scanning db row, err: ", err.Error())
			return
		}
		result = append(result, rowData)
	}

	return
}

func (s *storage) UpdateShipper(ctx context.Context, id int64, param m.ShipperRequest) (result m.ShipperResponse, err error) {
	res, err := s.ShipperDB.ExecContext(ctx, updateShipperQuery,
		param.Name,
		param.ImageURL,
		param.Description,
		param.MaxWeight,
		param.CreatedAt,
		param.CreatedBy,
		param.UpdatedAt,
		param.UpdatedBy,
		id,
	)
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[ShipperModule][UpdateShipper] no rows affected in db")
		return
	}

	result.ID = id

	return
}
