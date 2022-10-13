package productmodule

import (
	"context"
	"log"

	m "github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/model"
)

func (p *Module) AddProduct(ctx context.Context, data m.ProductRequest) (result m.ProductResponse, err error) {
	data, err = SanitizeInsert(data)
	if err != nil {
		log.Println("[ProductModule][AddProduct] bad request, err: ", err.Error())
		return
	}

	result, err = p.Storage.AddProduct(ctx, data)
	if err != nil {
		log.Println("[ProductModule][AddProduct] problem in getting from storage, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetProduct(ctx context.Context, id int64) (result m.ProductResponse, err error) {
	result, err = p.Storage.GetProduct(ctx, id)
	if err != nil {
		log.Println("[ProductModule][GetProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetProductAll(ctx context.Context) (result []m.ProductResponse, err error) {
	result, err = p.Storage.GetProductAll(ctx)
	if err != nil {
		log.Println("[ProductModule][GetProductAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) UpdateProduct(ctx context.Context, id int64, data m.ProductRequest) (result m.ProductResponse, err error) {
	result, err = p.Storage.UpdateProduct(ctx, id, data)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
