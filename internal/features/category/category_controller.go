package category

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryController struct {
	categoryService CategoryService
	validate        CategoryValidator
}

func CategoryRouters(r *mux.Router) {
	controller := new(CategoryController)
	r.HandleFunc("/category", controller.handleCreate).Methods("POST")
	r.HandleFunc("/category", controller.handleFindAll).Methods("GET")
	r.HandleFunc("/category/{id}", controller.handleDelete).Methods("DELETE")
}

func (c *CategoryController) handleCreate(w http.ResponseWriter, r *http.Request) {
	payload, err := c.validate.CreatePayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.categoryService.Create(payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "Create category success!", http.StatusCreated, nil)
}

func (c *CategoryController) handleFindAll(w http.ResponseWriter, r *http.Request) {
	pagination, whereCondition := c.validate.FindAllPayload(r)

	data, err := c.categoryService.FindAll(pagination, whereCondition)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Pagination(w, "Find all categories", http.StatusOK, data.Categories, pagination, data.Total)
}

func (c *CategoryController) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := c.validate.DeletePayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.categoryService.SoftDelete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "Delete category success!", http.StatusOK, nil)
}
