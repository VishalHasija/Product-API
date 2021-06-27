package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VishalHasija/Product-API.git/data"
	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
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

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Invalid URL ID unable to convert to number", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}

type KeyProduct struct{}

func (p *Product) MiddlewareJSONValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.NewProduct()
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error Decoding Data", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
