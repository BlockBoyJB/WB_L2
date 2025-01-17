package db

import (
	"event-server/internal/model/dbmodel"
	"event-server/internal/repo/dberrs"
	"sync"
	"time"
)

type EventRepo struct {
	mx      sync.RWMutex
	storage map[int64]map[string]dbmodel.Event
}

func NewEventRepo() *EventRepo {
	return &EventRepo{
		storage: make(map[int64]map[string]dbmodel.Event),
	}
}

func (r *EventRepo) Create(id int64, e dbmodel.Event) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.storage[id]; !ok { // вообще надо отдельный метод на создание пользователя, но это уже душно
		r.storage[id] = make(map[string]dbmodel.Event)
	}

	if _, ok := r.storage[id][e.Id]; ok {
		return dberrs.ErrAlreadyExists
	}
	r.storage[id][e.Id] = e
	return nil
}

func (r *EventRepo) FindById(id int64, eventId string) (dbmodel.Event, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	e, ok := r.storage[id][eventId]
	if !ok {
		//тут возвращается неочевидная ошибка: не существует либо пользователя, либо события
		return dbmodel.Event{}, dberrs.ErrNotFound
	}
	return e, nil
}

func (r *EventRepo) Update(id int64, e dbmodel.Event) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.storage[id]; !ok { // вообще надо отдельный метод на создание пользователя, но это уже душно
		r.storage[id] = make(map[string]dbmodel.Event)
		return dberrs.ErrNotFound
	}

	r.storage[id][e.Id] = e
	return nil
}

func (r *EventRepo) Delete(id int64, eventId string) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.storage[id]; !ok { // вообще надо отдельный метод на создание пользователя, но это уже душно
		r.storage[id] = make(map[string]dbmodel.Event)
		return dberrs.ErrNotFound
	}

	delete(r.storage[id], eventId)
	return nil
}

func (r *EventRepo) FindInterval(id int64, start, end time.Time) ([]dbmodel.Event, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.storage[id]; !ok { // вообще надо отдельный метод на создание пользователя, но это уже душно
		r.storage[id] = make(map[string]dbmodel.Event)
		return nil, dberrs.ErrNotFound
	}

	var result []dbmodel.Event
	for _, e := range r.storage[id] {
		// полуинтервал
		if e.Date.Equal(start) || (e.Date.After(start) && e.Date.Before(end)) {
			result = append(result, e)
		}
	}
	return result, nil
}
