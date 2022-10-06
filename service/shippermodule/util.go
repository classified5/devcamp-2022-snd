package shippermodule

import (
	"errors"
	"fmt"
	"time"

	m "github.com/classified5/devcamp-2022-snd/service/model"
)

func SanitizeInsert(param m.ShipperRequest) (m.ShipperRequest, error) {
	if param.Name == "" {
		return param, errors.New("name cannot be empty")
	}
	if param.ImageURL == "" {
		return param, errors.New("image url cannot be empty")
	}
	if param.Description == "" {
		return param, errors.New("description cannot be empty")
	}
	if param.MaxWeight < 0 {
		return param, errors.New("invalid rating range")
	}
	if param.CreatedAt.IsZero() {
		param.CreatedAt = time.Now()
	}
	if param.CreatedBy == 0 {
		param.CreatedBy = 99999
	}
	if param.UpdatedAt.IsZero() {
		param.UpdatedAt = time.Now()
	}
	if param.UpdatedBy == 0 {
		param.UpdatedBy = 99999
	}

	return param, nil
}

func BuildQuery(id int64, param m.ShipperRequest) (finalQuery string, fieldValues []interface{}) {
	var fieldQuery string
	fieldValues = make([]interface{}, 0)

	var i = 1
	if param.Name != "" {
		fieldQuery += fmt.Sprintf("name=$%d,", i)
		fieldValues = append(fieldValues, param.Name)
		i++
	}

	fieldQuery += fmt.Sprintf("updated_by=$%d,", i)
	fieldValues = append(fieldValues, param.UpdatedAt)
	i++

	finalQuery = fmt.Sprintf(updateShipperQuery, fieldQuery[:len(fieldQuery)-1], id)

	return
}
