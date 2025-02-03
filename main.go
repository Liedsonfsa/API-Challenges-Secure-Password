package main

import (
	"net/http"
	"secure-password/adapters"
)

func main() {
	handler := &adapters.HTTPHandler{}
	http.Handle("/validate-password", handler)
	
	http.ListenAndServe(":3000", nil)
}
