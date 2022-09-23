package shippermodule

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type storage struct {
	ShipperDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ShipperDB: db,
	}
}

func (s *storage) AddShipper(ctx context.Context, data InsertShipperRequest) (ShipperResponse, error) {
	var resp ShipperResponse

	var id int64
	if err := s.ShipperDB.QueryRowContext(ctx, addShipperQuery,
		data.Name,
		data.ImageURL,
		data.Description,
		data.MaxWeight,
	).Scan(&id); err != nil {
		log.Println("[ShipperModule][AddShipper][Storage] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ShipperResponse{
		ID: id,
	}
	return resp, nil
}

func (s *storage) GetShipper(ctx context.Context, id int64) (ShipperResponse, error) {
	var resp ShipperResponse

	if err := s.ShipperDB.QueryRowContext(ctx, getShipperQuery, id).Scan(
		&resp.Name,
		&resp.ImageURL,
		&resp.Description,
		&resp.MaxWeight,
	); err != nil {
		log.Println("[ShipperModule][GetShipper] problem querying to db, err: ", err.Error())
		return resp, err
	}
	resp.ID = id

	return resp, nil
}

func (s *storage) GetShipperAll(ctx context.Context) ([]ShipperResponse, error) {
	resp := make([]ShipperResponse, 0)

	rows, err := s.ShipperDB.QueryContext(ctx, getShipperAllQuery)
	if err != nil {
		log.Println("[ShipperModule][GetShipperAll] problem querying to db, err: ", err.Error())
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var rowData ShipperResponse
		if err := rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.ImageURL,
			&rowData.Description,
			&rowData.MaxWeight,
		); err != nil {
			log.Println("[ShipperModule][GetShipperAll] problem with scanning db row, err: ", err.Error())
			return resp, err
		}
		resp = append(resp, rowData)
	}

	return resp, nil
}

func (s *storage) UpdateShipper(ctx context.Context, id int64, data UpdateShipperRequest) (ShipperResponse, error) {
	var resp ShipperResponse

	query, values := data.BuildQuery(id)
	res, err := s.ShipperDB.ExecContext(ctx, query, values...)
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper] problem querying to db, err: ", err.Error())
		return resp, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper] problem querying to db, err: ", err.Error())
		return resp, err
	}
	if rowsAffected == 0 {
		log.Println("[ShipperModule][UpdateShipper] no rows affected in db")
		return resp, errors.New("no rows affected in db")
	}

	return ShipperResponse{
		ID: id,
	}, nil
}
