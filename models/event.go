package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateEventTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func (e *Event) Create(db *sql.DB) error {
	panic("not implemented")
}

func (e *Event) Update(db *sql.DB) error {
	panic("not implemented")
}

func (e *Event) Delete(db *sql.DB) error {
	panic("not implemented")
}

func GetEventByID(db *sql.DB, id int) (*Event, error) {
	panic("not implemented")
}

func GetEventsByDateRange(db *sql.DB, start, end time.Time) ([]*Event, error) {
	panic("not implemented")
}