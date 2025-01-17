package v1

import (
	"errors"
	"event-server/internal/service"
	"net/http"
	"time"
)

type eventRouter struct {
	event service.Event
}

func newEventRouter(event service.Event) {
	e := &eventRouter{
		event: event,
	}
	http.HandleFunc("/create_event", loggingMiddleware(e.create))
	http.HandleFunc("/find_event", loggingMiddleware(e.findById))
	http.HandleFunc("/update_event", loggingMiddleware(e.update))
	http.HandleFunc("/delete_event", loggingMiddleware(e.delete))

	http.HandleFunc("/events_for_day", loggingMiddleware(e.findForDay))
	http.HandleFunc("/events_for_week", loggingMiddleware(e.findForWeek))
	http.HandleFunc("/events_for_month", loggingMiddleware(e.findForMonth))
}

func (er *eventRouter) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	userId, err := parseUserId(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	e, err := parseEvent(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	eventId, err := er.event.Create(userId, e)
	if err != nil {
		errorResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	jsonResponse(w, http.StatusCreated, eventId)
}

func (er *eventRouter) findById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	userId, err := parseUserId(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	e, err := er.event.FindById(userId, r.FormValue("event_id"))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) && errors.Is(err, service.ErrEventNotFound) {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		errorResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	jsonResponse(w, http.StatusOK, e)
}

func (er *eventRouter) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userId, err := parseUserId(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	eventId := r.FormValue("event_id")
	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = er.event.Update(userId, eventId, event); err != nil {
		if errors.Is(err, service.ErrEventNotFound) {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		errorResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (er *eventRouter) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userId, err := parseUserId(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	eventId := r.FormValue("event_id")
	if err = er.event.Delete(userId, eventId); err != nil {
		if errors.Is(err, service.ErrEventNotFound) {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		errorResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (er *eventRouter) findForDay(w http.ResponseWriter, r *http.Request) {
	findInterval(w, r, er.event.FindForDay)
}

func (er *eventRouter) findForWeek(w http.ResponseWriter, r *http.Request) {
	findInterval(w, r, er.event.FindForWeek)
}

func (er *eventRouter) findForMonth(w http.ResponseWriter, r *http.Request) {
	findInterval(w, r, er.event.FindForMonth)
}

func findInterval(w http.ResponseWriter, r *http.Request, f func(int64, time.Time) ([]service.EventOutput, error)) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	userId, err := parseUserId(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	date, err := parseDate(r)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	e, err := f(userId, date)
	if err != nil {
		if errors.Is(err, service.ErrEventNotFound) {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		errorResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	jsonResponse(w, http.StatusOK, e)
}
