package modelitem

import (
	"testing"

	"github.com/maszuari/apinovi/db"
	"github.com/stretchr/testify/assert"
)

func TestTwoForOnePromo(t *testing.T) {

	vou := db.Product{Code: "VOUCHER", Name: "NoviCap Voucher", Price: 5.00, TwoForOne: true, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}
	tsh := db.Product{Code: "TSHIRT", Name: "NoviCap T-Shirt", Price: 20.00, TwoForOne: false, BulkPurchase: true, BulkPurchasePrice: 19, BulkPurchaseMin: 3}
	mug := db.Product{Code: "MUG", Name: "NoviCap Mug", Price: 7.50, TwoForOne: false, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}

	data := []db.Product{}
	data = append(data, vou)
	data = append(data, tsh)
	data = append(data, mug)

	lists := db.Products{}
	lists.Products = data

	i1 := "VOUCHER"
	i2 := "TSHIRT"
	i3 := "VOUCHER"
	i4 := "VOUCHER"
	i5 := "VOUCHER"
	i6 := "VOUCHER"
	i7 := "MUG"

	items := make([]string, 7)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)
	items = append(items, i5)
	items = append(items, i6)
	items = append(items, i7)

	cart := Cart{}
	cart.Items = items

	bk := new(BookKeeper)
	bk.Tmp = make(map[string]float64)
	bk.TwoForOnePromo(vou.Code, vou.Price, cart)
	amount := bk.Tmp[vou.Code]

	expect := 15
	assert.Equal(t, float64(expect), amount)
	//assert.Equal(t, "NoviCap Voucher", lists.Products[0].Name)

}

func TestCalculate2For1Promo(t *testing.T) {

	price := 5.00
	num := 5
	r := Calculate2For1Promo(price, num)

	assert.Equal(t, float64(15), r)
}

func TestBulkPurchasePromoMoreOrEqualBulkMin(t *testing.T) {

	vou := db.Product{Code: "VOUCHER", Name: "NoviCap Voucher", Price: 5.00, TwoForOne: true, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}
	tsh := db.Product{Code: "TSHIRT", Name: "NoviCap T-Shirt", Price: 20.00, TwoForOne: false, BulkPurchase: true, BulkPurchasePrice: 19, BulkPurchaseMin: 3}
	mug := db.Product{Code: "MUG", Name: "NoviCap Mug", Price: 7.50, TwoForOne: false, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}

	data := []db.Product{}
	data = append(data, vou)
	data = append(data, tsh)
	data = append(data, mug)

	lists := db.Products{}
	lists.Products = data

	i1 := "TSHIRT"
	i2 := "TSHIRT"
	i3 := "VOUCHER"
	i4 := "TSHIRT"

	items := make([]string, 4)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)

	cart := Cart{}
	cart.Items = items

	bk := new(BookKeeper)
	bk.Tmp = make(map[string]float64)
	bk.BulkPurchasePromo(tsh.Code, tsh.Price, tsh.BulkPurchaseMin, tsh.BulkPurchasePrice, cart)
	amount := bk.Tmp[tsh.Code]

	var expect float64 = 57
	assert.Equal(t, expect, amount)

}

func TestBulkPurchasePromoLessThanBulkMin(t *testing.T) {

	tsh := db.Product{Code: "TSHIRT", Name: "NoviCap T-Shirt", Price: 20.00, TwoForOne: false, BulkPurchase: true, BulkPurchasePrice: 19, BulkPurchaseMin: 3}

	i1 := "TSHIRT"
	i2 := "TSHIRT"
	i3 := "VOUCHER"
	i4 := "MUG"

	items := make([]string, 4)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)

	cart := Cart{}
	cart.Items = items

	bk := new(BookKeeper)
	bk.Tmp = make(map[string]float64)
	bk.BulkPurchasePromo(tsh.Code, tsh.Price, tsh.BulkPurchaseMin, tsh.BulkPurchasePrice, cart)
	amount := bk.Tmp[tsh.Code]

	var expect float64 = 40
	assert.Equal(t, expect, amount)

}

func TestNormalPurchase(t *testing.T) {

	//vou := db.Product{Code: "VOUCHER", Name: "NoviCap Voucher", Price: 5.00, TwoForOne: true, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}
	//tsh := db.Product{Code: "TSHIRT", Name: "NoviCap T-Shirt", Price: 20.00, TwoForOne: false, BulkPurchase: true, BulkPurchasePrice: 19, BulkPurchaseMin: 3}
	mug := db.Product{Code: "MUG", Name: "NoviCap Mug", Price: 7.50, TwoForOne: false, BulkPurchase: false, BulkPurchasePrice: 0, BulkPurchaseMin: 0}

	i1 := "TSHIRT"
	i2 := "TSHIRT"
	i3 := "VOUCHER"
	i4 := "MUG"

	items := make([]string,4)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)

	cart := Cart{}
	cart.Items = items

	bk := new(BookKeeper)
	bk.Tmp = make(map[string]float64)
	bk.NormalPurchase(mug.Code, mug.Price, cart)
	amount := bk.Tmp[mug.Code]
	expect := 7.5
	assert.Equal(t, expect, amount)
}

func TestCheckoutEx1(t *testing.T) {
	data := createProducts()
	list := db.Products{}
	list.Products = data

	i1 := "TSHIRT"
	i2 := "MUG"
	i3 := "VOUCHER"

	items := make([]string, 3)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)

	cart := Cart{}
	cart.Items = items

	imodel := NewItemModel(list)
	var total float64
	total = imodel.Checkout(cart)

	var expect float64 = 32.50
	assert.Equal(t, expect, total)

}

func TestCheckoutEx2(t *testing.T) {
	data := createProducts()
	list := db.Products{}
	list.Products = data

	i1 := "VOUCHER"
	i2 := "TSHIRT"
	i3 := "VOUCHER"

	items := make([]string, 3)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)

	cart := Cart{}
	cart.Items = items

	imodel := NewItemModel(list)
	var total float64
	total = imodel.Checkout(cart)

	var expect float64 = 25
	assert.Equal(t, expect, total)
}

func TestCheckoutEx3(t *testing.T) {
	data := createProducts()
	list := db.Products{}
	list.Products = data

	i1 := "TSHIRT"
	i2 := "TSHIRT"
	i3 := "TSHIRT"
	i4 := "VOUCHER"
	i5 := "TSHIRT"

	items := make([]string,5)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)
	items = append(items, i5)

	cart := Cart{}
	cart.Items = items

	imodel := NewItemModel(list)
	var total float64
	total = imodel.Checkout(cart)

	var expect float64 = 81
	assert.Equal(t, expect, total)
}

func TestCheckoutEx4(t *testing.T) {
	data := createProducts()
	list := db.Products{}
	list.Products = data

	i1 := "VOUCHER"
	i2 := "TSHIRT"
	i3 := "VOUCHER"
	i4 := "VOUCHER"
	i5 := "MUG"
	i6 := "TSHIRT"
	i7 := "TSHIRT"

	items := make([]string,7)
	items = append(items, i1)
	items = append(items, i2)
	items = append(items, i3)
	items = append(items, i4)
	items = append(items, i5)
	items = append(items, i6)
	items = append(items, i7)

	cart := Cart{}
	cart.Items = items

	imodel := NewItemModel(list)
	var total float64
	total = imodel.Checkout(cart)

	var expect float64 = 74.5
	assert.Equal(t, expect, total)
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
