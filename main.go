package main

import (
	"log"
	"net/http"
	"test/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/twosum", controller.TwoSumHandler).Methods("POST")
	r.HandleFunc("/threesum", controller.ThreeSumHandler).Methods("POST")
	r.HandleFunc("/finsubstring", controller.FindSubstringHandler).Methods("POST")

	http.Handle("/", r)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}