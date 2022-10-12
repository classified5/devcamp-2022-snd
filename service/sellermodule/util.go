package sellermodule

import (
	"errors"
	"time"

	m "github.com/classified5/devcamp-2022-snd/service/model"
)

func SanitizeInsert(param m.SellerRequest) (m.SellerRequest, error) {
	if param.Name == "" {
		return param, errors.New("name cannot be empty")
	}
	if param.Password == "" {
		return param, errors.New("password cannot be empty")
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
