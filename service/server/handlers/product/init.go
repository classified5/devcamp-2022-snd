package product

import "github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/productmodule"

type InsertProductResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	product *productmodule.Module
}

func NewProductHandler(p *productmodule.Module) *Handler {
	return &Handler{
		product: p,
	}
}
