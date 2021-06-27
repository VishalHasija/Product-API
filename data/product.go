package data

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

func NewProduct() *Product {
	return &Product{}
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	log.Println("JSON Decoding....")
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	return products
}

func AddProducts(prod *Product) {
	prod.ID = getID()
	products = append(products, prod)
}

func UpdateProduct(id int, prod *Product) error {
	idxInDB, err := findProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	products[idxInDB] = prod
	return nil

}

var ErrorProductNotListed = errors.New("Product not listed in Database")

func findProduct(id int) (int, error) {
	for idx, product := range products {
		if product.ID == id {
			return idx, nil
		}
	}
	return 0, ErrorProductNotListed
}

func getID() int {
	return products[len(products)-1].ID + 1
}

var products = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milk Coffee",
		Price:       2.45,
		SKU:         "lat245",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong Coffee without milk",
		Price:       3.45,
		SKU:         "esp345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
