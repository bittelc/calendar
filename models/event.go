package models

import (
	"database/sql"
	"errors"
	"time"
)

type Event struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	Attendees   []Attendee  `json:"attendees"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Attendee struct {
	Name               string     `json:"name"`
	Email              string     `json:"email"`
	Role               string     `json:"role"`                 // "organizer", "required", "optional"
	Status             string     `json:"status"`               // "accepted", "declined", "tentative", "no-response"
	InvitedBy          string     `json:"invited_by"`           // email of inviter
	InvitedAt          time.Time  `json:"invited_at"`
	ResponseRequiredBy *time.Time `json:"response_required_by,omitempty"` // deadline for response
	ResponseAt         *time.Time `json:"response_at,omitempty"`          // nil if no response yet
	Note               string     `json:"note,omitempty"`                 // optional response message
}

func CreateEventTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		attendees TEXT, -- JSON array of attendees
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func (e *Event) Create(db *sql.DB) error {
	return errors.New("Event.Create not implemented")
}

func (e *Event) Update(db *sql.DB) error {
	return errors.New("Event.Update not implemented")
}

func (e *Event) Delete(db *sql.DB) error {
	return errors.New("Event.Delete not implemented")
}

func GetEventByID(db *sql.DB, id int) (*Event, error) {
	return nil, errors.New("GetEventByID not implemented")
}

func GetEventsByDateRange(db *sql.DB, start, end time.Time) ([]*Event, error) {
	return nil, errors.New("GetEventsByDateRange not implemented")
}

func ProcessExpiredInvitations(db *sql.DB) error {
	return errors.New("ProcessExpiredInvitations not implemented")
}