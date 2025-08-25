package models

import (
	"database/sql"
	"errors"
	"time"
)

type Event struct {
	ID            int        `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	CreatedBy     string     `json:"created_by"`     // Email of event creator
	CollectionIDs []int      `json:"collection_ids"` // Events can belong to multiple collections
	Attendees     []Attendee `json:"attendees"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Attendee struct {
	Name               string     `json:"name"`
	Email              string     `json:"email"`
	Role               string     `json:"role"`       // "organizer", "required", "optional"
	Status             string     `json:"status"`     // "accepted", "declined", "tentative", "no-response"
	InvitedBy          string     `json:"invited_by"` // email of inviter
	InvitedAt          time.Time  `json:"invited_at"`
	ResponseRequiredBy *time.Time `json:"response_required_by,omitempty"` // deadline for response
	ResponseAt         *time.Time `json:"response_at,omitempty"`          // nil if no response yet
	Note               string     `json:"note,omitempty"`                 // optional response message
}

type Permission string

const (
	PermissionView    Permission = "view"        // Can read event and propose changes for event to organizers and creator
	PermissionEdit    Permission = "contributor" // Can edit event title, location, attendees, description
	PermissionAdmin   Permission = "organizer"   // Can edit event timing, + contributor permissions
	PermissionCreator Permission = "creator"     // Admin permissions on everything
)

type ShareStatus string

const (
	ShareStatusPending  ShareStatus = "pending"
	ShareStatusAccepted ShareStatus = "accepted"
	ShareStatusDeclined ShareStatus = "declined"
	ShareStatusRevoked  ShareStatus = "revoked"
)

type EventCollection struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"` // e.g., "Emma's School Events"
	Description string                 `json:"description"`
	Color       string                 `json:"color"`          // Hex color for UI
	Icon        string                 `json:"icon,omitempty"` // Optional emoji or icon
	CreatedBy   string                 `json:"created_by"`     // Email of creator
	EventIDs    []int                  `json:"event_ids"`      // Events in this collection
	Shares      []EventCollectionShare `json:"shares"`         // Who has access
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type EventCollectionShare struct {
	ID           int         `json:"id"`
	CollectionID int         `json:"collection_id"`
	SharedWith   string      `json:"shared_with"` // Email of person shared with
	Permission   Permission  `json:"permission"`  // Enum for type safety
	SharedBy     string      `json:"shared_by"`   // Email of person who shared
	SharedAt     time.Time   `json:"shared_at"`
	AcceptedAt   *time.Time  `json:"accepted_at,omitempty"`
	Status       ShareStatus `json:"status"`            // Enum for status
	Message      string      `json:"message,omitempty"` // Optional message when sharing
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
