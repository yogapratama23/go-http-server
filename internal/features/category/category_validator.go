package category

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryValidator struct{}

func (v *CategoryValidator) CreatePayload(r *http.Request) (*CreateCategoryInput, error) {
	var payload CreateCategoryInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	if payload.Name == "" {
		return nil, errors.New("name is required")
	}

	return &payload, nil
}

func (v *CategoryValidator) DeletePayload(r *http.Request) (*int, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 {
		return &id, errors.New("id is required")
	}

	return &id, nil
}

func (v *CategoryValidator) FindAllPayload(r *http.Request) (*response.PaginationInput, *FindAllWhereCond) {
	skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
	take, _ := strconv.Atoi(r.URL.Query().Get("take"))
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	search := r.URL.Query().Get("search")

	pagination := &response.PaginationInput{
		Skip: skip,
		Take: take,
	}

	whereCondition := &FindAllWhereCond{
		Id:     id,
		Search: search,
	}

	return pagination, whereCondition
}

func (v *CategoryValidator) UpdatePayload(r *http.Request) (*int, *UpdateCategoryInput, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 {
		return &id, nil, errors.New("id is required")
	}

	var payload UpdateCategoryInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, nil, err
	}

	if payload.Name == "" {
		return nil, nil, errors.New("name is required")
	}

	return &id, &payload, nil
}
