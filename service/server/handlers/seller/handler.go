package seller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/classified5/devcamp-2022-snd/service/model"
	"github.com/classified5/devcamp-2022-snd/service/server"
)

func (p *Handler) AddSellerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering AddSellerHandler")
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[SellerHandler][AddSeller] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data model.SellerRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[SellerHandler][AddSeller] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	res, err := p.seller.AddSeller(context.Background(), data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp := InsertSellerResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) GetSellerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetSeller Handler")
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[SellerHandler][GetSeller] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.seller.GetSeller(context.Background(), queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) GetSellerAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetSellerAll Handler")
	timeStart := time.Now()
	var err error

	resp, err := p.seller.GetSellerAll(context.Background())
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}
