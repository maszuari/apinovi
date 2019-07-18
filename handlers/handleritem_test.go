package handleritem

import (

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	modelitem "github.com/maszuari/apinovi/models"
	"github.com/maszuari/apinovi/db"
)

type ItemModelStub struct{}

func (im *ItemModelStub) Checkout(cart modelitem.Cart) float64{
	return 99
} 

func TestCheckoutEx1(t *testing.T){

	str := `{"items":["VOUCHER","TSHIRT","MUG"]}`

	body := []byte(str)
	req, err := http.NewRequest("POST", "/checkout", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	data := createProducts()
	list := db.Products{}
	list.Products = data

	imodel := modelitem.NewItemModel(list)
	h := NewHandler(imodel)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/checkout", h.Checkout)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expect,err := createJSONOutput("n","Success", 32.5)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, expect, actual)
}

func TestCheckoutEx2(t *testing.T){
	str := `{"items":["VOUCHER","TSHIRT","VOUCHER"]}`

	body := []byte(str)
	req, err := http.NewRequest("POST", "/checkout", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	data := createProducts()
	list := db.Products{}
	list.Products = data

	imodel := modelitem.NewItemModel(list)
	h := NewHandler(imodel)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/checkout", h.Checkout)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expect,err := createJSONOutput("n","Success", 25.0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, expect, actual)
}

func TestCheckoutEx3(t *testing.T){
	str := `{"items":["TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"]}`

	body := []byte(str)
	req, err := http.NewRequest("POST", "/checkout", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	data := createProducts()
	list := db.Products{}
	list.Products = data

	imodel := modelitem.NewItemModel(list)
	h := NewHandler(imodel)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/checkout", h.Checkout)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expect,err := createJSONOutput("n","Success", 81.0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, expect, actual)
}

func TestCheckoutEx4(t *testing.T){
	str := `{"items":["VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"]}`

	body := []byte(str)
	req, err := http.NewRequest("POST", "/checkout", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	data := createProducts()
	list := db.Products{}
	list.Products = data

	imodel := modelitem.NewItemModel(list)
	h := NewHandler(imodel)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/checkout", h.Checkout)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expect,err := createJSONOutput("n","Success", 74.5)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, expect, actual)
}

func createJSONOutput(e string, m string, t float64) (string, error) {

	s := fmt.Sprintf("%.2f", t)
	res:= Output{Error: e, Message: m, Total: s}
	out, err := json.Marshal(res)
	if err != nil {
		return "none", err
	}
	return string(out), nil
}

func createProducts() []db.Product {
	p1 := db.Product{Code: "VOUCHER", Name: "NoviCap Voucher", Price: 5.00, TwoForOne: true, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}
	p2 := db.Product{Code: "TSHIRT", Name: "NoviCap T-Shirt", Price: 20.00, TwoForOne: false, BulkPurchase: true, BulkPurchasePrice: 19, BulkPurchaseMin: 3}
	p3 := db.Product{Code: "MUG", Name: "NoviCap Mug", Price: 7.50, TwoForOne: false, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}

	data := []db.Product{}
	data = append(data, p1)
	data = append(data, p2)
	data = append(data, p3)

	return data
}