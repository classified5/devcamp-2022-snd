package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/classified5/devcamp-2022-snd/service/database"
	"github.com/classified5/devcamp-2022-snd/service/server"
	shipperHandler "github.com/classified5/devcamp-2022-snd/service/server/handlers/shipper"
	"github.com/classified5/devcamp-2022-snd/service/shippermodule"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "12345",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}

	// Init DB connection
	log.Println("Initializing DB Connection")
	db := database.GetDatabaseConnection(dbConfig)

	// Init shipper usecase
	log.Println("Initializing Usecase")
	sm := shippermodule.NewShipperModule(db)

	// Init shipper handler
	log.Println("Initializing Handler")
	sh := shipperHandler.NewShipperHandler(sm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/shipper", sh.AddShipperHandler).Methods(http.MethodPost)
	router.HandleFunc("/shipper/{id}", sh.UpdateShipperHandler).Methods(http.MethodPut)
	router.HandleFunc("/shipper/{id}", sh.GetShipperHandler).Methods(http.MethodGet)
	router.HandleFunc("/shippers", sh.GetShipperAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/shippers/{id}", sh.DeleteShipperHandler).Methods(http.MethodDelete)
	router.HandleFunc("/", sh.RootHandler).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022-snd shipper service service is starting...")

	server.Serve(serverConfig, router)
}
