package ports

import "net/http"

type PasswordHandler interface {
	ServerHTTP(rw http.ResponseWriter, r* http.Request)
}