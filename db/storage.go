package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Code              string
	Name              string
	Price             float64
	TwoForOne         bool    `json:"2for1"`
	BulkPurchase      bool    `json:"bulkpurchase"`
	BulkPurchasePrice float64 `json:"bulkpurchaseprice"`
	BulkPurchaseMin   int     `json:"bulkpurchasemin"`
}

func GetProducts() Products {

	jsonFile, err := os.Open("products.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products Products
	json.Unmarshal(byteValue, &products)
	return products
}
