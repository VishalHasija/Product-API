package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		p.UpdateProduct(rw, r)
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

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")
	path := r.URL.Path
	reg := regexp.MustCompile("/([0-9]+)")
	g := reg.FindAllStringSubmatch(path, -1)
	p.l.Println(g)
	if len(g) != 1 {
		p.l.Println("Invalid URI more than 1 ID")
		http.Error(rw, "Invalid URL ID", http.StatusBadRequest)
		return
	}
	p.l.Println(g[0], " Capture group")
	if len(g[0]) != 2 {
		p.l.Println("Invalid URL more than one capture group")
		http.Error(rw, "Invalid URL ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(g[0][1])
	if err != nil {
		http.Error(rw, "Invalid URL ID unable to convert to number", http.StatusBadRequest)
		return
	}

	prod := data.NewProduct()
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error Decoding Data", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}
