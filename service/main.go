package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/classified5/devcamp-2022-snd/service/database"
	"github.com/classified5/devcamp-2022-snd/service/sellermodule"
	"github.com/classified5/devcamp-2022-snd/service/server"
	sellerHandler "github.com/classified5/devcamp-2022-snd/service/server/handlers/seller"
	shipperHandler "github.com/classified5/devcamp-2022-snd/service/server/handlers/shipper"
	"github.com/classified5/devcamp-2022-snd/service/shippermodule"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "12345",
		DBName:   "devcamp",
		Port:     9001,
		Host:     "localhost",
		SSLMode:  "disable",
	}

	// Init DB connection
	log.Println("Initializing DB Connection")
	db := database.GetDatabaseConnection(dbConfig)

	router := mux.NewRouter()

	initSeller(db, router)
	initShipper(db, router)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022-snd shipper and seller service service is starting...")

	server.Serve(serverConfig, router)
}

func initShipper(db *sql.DB, router *mux.Router) {
	// Init shipper usecase
	log.Println("Initializing Usecase")
	sm := shippermodule.NewShipperModule(db)

	// Init shipper handler
	log.Println("Initializing Handler")
	sh := shipperHandler.NewShipperHandler(sm)

	// REST Handlers
	router.HandleFunc("/shipper", sh.AddShipperHandler).Methods(http.MethodPost)
	router.HandleFunc("/shipper/{id}", sh.UpdateShipperHandler).Methods(http.MethodPut)
	router.HandleFunc("/shipper/{id}", sh.GetShipperHandler).Methods(http.MethodGet)
	router.HandleFunc("/shipper/{id}", sh.DeleteShipperHandler).Methods(http.MethodDelete)
	router.HandleFunc("/shippers", sh.GetShipperAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", sh.RootHandler).Methods(http.MethodGet)
}

func initSeller(db *sql.DB, router *mux.Router) {
	// Init shipper usecase
	log.Println("Initializing Usecase")
	sm := sellermodule.NewSellerModule(db)

	// Init shipper handler
	log.Println("Initializing Handler")
	sh := sellerHandler.NewSellerHandler(sm)

	// REST Handlers
	router.HandleFunc("/seller", sh.AddSellerHandler).Methods(http.MethodPost)
	router.HandleFunc("/seller/{id}", sh.GetSellerHandler).Methods(http.MethodGet)
	router.HandleFunc("/sellers", sh.GetSellerAllHandler).Methods(http.MethodGet)
}
