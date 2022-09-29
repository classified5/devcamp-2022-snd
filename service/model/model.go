package model

import "time"

type ShipperResponse struct {
	ID          int64     `json:"shipper_id,omitempty" db:"id"`
	Name        string    `json:"shipper_name,omitempty" db:"name"`
	ImageURL    string    `json:"shipper_image,omitempty" db:"image_url"`
	Description string    `json:"shipper_description,omitempty" db:"description"`
	MaxWeight   int       `json:"max_weight,omitempty" db:"max_weight"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	CreatedBy   int       `json:"created_by,omitempty" db:"created_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy   int       `json:"updated_by,omitempty" db:"updated_by"`
}

type ShipperRequest struct {
	Name        string    `json:"shipper_name"`
	ImageURL    string    `json:"shipper_image"`
	Description string    `json:"shipper_description"`
	MaxWeight   int       `json:"max_weight"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}
