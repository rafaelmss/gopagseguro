package main

import (
	"log"
	"net/http"
	"../goPagSeguro/router"
)

func main() {

	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
