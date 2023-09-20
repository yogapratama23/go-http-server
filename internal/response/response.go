package response

import (
	"encoding/json"
	"math"
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

func Pagination(w http.ResponseWriter, msg string, code int, data interface{}, p *PaginationInput, total int) {
	if (p.Skip == 0) && (p.Take == 0) {
		p.Skip = 0
		p.Take = 10
	}

	pagination := map[string]interface{}{
		"page":       int(math.Ceil(float64(p.Skip)/float64(p.Take))) + 1,
		"per_page":   p.Take,
		"page_count": int(math.Ceil(float64(total) / float64(p.Take))),
		"total":      total,
	}

	resp, _ := json.Marshal(map[string]interface{}{
		"message":    msg,
		"data":       data,
		"pagination": pagination,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
