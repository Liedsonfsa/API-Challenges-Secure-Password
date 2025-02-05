package core

import "strings"

func CheckPassword(password string) []string {
	var errors []string
    
	if len(password) < 8 {
		errors = append(errors, "A sua senha deve conter no mínimo 8 caracteres!")
		return errors
	}

	if !hasUpperCase(password) {
		errors = append(errors, "Sua senha precisa ter ao menos um caracter maiúsculo!")
	}

	if !hasLowerCase(password) {
		errors = append(errors, "Sua senha precisa ter ao menos um caracter minúsculo!")
	}

	if !hasNumber(password) {
		errors = append(errors, "Sua senha precisa ter ao menos um número!")
	}

	if !hasSpecialChar(password) {
		errors = append(errors, "Sua senha precisa ter ao menos um caracter especial!")
	}

	return errors
}

func hasUpperCase(password string) bool {
    for _, char := range password {
        if char >= 'A' && char <= 'Z' {
            return true
        }
    }
    return false
}

func hasLowerCase(password string) bool {
    for _, char := range password {
        if char >= 'a' && char <= 'z' {
            return true
        }
    }
    return false
}

func hasNumber(password string) bool {
    for _, char := range password {
        if char >= '0' && char <= '9' {
            return true
        }
    }
    return false
}

func hasSpecialChar(password string) bool {
    especiais := "!@#$%^&*()_+{}[]:;<>,.?/~`"
    return strings.ContainsAny(password, especiais)
}