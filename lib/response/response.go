package response

import (
	"net/http"
)

func SendResponse(w http.ResponseWriter, httpStatus int, jsonResponse []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(jsonResponse)
}
