package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryController struct {
	categoryService CategoryService
}

func CategoryRouters(r *mux.Router) {
	controller := new(CategoryController)
	r.HandleFunc("/category", controller.handleCreate).Methods("POST")
	r.HandleFunc("/category", controller.handleFindAll).Methods("GET")
	r.HandleFunc("/category/{id}", controller.handleDelete).Methods("DELETE")
}

func (c *CategoryController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateCategoryInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.categoryService.Create(&payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "Create category success!", http.StatusCreated, nil)
}

func (c *CategoryController) handleFindAll(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
	take, _ := strconv.Atoi(r.URL.Query().Get("take"))
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	search := r.URL.Query().Get("search")

	pagination := response.PaginationInput{
		Skip: skip,
		Take: take,
	}

	whereCondition := &FindAllWhereCond{
		Id:     id,
		Search: search,
	}
	data, err := c.categoryService.FindAll(&pagination, whereCondition)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Pagination(w, "Find all categories", http.StatusOK, data.Categories, &pagination, data.Total)
}

func (c *CategoryController) handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.categoryService.SoftDelete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "Delete category success!", http.StatusOK, nil)
}
