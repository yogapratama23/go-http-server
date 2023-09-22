package category

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/constants/message"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryController struct {
	service   CategoryService
	validator CategoryValidator
}

func CategoryRouters(r *mux.Router) {
	controller := new(CategoryController)
	r.HandleFunc("/category", controller.handleCreate).Methods(http.MethodPost)
	r.HandleFunc("/category", controller.handleFindAll).Methods(http.MethodGet)
	r.HandleFunc("/category-products", controller.handleFindAllWithProducts).Methods(http.MethodGet)
	r.HandleFunc("/category/{id}", controller.handleDelete).Methods(http.MethodDelete)
	r.HandleFunc("/category/{id}", controller.handleUpdate).Methods(http.MethodPut)
}

func (c *CategoryController) handleCreate(w http.ResponseWriter, r *http.Request) {
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

	response.Success(w, message.CreateCategorySuccess, http.StatusCreated, nil)
}

func (c *CategoryController) handleFindAll(w http.ResponseWriter, r *http.Request) {
	pagination, whereCondition := c.validator.FindAllPayload(r)

	data, err := c.service.FindAll(pagination, whereCondition)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Pagination(w, message.FindAllCategories, http.StatusOK, data.Categories, pagination, data.Total)
}

func (c *CategoryController) handleFindAllWithProducts(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.FindAllWithProducts()
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "find categories with products", http.StatusOK, data)
}

func (c *CategoryController) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := c.validator.DeletePayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.SoftDelete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, message.DeleteCategorySuccess, http.StatusOK, nil)
}

func (c *CategoryController) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, payload, err := c.validator.UpdatePayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.Update(id, payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, message.UpdateCategorySuccess, http.StatusOK, nil)
}
