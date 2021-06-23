package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VishalHasija/Product-API.git/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Request")
	prod := data.NewProduct()
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error Decoding Data", http.StatusBadRequest)
		return
	}
	data.AddProducts(prod)
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprintf(rw, "Product successfully added")

}
