package shippermodule

import (
	"errors"
	"fmt"
	"time"
)

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

type InsertShipperRequest struct {
	Name        string    `json:"shipper_name"`
	ImageURL    string    `json:"shipper_image"`
	Description string    `json:"shipper_description"`
	MaxWeight   int       `json:"max_weight"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}

func (p InsertShipperRequest) Sanitize() error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if p.MaxWeight == 0 {
		return errors.New("price cannot be empty")
	}
	if p.MaxWeight < 0 {
		return errors.New("invalid rating range")
	}
	return nil
}

type UpdateShipperRequest struct {
	Name        string    `json:"shipper_name"`
	ImageURL    string    `json:"shipper_image"`
	Description string    `json:"shipper_description"`
	MaxWeight   int       `json:"max_weight"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}

func (p UpdateShipperRequest) BuildQuery(id int64) (string, []interface{}) {
	var fieldQuery string
	fieldValues := make([]interface{}, 0)

	var i = 1
	if p.Name != "" {
		fieldQuery += fmt.Sprintf("name=$%d,", i)
		fieldValues = append(fieldValues, p.Name)
		i++
	}
	if p.ImageURL != "" {
		fieldQuery += fmt.Sprintf("image_url=$%d,", i)
		fieldValues = append(fieldValues, p.ImageURL)
		i++
	}
	if p.Description != "" {
		fieldQuery += fmt.Sprintf("description=$%d,", i)
		fieldValues = append(fieldValues, p.Description)
		i++
	}
	if p.MaxWeight != 0 {
		fieldQuery += fmt.Sprintf("price=$%d,", i)
		fieldValues = append(fieldValues, p.MaxWeight)
		i++
	}

	finalQuery := fmt.Sprintf(updateShipperQuery, fieldQuery[:len(fieldQuery)-1], id)

	return finalQuery, fieldValues
}
