package v1

import (
	"log"
	"net/http"
)

// К сожалению, этим придется оборачивать каждую ручку (либо я не нашел как в стд добавить middleware)
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s | %s %s ", r.RemoteAddr, r.Method, r.RequestURI)
		next(w, r)
	}
}
