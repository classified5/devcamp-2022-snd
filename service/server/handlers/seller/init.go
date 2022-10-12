package seller

import "github.com/classified5/devcamp-2022-snd/service/sellermodule"

type InsertSellerResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	seller *sellermodule.Module
}

func NewSellerHandler(p *sellermodule.Module) *Handler {
	return &Handler{
		seller: p,
	}
}
