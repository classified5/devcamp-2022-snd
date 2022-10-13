package shipper

import "github.com/classified5/devcamp-2022-snd/service/shippermodule"

type InsertShipperResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	shipper *shippermodule.Module
}

func NewShipperHandler(p *shippermodule.Module) *Handler {
	return &Handler{
		shipper: p,
	}
}
