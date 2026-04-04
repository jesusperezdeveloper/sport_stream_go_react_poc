package httputil

import (
	"encoding/json"
	"net/http"
	"time"
)

type Meta struct {
	Total     int    `json:"total"`
	Timestamp string `json:"timestamp"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data interface{}, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := SuccessResponse{
		Data: data,
		Meta: Meta{
			Total:     total,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	json.NewEncoder(w).Encode(resp)
}

func JSONError(w http.ResponseWriter, status int, code, message, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
	json.NewEncoder(w).Encode(resp)
}

func DecodeJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
