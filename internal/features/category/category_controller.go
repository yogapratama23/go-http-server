package category

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	categoryService CategoryService
}

func CategoryRouters(r *mux.Router) {
	controller := new(CategoryController)
	r.HandleFunc("/category", controller.handleCreate).Methods("POST")
	r.HandleFunc("/category", controller.handleFindAll).Methods("GET")
}

func (c *CategoryController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateCategoryInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp, _ := json.Marshal(map[string]interface{}{
			"message": err.Error(),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	err = c.categoryService.Create(&payload)
	if err != nil {
		resp, _ := json.Marshal(map[string]interface{}{
			"message": err.Error(),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(map[string]interface{}{
		"message": "Create category success!",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (c *CategoryController) handleFindAll(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	search := r.URL.Query().Get("search")

	if id != "" {
		data, err := c.categoryService.FindById(id)
		if err != nil {
			resp, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
			})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(map[string]interface{}{
			"message": "Get category success!",
			"data":    data,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	}

	if search != "" {
		data, err := c.categoryService.FindByName(&search)
		if err != nil {
			resp, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
			})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(map[string]interface{}{
			"message": "Get category success!",
			"data":    data,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	}

	data, err := c.categoryService.FindAll()
	if err != nil {
		resp, _ := json.Marshal(map[string]interface{}{
			"message": err.Error(),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(map[string]interface{}{
		"message": "Get all category success!",
		"data":    data,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
