package productmodule

import (
	"errors"
	"fmt"
	"time"

	m "github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/model"
)

func SanitizeInsert(param m.ProductRequest) (m.ProductRequest, error) {
	if param.Name == "" {
		return param, errors.New("name cannot be empty")
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

func BuildQuery(id int64, param m.ProductRequest) (finalQuery string, fieldValues []interface{}) {
	var fieldQuery string
	fieldValues = make([]interface{}, 0)

	var i = 1
	if param.Name != "" {
		fieldQuery += fmt.Sprintf("name=$%d,", i)
		fieldValues = append(fieldValues, param.Name)
		i++
	}
	if param.UpdatedBy != 0 {
		fieldQuery += fmt.Sprintf("updated_by=$%d,", i)
		fieldValues = append(fieldValues, param.UpdatedBy)
		i++
	} else {
		fieldQuery += fmt.Sprintf("updated_by=$%d,", i)
		fieldValues = append(fieldValues, 99999)
		i++
	}

	fieldQuery += fmt.Sprintf("updated_at=$%d,", i)
	fieldValues = append(fieldValues, param.UpdatedAt)
	i++

	finalQuery = fmt.Sprintf(updateProductQuery, fieldQuery[:len(fieldQuery)-1], id)

	return
}
