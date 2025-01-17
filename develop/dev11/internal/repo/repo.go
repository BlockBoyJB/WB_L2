package repo

import (
	"event-server/internal/model/dbmodel"
	"event-server/internal/repo/db"
	"time"
)

type Event interface {
	Create(id int64, e dbmodel.Event) error
	FindById(id int64, eventId string) (dbmodel.Event, error)
	Update(id int64, e dbmodel.Event) error
	Delete(id int64, eventId string) error
	FindInterval(id int64, start, end time.Time) ([]dbmodel.Event, error)
}

type Repositories struct {
	Event
}

func NewRepositories() *Repositories {
	return &Repositories{
		Event: db.NewEventRepo(),
	}
}
