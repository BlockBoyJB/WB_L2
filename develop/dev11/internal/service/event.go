package service

import (
	"errors"
	"event-server/internal/model/dbmodel"
	"event-server/internal/repo"
	"event-server/internal/repo/dberrs"
	"github.com/google/uuid"
	"time"
)

type eventService struct {
	event repo.Event
}

func newEventService(event repo.Event) *eventService {
	return &eventService{
		event: event,
	}
}

func (s *eventService) Create(id int64, input EventInput) (string, error) {
	eventId := uuid.NewString()
	e := dbmodel.Event{
		Id:    eventId,
		Title: input.Title,
		Text:  input.Text,
		Date:  input.Date,
	}
	if err := s.event.Create(id, e); err != nil {
		if errors.Is(err, dberrs.ErrNotFound) {
			return "", ErrUserNotFound
		}
		if errors.Is(err, dberrs.ErrAlreadyExists) {
			return "", ErrEventAlreadyExist
		}
		// тут log
		return "", err
	}
	return eventId, nil
}

func (s *eventService) FindById(id int64, eventId string) (EventOutput, error) {
	e, err := s.event.FindById(id, eventId)
	if err != nil {
		if errors.Is(err, dberrs.ErrNotFound) {
			return EventOutput{}, ErrEventNotFound
		}
		return EventOutput{}, err
	}
	return EventOutput{
		Id:    e.Id,
		Title: e.Title,
		Text:  e.Text,
		Date:  e.Date,
	}, nil
}

func (s *eventService) Update(id int64, eventId string, input EventInput) error {
	err := s.event.Update(id, dbmodel.Event{
		Id:    eventId,
		Title: input.Title,
		Text:  input.Text,
		Date:  input.Date,
	})
	if err != nil {
		if errors.Is(err, dberrs.ErrNotFound) {
			return ErrEventNotFound
		}
		return err
	}
	return nil
}

func (s *eventService) Delete(id int64, eventId string) error {
	if err := s.event.Delete(id, eventId); err != nil {
		if errors.Is(err, dberrs.ErrNotFound) {
			return ErrEventNotFound
		}
		return err
	}
	return nil
}

func (s *eventService) FindForDay(id int64, day time.Time) ([]EventOutput, error) {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day()+1, 0, 0, 0, 0, day.Location())
	return s.findInterval(id, start, end)
}

func (s *eventService) FindForWeek(id int64, week time.Time) ([]EventOutput, error) {
	start := time.Date(week.Year(), week.Month(), 1, 0, 0, 0, 0, week.Location())
	end := start.AddDate(0, 0, 7)
	return s.findInterval(id, start, end)
}

func (s *eventService) FindForMonth(id int64, month time.Time) ([]EventOutput, error) {
	start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	end := start.AddDate(0, 1, 0)
	return s.findInterval(id, start, end)
}

func (s *eventService) findInterval(id int64, start, end time.Time) ([]EventOutput, error) {
	events, err := s.event.FindInterval(id, start, end)
	if err != nil {
		if errors.Is(err, dberrs.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	result := make([]EventOutput, 0, len(events))
	for _, e := range events {
		result = append(result, EventOutput{
			Id:    e.Id,
			Title: e.Title,
			Text:  e.Text,
			Date:  e.Date,
		})
	}
	return result, nil
}
