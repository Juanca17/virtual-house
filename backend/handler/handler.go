package handler

import (
	"encoding/json"
	"net/http"

	"../connection"
)

func ResponseWriter(message interface{}, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var res connection.JsonResponse
	res.Response = message
	res.Status = status
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
