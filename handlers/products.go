package handlers

import (
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

func (P *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetProducts(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode data to json", http.StatusInternalServerError)
		return
	}
}
