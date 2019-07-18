package handleritem

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	modelitem "github.com/maszuari/apinovi/models"
)

type Output struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Total   string    `json:"total"`
}

type handler struct {
	ItemModel *modelitem.ItemModel
}

func NewHandler(imodel *modelitem.ItemModel) *handler {
	return &handler{imodel}
}

func (h *handler) Checkout(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dec := json.NewDecoder(r.Body)
	var cart modelitem.Cart
	err := dec.Decode(&cart)
	r.Body.Close()

	o := &Output{}

	if err != nil {
		log.Println(err)
		o.Error = "y"
		o.Message = "Failed to read client request"
		s := fmt.Sprintf("%.2f", 0.00)
		o.Total = s
	} else {
		
		total := h.ItemModel.Checkout(cart)
		o.Error = "n"
		o.Message = "Success"
		s := fmt.Sprintf("%.2f", total)
		o.Total = s
	}

	json.NewEncoder(w).Encode(o)
}

func (h *handler) Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello "+name)
}