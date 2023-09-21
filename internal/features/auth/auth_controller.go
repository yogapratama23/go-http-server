package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type AuthController struct {
	validator AuthValidator
	service   AuthService
}

func AuthRouters(r *mux.Router) {
	controller := new(AuthController)
	r.HandleFunc("/signup", controller.handleSignup).Methods(http.MethodPost)
	r.HandleFunc("/signin", controller.handleSignin).Methods(http.MethodPost)
}

func (c *AuthController) handleSignup(w http.ResponseWriter, r *http.Request) {
	payload, err := c.validator.SignupPayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.SignUp(payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "signup success", http.StatusCreated, nil)
}

func (c *AuthController) handleSignin(w http.ResponseWriter, r *http.Request) {
	payload, err := c.validator.SigninPayload(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := c.service.Signin(payload)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, "signin success", http.StatusCreated, token)
}
