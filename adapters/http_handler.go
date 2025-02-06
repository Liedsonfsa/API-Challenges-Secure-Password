package adapters

import (
	"encoding/json"
	"net/http"
	"secure-password/core"
)

type HTTPHandler struct {}

// ServeHTTP recebe a requisição, decodifica e envia a senha para a validação. Por fim, define o status da requisição e retorna eventuais erros.
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
		
		response := map[string]string{
			"error": "O corpo da requisição contém um JSON inválido ou malformado.",
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
		return
	}

	validationErrors := core.CheckPassword(request.Password)

	if len(validationErrors) > 0 {
		rw.WriteHeader(http.StatusBadRequest)
		
		response := map[string][]string{
			"error": validationErrors,
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}