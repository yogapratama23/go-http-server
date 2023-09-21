package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yogapratama23/go-http-server/internal/constants/message"
	db "github.com/yogapratama23/go-http-server/internal/database"
	"github.com/yogapratama23/go-http-server/internal/features/auth"
	"github.com/yogapratama23/go-http-server/internal/features/category"
	"github.com/yogapratama23/go-http-server/internal/middlewares"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(message.FailedToLoadEnv)
	}

	// connect db
	db.Init()
	defer db.Connect.Close()

	port := os.Getenv("APP_PORT")
	addr := fmt.Sprintf(":%s", port)

	r := mux.NewRouter()

	// register middlewares
	r.Use(middlewares.Logging)
	r.Use(middlewares.Auth)

	// register routes
	r.HandleFunc("/", home).Methods("GET")
	category.CategoryRouters(r)
	auth.AuthRouters(r)

	http.ListenAndServe(addr, r)
}

func home(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(map[string]interface{}{
		"message": "Hello World!",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
