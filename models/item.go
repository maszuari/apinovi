package modelitem

import (
	"math"
)

import (
	"github.com/maszuari/apinovi/db"
)

type ItemModelImp interface {
	Checkout(cart Cart) float64
}

type ItemModel struct {
	Products db.Products
}

type BookKeeper struct {
	Tmp map[string]float64
}

type Cart struct {
	Items []string `json:"items"`
}

func NewItemModel(list db.Products) *ItemModel {
	return &ItemModel{Products: list}
}

func (im *ItemModel) Checkout(cart Cart) float64 {

	bk := new(BookKeeper)
	bk.Tmp = make(map[string]float64)

	for _, p := range im.Products.Products {
		if p.TwoForOne {
			bk.TwoForOnePromo(p.Code, p.Price, cart)
		} else if p.BulkPurchase {
			bk.BulkPurchasePromo(p.Code, p.Price, p.BulkPurchaseMin, p.BulkPurchasePrice, cart)
		} else if !p.TwoForOne && !p.BulkPurchase {
			bk.NormalPurchase(p.Code, p.Price, cart)
		}
	}

	var total float64 = 0
	for _, p := range bk.Tmp {
		total = total + p
	}

	return total
}

func (bk *BookKeeper) TwoForOnePromo(productcode string, productprice float64, cart Cart) {

	num := 0
	for _, c := range cart.Items {
		if c == productcode {
			num++
		}
	}
	if num > 0 {
		amount := Calculate2For1Promo(productprice, num)
		bk.Tmp[productcode] = amount
	}

}

//Divide by 2
//Round up the result. Then multiple with price.
func Calculate2For1Promo(price float64, num int) float64 {

	x := float64(num) / 2
	y := math.Round(x)
	rs := price * y
	return rs
}

func (bk *BookKeeper) BulkPurchasePromo(productcode string, productprice float64, bulkmin int, bulkprice float64, cart Cart) {

	num := 0
	for _, c := range cart.Items {
		if c == productcode {
			num++
		}
	}
	if num >= bulkmin {
		amount := bulkprice * float64(num)
		bk.Tmp[productcode] = amount
	} else {
		amount := productprice * float64(num)
		bk.Tmp[productcode] = amount
	}
}

func (bk *BookKeeper) NormalPurchase(productcode string, productprice float64, cart Cart) {

	for _, c := range cart.Items {
		if c == productcode {
			_, ok := bk.Tmp[productcode]
			if ok {
				bk.Tmp[productcode] = bk.Tmp[productcode] + productprice
			} else {
				bk.Tmp[productcode] = productprice
			}
		}
	}
}