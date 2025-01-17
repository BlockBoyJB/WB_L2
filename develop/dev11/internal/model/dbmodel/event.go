package dbmodel

import "time"

type Event struct {
	Id    string    `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Date  time.Time `json:"date"`
}
