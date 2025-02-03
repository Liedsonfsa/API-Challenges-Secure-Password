package adapters

import (
	"encoding/json"
	// "log"
	"net/http"
	"secure-password/core"
)

type HTTPHandler struct {}

func (h *HTTPHandler) ServeHTTP(rw http.ResponseWriter, r* http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	isSecure := core.CheckPassword(request.Password)

	if !isSecure {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}