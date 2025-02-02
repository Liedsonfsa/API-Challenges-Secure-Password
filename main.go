package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/password", isSecurePassword)

	http.ListenAndServe(":3000", nil)
}

func isSecurePassword(rw http.ResponseWriter, r* http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	isSecure := checkPassword(request.Password)
	response := map[string]bool{"secure": isSecure}
	rw.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(rw).Encode(response)
}

func checkPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var numerico, especial, maisculo, minusculo bool = false, false, false, false

	for _, char := range password {
		if (char < 48 || char > 57) && (char < 65 || char > 90) && (char < 97 || char > 122) {
			especial = true
		}

		if char >= 65 && char <= 90 {
			maisculo = true
		}

		if char >= 97 && char <= 122 {
			minusculo = true
		}

		if char >= 48 && char <= 57 {
			numerico = true
		}

		if especial && maisculo && minusculo && numerico {
			return true
		}
	}

	return false
}