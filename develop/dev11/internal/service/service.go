package service

import (
	"event-server/internal/repo"
	"time"
)

type (
	EventInput struct {
		Title string
		Text  string
		Date  time.Time
	}
	EventOutput struct {
		Id    string    `json:"id"`
		Title string    `json:"title"`
		Text  string    `json:"text"`
		Date  time.Time `json:"date"`
	}
)

type Event interface {
	Create(id int64, input EventInput) (string, error)
	FindById(id int64, eventId string) (EventOutput, error)
	Update(id int64, eventId string, input EventInput) error
	Delete(id int64, eventId string) error
	FindForDay(id int64, day time.Time) ([]EventOutput, error)
	FindForWeek(id int64, week time.Time) ([]EventOutput, error)
	FindForMonth(id int64, month time.Time) ([]EventOutput, error)
}

type (
	Services struct {
		Event Event
	}
	ServicesDependencies struct {
		Repos *repo.Repositories
	}
)

func NewServices(d *ServicesDependencies) *Services {
	return &Services{
		Event: newEventService(d.Repos.Event),
	}
}
