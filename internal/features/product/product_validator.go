package product

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ProductValidator struct{}

func (v *ProductValidator) CreatePayload(r *http.Request) (*CreateProductInput, error) {
	var payload CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if payload.Name == "" {
		return nil, errors.New("name is required")
	}

	if payload.CategoryId == 0 {
		return nil, errors.New("category id is required")
	}

	return &payload, nil
}
