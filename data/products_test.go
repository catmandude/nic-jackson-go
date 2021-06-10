package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Austin",
		Price: 1.00,
		SKU: "abs-asdfas",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
