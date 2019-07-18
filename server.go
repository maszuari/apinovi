package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handleritem "github.com/maszuari/apinovi/handlers"
	modelitem "github.com/maszuari/apinovi/models"
	"github.com/maszuari/apinovi/db"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	storage := db.GetProducts()

	imodel := modelitem.NewItemModel(storage)
	h := handleritem.NewHandler(imodel)

	r := mux.NewRouter()
	r.HandleFunc("/checkout", h.Checkout).Methods("POST")
	r.HandleFunc("/hello/{name}/", h.Hello).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
