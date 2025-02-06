package core

import "strings"

// CheckPassword verifica se a senha obedece a todos os critérios de validação
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

// hasUpperCase verifica a existência de caracteres maiúsculos
func hasUpperCase(password string) bool {
    for _, char := range password {
        if char >= 'A' && char <= 'Z' {
            return true
        }
    }
    return false
}

// hasLowerCase verifica a existência de caracteres minúsculos
func hasLowerCase(password string) bool {
    for _, char := range password {
        if char >= 'a' && char <= 'z' {
            return true
        }
    }
    return false
}

// hasNumber verifica a existência de caracteres númericos
func hasNumber(password string) bool {
    for _, char := range password {
        if char >= '0' && char <= '9' {
            return true
        }
    }
    return false
}

// hasSpecialChar verifia a existência de caracteres especiais
func hasSpecialChar(password string) bool {
    especiais := "!@#$%^&*()_+{}[]:;<>,.?/~`"
    return strings.ContainsAny(password, especiais)
}