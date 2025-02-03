package core

import "strings"

func CheckPassword(password string) bool {
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