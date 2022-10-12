package sellermodule

import (
	"context"
	"log"

	m "github.com/classified5/devcamp-2022-snd/service/model"
)

func (p *Module) AddSeller(ctx context.Context, data m.SellerRequest) (result m.SellerResponse, err error) {
	data, err = SanitizeInsert(data)
	if err != nil {
		log.Println("[SellerModule][AddSeller] bad request, err: ", err.Error())
		return
	}

	result, err = p.Storage.AddSeller(ctx, data)
	if err != nil {
		log.Println("[SellerModule][AddSeller] problem in getting from storage, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetSeller(ctx context.Context, id int64) (result m.SellerResponse, err error) {
	result, err = p.Storage.GetSeller(ctx, id)
	if err != nil {
		log.Println("[SellerModule][GetSeller] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetSellerAll(ctx context.Context) (result []m.SellerResponse, err error) {
	result, err = p.Storage.GetSellerAll(ctx)
	if err != nil {
		log.Println("[SellerModule][GetSellerAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
