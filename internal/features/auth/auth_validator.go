package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AuthValidator struct{}

func (v *AuthValidator) SignupPayload(r *http.Request) (*SignupInput, error) {
	var payload SignupInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	if payload.Username == "" {
		return nil, errors.New("username is required")
	}

	if payload.Password == "" {
		return nil, errors.New("password is required")
	}

	return &payload, nil
}

func (v *AuthValidator) SigninPayload(r *http.Request) (*SigninInput, error) {
	var payload SigninInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	if payload.Username == "" {
		return nil, errors.New("username is required")
	}

	if payload.Password == "" {
		return nil, errors.New("password is required")
	}

	return &payload, nil
}
