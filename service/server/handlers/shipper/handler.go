package shipper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/classified5/devcamp-2022-snd/service/server"
	"github.com/classified5/devcamp-2022-snd/service/shippermodule"
)

type Handler struct {
	shipper *shippermodule.Module
}

func NewShipperHandler(p *shippermodule.Module) *Handler {
	return &Handler{
		shipper: p,
	}
}

func (p *Handler) AddShipper(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ShipperHandler][AddShipper] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data shippermodule.InsertShipperRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ShipperHandler][AddShipper] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	res, err := p.shipper.AddShipper(context.Background(), data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp := cuShipperResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) GetShipper(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ShipperHandler][GetShipper] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.shipper.GetShipper(context.Background(), queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) GetShipperAll(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	var err error

	resp, err := p.shipper.GetShipperAll(context.Background())
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) UpdateShipper(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ShipperHandler][UpdateShipper] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ShipperHandler][UpdateShipper] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data shippermodule.UpdateShipperRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ShipperHandler][UpdateShipper] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.shipper.UpdateShipper(context.Background(), queryID, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}
