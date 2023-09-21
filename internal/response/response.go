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
	page := 1
	perPage := total
	pageCount := 1

	if (p.Skip != 0) && (p.Take != 0) {
		page = int(math.Ceil(float64(p.Skip)/float64(p.Take))) + 1
		perPage = p.Take
		pageCount = int(math.Ceil(float64(total) / float64(p.Take)))
	}

	pagination := map[string]interface{}{
		"page":       page,
		"per_page":   perPage,
		"page_count": pageCount,
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
