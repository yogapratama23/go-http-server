package response

import (
	"encoding/json"
	"net/http"
)

type PaginationInput struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

func Error(w http.ResponseWriter, error string, code int) {

	resp, _ := json.Marshal(map[string]interface{}{
		"message": error,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func Success(w http.ResponseWriter, msg string, code int, data interface{}) {
	resp, _ := json.Marshal(map[string]interface{}{
		"message": msg,
		"data":    data,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
