package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type ProductController struct {
	validator ProductValidator
	service   ProductService
}

func ProductRouters(r *mux.Router) {
	controller := new(ProductController)
	r.HandleFunc("/product", controller.handleCreate).Methods(http.MethodPost)
	r.HandleFunc("/product", controller.handleFindAll).Methods(http.MethodGet)
}

func (c *ProductController) handleCreate(w http.ResponseWriter, r *http.Request) {
	payload, err := c.validator.CreatePayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.Create(payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "create product success", http.StatusCreated, nil)
}

func (c *ProductController) handleFindAll(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.FindAll()
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "find all products with details", http.StatusOK, products)
}
