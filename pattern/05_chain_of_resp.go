package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Применяется, когда программа должна обрабатывать разнообразные запросы несколькими способами,
но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся

Плюсы
	1) Низкая зависимость между клиентом и обработчиками
	2) Реализует первые 2 принципа SOLID (single responsibility + open-closed)

Минусы
	1) Запрос может остаться никем не обработанным

Что в примере?
Реализована цепочка обработчиков для разных методов запросов
*/

type request uint8

const (
	Get request = iota
	Post
	Delete
)

type Handler interface {
	Handle(r request)
}

type GetHandler struct {
	next Handler
}

func (h *GetHandler) Handle(r request) {
	if r == Get {
		fmt.Println("handle get request")
	} else if h.next != nil {
		h.next.Handle(r)
	}
}

type PostHandler struct {
	next Handler
}

func (h *PostHandler) Handle(r request) {
	if r == Post {
		fmt.Println("handle post request")
	} else if h.next != nil {
		h.next.Handle(r)
	}
}

type DeleteHandler struct {
	next Handler
}

func (h *DeleteHandler) Handle(r request) {
	if r == Delete {
		fmt.Println("handle delete request")
	} else if h.next != nil {
		h.next.Handle(r)
	}
}
