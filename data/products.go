package data

import "time"

type Product struct {
	ID int
	Name string
	Description string
	Price float32
	SKU string
	CreatedOn string
	UpdatedOn string
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milk thing",
		Price: 2.45,
		SKU: "abc322",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee",
		Price: 3.45,
		SKU: "faa322",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	}
}