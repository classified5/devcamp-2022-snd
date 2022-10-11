package shipper

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

	"github.com/classified5/devcamp-2022-snd/service/model"
	"github.com/classified5/devcamp-2022-snd/service/server"
)

func (p *Handler) RootHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Entering RootHandler")
	fmt.Fprintf(w, "Hello Devcamp-2022-snd!")
}

func (p *Handler) AddShipperHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering AddShipperHandler")
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ShipperHandler][AddShipper] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data model.ShipperRequest
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

	resp := InsertShipperResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) GetShipperHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetShipper Handler")
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

func (p *Handler) DeleteShipperHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering DeleteShipper Handler")
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ShipperHandler][DeleteShipper] bad request, err: ", err.Error())
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

func (p *Handler) GetShipperAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetShipperAll Handler")
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

func (p *Handler) UpdateShipperHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering UpdateShipper Handler")
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

	var data model.ShipperRequest
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
