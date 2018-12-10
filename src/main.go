package main

import (
	"log"
	"net/http"

	ctrl "./controller"
	"github.com/gorilla/mux"

	sendGrid "./service"
)

func main() {
	sendGrid.SendEmail()
	router := mux.NewRouter()
	router.HandleFunc("/email/list", ctrl.GetEmails).Methods("GET")
	router.HandleFunc("/email", ctrl.SendEmail).Methods("POST")

	log.Println("Server started at port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}
