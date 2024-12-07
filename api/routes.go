package api

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", GetProductByID).Methods("GET")
	r.HandleFunc("/products", GetProductsByUser).Methods("GET")

	return r
}
