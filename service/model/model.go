package model

import "time"

type ProductResponse struct {
	ID        int64     `json:"product_id,omitempty" db:"id"`
	Name      string    `json:"product_name,omitempty" db:"name"`
	Price     string    `json:"product_price,omitempty" db:"price"`
	Discount  string    `json:"product_discount,omitempty" db:"discount"`
	Stock     int       `json:"stock,omitempty" db:"stock"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	CreatedBy int       `json:"created_by,omitempty" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy int       `json:"updated_by,omitempty" db:"updated_by"`
}

type ProductRequest struct {
	Name      string    `json:"product_name"`
	Price     string    `json:"product_price"`
	Discount  string    `json:"product_discount"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}
