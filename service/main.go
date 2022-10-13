package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/database"
	"github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/server"
	productHandler "github.com/adhafajri/devcamp_pe_muhammad_adha_fajri_jonison/service/server/handlers/product"
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

	// Init product usecase
	log.Println("Initializing Usecase")
	sm := productHandler.NewProductModule(db)

	// Init product handler
	log.Println("Initializing Handler")
	sh := productHandler.NewProductHandler(sm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/product", sh.AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", sh.UpdateProductHandler).Methods(http.MethodPut)
	router.HandleFunc("/product/{id}", sh.GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", sh.GetProductAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", sh.RootHandler).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022-snd product service service is starting...")

	server.Serve(serverConfig, router)
}
