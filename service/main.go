package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Devcamp-2022-snd!")
}

func main() {
	log.Println("Devcamp-2022-snd backend service is starting...")

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":9090", nil)
}
