package productmodule

import (
	"context"
	"database/sql"
	"log"

	m "github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/model"
)

type storage struct {
	ProductDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ProductDB: db,
	}
}

func (s *storage) AddProduct(ctx context.Context, data m.ProductRequest) (result m.ProductResponse, err error) {
	var id int64
	if err := s.ProductDB.QueryRowContext(ctx, addProductQuery,
		data.Name,
		data.Price,
		data.Discount,
		data.Stock,
		data.CreatedAt,
		data.CreatedBy,
		data.UpdatedAt,
		data.UpdatedBy,
	).Scan(&id); err != nil {
		log.Println("[ProductModule][AddProduct][Storage] problem querying to db, err: ", err.Error())
		return result, err
	}

	result.ID = id

	return
}

func (s *storage) GetProduct(ctx context.Context, id int64) (result m.ProductResponse, err error) {
	if err := s.ProductDB.QueryRowContext(ctx, getProductQuery, id).Scan(
		&result.Name,
		&result.Price,
		&result.Discount,
		&result.Stock,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	); err != nil {
		log.Println("[ProductModule][GetProduct] problem querying to db, err: ", err.Error())
		return result, err
	}
	result.ID = id

	return
}

func (s *storage) GetProductAll(ctx context.Context) (result []m.ProductResponse, err error) {
	result = make([]m.ProductResponse, 0)

	rows, err := s.ProductDB.QueryContext(ctx, getProductAllQuery)
	if err != nil {
		log.Println("[ProductModule][GetProductAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowData m.ProductResponse
		if err = rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.Price,
			&rowData.Discount,
			&rowData.Stock,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.UpdatedAt,
			&rowData.UpdatedBy,
		); err != nil {
			log.Println("[ProductModule][GetProductAll] problem with scanning db row, err: ", err.Error())
			return
		}
		result = append(result, rowData)
	}

	return
}

func (s *storage) UpdateProduct(ctx context.Context, id int64, param m.ProductRequest) (result m.ProductResponse, err error) {
	res, err := s.ProductDB.ExecContext(ctx, updateProductQuery,
		param.Name,
		param.Price,
		param.Discount,
		param.Stock,
		param.CreatedAt,
		param.CreatedBy,
		param.UpdatedAt,
		param.UpdatedBy,
		id,
	)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[ProductModule][UpdateProduct] no rows affected in db")
		return
	}

	result.ID = id

	return
}
