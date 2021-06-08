package handlers

import (
	"github.com/catmandude/nic-jackson-go/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Unneccessary because of Gorilla Mux
//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		p.getProducts(rw, r)
//		return
//	} else if r.Method == http.MethodPost {
//		p.addProduct(rw, r)
//		return
//	} else if r.Method == http.MethodPut {
//		p.l.Println("Put")
//		reg := regexp.MustCompile(`/([0-9]+)`)
//		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
//
//		if len(g) != 1 {
//			http.Error(rw, "Invalid URI", http.StatusBadRequest)
//			return
//		}
//
//		if len(g[0]) != 2 {
//			http.Error(rw, "invalid URI", http.StatusBadRequest)
//			return
//		}
//
//		idString := g[0][1]
//		id, err := strconv.Atoi(idString)
//		if err != nil {
//			http.Error(rw, "invalid URI", http.StatusBadRequest)
//			return
//		}
//
//		p.updateProduct(id, rw, r)
//		p.l.Println("got id", id)
//	}
//
//	//catch all
//	rw.WriteHeader(http.StatusMethodNotAllowed)
//}

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle get products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		http.Error(rw, "ID error", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT products")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
	}
}