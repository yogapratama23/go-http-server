package category

import (
	"encoding/json"
	"net/http"

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
	id := r.URL.Query().Get("id")
	search := r.URL.Query().Get("search")

	if id != "" {
		data, err := c.categoryService.FindById(id)
		if err != nil {
			response.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Success(w, "Find category by id!", http.StatusOK, data)
		return
	}

	if search != "" {
		data, err := c.categoryService.FindByName(&search)
		if err != nil {
			response.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Success(w, "Find category by search!", http.StatusOK, data)
		return
	}

	data, err := c.categoryService.FindAll()
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "Find all category!", http.StatusOK, data)
}

func (c *CategoryController) handleDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.categoryService.SoftDelete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(map[string]interface{}{
		"message": "Delete category success!",
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
