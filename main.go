package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

	var hasNumber, hasSpecialChar, hasUpperCase, hasLowerCase bool = false, false, false, false
	especialChar := "!@#$%^&*()_+{}[]:;<>,.?/~`"

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowerCase = true
		} else if char >= '0' && char <= '9' {
			hasNumber = true
		} else if(strings.ContainsAny(password, especialChar)) {
			hasSpecialChar = true	
		}

		if hasSpecialChar && hasUpperCase && hasLowerCase && hasNumber {
			return true
		}
	}

	return false
}