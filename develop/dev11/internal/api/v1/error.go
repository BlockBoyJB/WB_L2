package v1

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidUserInput  = errors.New("invalid user input: required user id")
	ErrInvalidEventInput = errors.New("invalid event input: cannot parse")
)

func errorResponse(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
}
