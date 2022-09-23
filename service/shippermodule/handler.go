package shippermodule

import (
	"context"
	"log"
)

func (p *Module) AddShipper(ctx context.Context, data InsertShipperRequest) (ShipperResponse, error) {
	if err := data.Sanitize(); err != nil {
		log.Println("[ShipperModule][AddShipper] bad request, err: ", err.Error())
		return ShipperResponse{}, err
	}

	resp, err := p.Storage.AddShipper(ctx, data)
	if err != nil {
		log.Println("[ShipperModule][AddShipper] problem in getting from storage, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (p *Module) GetShipper(ctx context.Context, id int64) (ShipperResponse, error) {
	var resp ShipperResponse
	var err error

	resp, err = p.Storage.GetShipper(ctx, id)
	if err != nil {
		log.Println("[ShipperModule][GetShipper] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (p *Module) GetShipperAll(ctx context.Context) ([]ShipperResponse, error) {
	resp, err := p.Storage.GetShipperAll(ctx)
	if err != nil {
		log.Println("[ShipperModule][GetShipperAll] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (p *Module) UpdateShipper(ctx context.Context, id int64, data UpdateShipperRequest) (ShipperResponse, error) {
	resp, err := p.Storage.UpdateShipper(ctx, id, data)
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}
