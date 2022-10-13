package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/model"
	"github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/server"
)

func (p *Handler) RootHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Entering RootHandler")
	fmt.Fprintf(w, "Hello Devcamp-2022-snd!")
}

func (p *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering AddProductHandler")
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][AddProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data model.ProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][AddProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	res, err := p.product.AddProduct(context.Background(), data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp := InsertProductResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetProduct Handler")
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][GetProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.product.GetProduct(context.Background(), queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) GetProductAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetProductAll Handler")
	timeStart := time.Now()
	var err error

	resp, err := p.product.GetProductAll(context.Background())
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering UpdateProduct Handler")
	timeStart := time.Now()
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data model.ProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.product.UpdateProduct(context.Background(), queryID, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}
