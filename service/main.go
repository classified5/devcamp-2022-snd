package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/classified5/devcamp-2022-snd/service/database"
	"github.com/classified5/devcamp-2022-snd/service/server"
	shipperHandler "github.com/classified5/devcamp-2022-snd/service/server/handlers/shipper"
	"github.com/classified5/devcamp-2022-snd/service/shippermodule"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Devcamp-2022-snd!")
}

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "admin",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}
	db := database.GetDatabaseConnection(dbConfig)

	sm := shippermodule.NewShipperModule(db)
	sh := shipperHandler.NewShipperHandler(sm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/shipper", sh.AddShipper).Methods(http.MethodPost)
	router.HandleFunc("/shipper/{id:[0-9]+}", sh.UpdateShipper).Methods(http.MethodPut)
	router.HandleFunc("/shipper/{id:[0-9]+}", sh.GetShipper).Methods(http.MethodGet)
	router.HandleFunc("/shippers", sh.GetShipperAll).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/").Handler(fs)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)

	log.Println("Devcamp-2022-snd service service is starting...")

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":9090", nil)
}
