package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sraynitjsr/controller"
)

func main() {
	router := mux.NewRouter()
	controller.RegisterStudentRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
