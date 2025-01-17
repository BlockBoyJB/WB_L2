package v1

import (
	"encoding/json"
	"event-server/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func parseEvent(r *http.Request) (service.EventInput, error) {
	title := r.FormValue("title")
	text := r.FormValue("text")
	if title == "" || text == "" {
		return service.EventInput{}, ErrInvalidEventInput
	}

	date, err := parseDate(r)
	if err != nil {
		return service.EventInput{}, err
	}

	return service.EventInput{
		Title: title,
		Text:  text,
		Date:  date,
	}, nil
}

func parseUserId(r *http.Request) (int64, error) {
	userId := r.FormValue("user_id")
	if userId == "" {
		return 0, ErrInvalidUserInput
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, ErrInvalidUserInput
	}
	return id, nil
}

func parseDate(r *http.Request) (time.Time, error) {
	date, err := time.Parse(time.DateOnly, r.FormValue("date"))
	if err != nil {
		return time.Time{}, ErrInvalidEventInput
	}
	return date, nil
}

func jsonResponse(w http.ResponseWriter, code int, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"result": %s}`, string(b))))
}
