package ports

import "net/http"

// PasswordHandler define a interface para tipos que lidam com requisições HTTP relacionadas a senhas. 
type PasswordHandler interface {
	ServerHTTP(rw http.ResponseWriter, r* http.Request)
}