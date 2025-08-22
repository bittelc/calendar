package models

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestEventModel(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	t.Run("CreateEvent", func(t *testing.T) {
		event := &Event{
			Title:       "Test Event",
			Description: "A test event",
			StartTime:   time.Now(),
			EndTime:     time.Now().Add(time.Hour),
		}

		err := event.Create(db)
		require.NoError(t, err)
		assert.NotZero(t, event.ID)
		assert.NotZero(t, event.CreatedAt)
		assert.NotZero(t, event.UpdatedAt)
	})

	t.Run("GetEventByID", func(t *testing.T) {
		// Create test event
		event := &Event{
			Title:     "Get Test Event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Hour),
		}
		require.NoError(t, event.Create(db))

		// Get event by ID
		retrieved, err := GetEventByID(db, event.ID)
		require.NoError(t, err)
		assert.Equal(t, event.Title, retrieved.Title)
		assert.Equal(t, event.ID, retrieved.ID)
	})

	t.Run("GetEventByID_NotFound", func(t *testing.T) {
		_, err := GetEventByID(db, 999999)
		assert.Error(t, err)
	})

	t.Run("UpdateEvent", func(t *testing.T) {
		// Create test event
		event := &Event{
			Title:     "Update Test Event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Hour),
		}
		require.NoError(t, event.Create(db))

		// Update event
		event.Title = "Updated Title"
		err := event.Update(db)
		require.NoError(t, err)

		// Verify update
		retrieved, err := GetEventByID(db, event.ID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Title", retrieved.Title)
	})

	t.Run("DeleteEvent", func(t *testing.T) {
		// Create test event
		event := &Event{
			Title:     "Delete Test Event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Hour),
		}
		require.NoError(t, event.Create(db))

		// Delete event
		err := event.Delete(db)
		require.NoError(t, err)

		// Verify deletion
		_, err = GetEventByID(db, event.ID)
		assert.Error(t, err)
	})

	t.Run("GetEventsByDateRange", func(t *testing.T) {
		now := time.Now()
		
		// Create events
		event1 := &Event{
			Title:     "Event 1",
			StartTime: now,
			EndTime:   now.Add(time.Hour),
		}
		event2 := &Event{
			Title:     "Event 2",
			StartTime: now.Add(2 * time.Hour),
			EndTime:   now.Add(3 * time.Hour),
		}
		event3 := &Event{
			Title:     "Event 3",
			StartTime: now.Add(24 * time.Hour), // Next day
			EndTime:   now.Add(25 * time.Hour),
		}
		
		require.NoError(t, event1.Create(db))
		require.NoError(t, event2.Create(db))
		require.NoError(t, event3.Create(db))

		// Get events for today only
		events, err := GetEventsByDateRange(db, now, now.Add(12*time.Hour))
		require.NoError(t, err)
		assert.Len(t, events, 2)
		assert.Contains(t, []string{events[0].Title, events[1].Title}, "Event 1")
		assert.Contains(t, []string{events[0].Title, events[1].Title}, "Event 2")
	})

	t.Run("CreateEventWithAttendees", func(t *testing.T) {
		now := time.Now()
		event := &Event{
			Title:     "Meeting with Attendees",
			StartTime: now,
			EndTime:   now.Add(time.Hour),
			Attendees: []Attendee{
				{
					Name:      "John Doe",
					Email:     "john@example.com",
					Role:      "required",
					Status:    "accepted",
					InvitedBy: "organizer@example.com",
					InvitedAt: now,
				},
				{
					Name:      "Jane Smith",
					Email:     "jane@example.com",
					Role:      "optional",
					Status:    "no-response",
					InvitedBy: "organizer@example.com",
					InvitedAt: now,
				},
			},
		}
		
		err := event.Create(db)
		require.NoError(t, err)
		assert.NotZero(t, event.ID)
		assert.Len(t, event.Attendees, 2)
	})

	t.Run("GetEventWithAttendees", func(t *testing.T) {
		now := time.Now()
		// Create event with attendees
		event := &Event{
			Title:     "Team Meeting",
			StartTime: now,
			EndTime:   now.Add(time.Hour),
			Attendees: []Attendee{
				{
					Name:      "Alice",
					Email:     "alice@example.com",
					Role:      "organizer",
					Status:    "accepted",
					InvitedBy: "alice@example.com",
					InvitedAt: now,
				},
			},
		}
		require.NoError(t, event.Create(db))

		// Retrieve event and verify attendees are loaded
		retrieved, err := GetEventByID(db, event.ID)
		require.NoError(t, err)
		assert.Len(t, retrieved.Attendees, 1)
		assert.Equal(t, "Alice", retrieved.Attendees[0].Name)
		assert.Equal(t, "organizer", retrieved.Attendees[0].Role)
	})

	t.Run("UpdateEventAttendees", func(t *testing.T) {
		now := time.Now()
		responseTime := now.Add(time.Hour)
		
		// Create event with one attendee
		event := &Event{
			Title:     "Status Update Meeting",
			StartTime: now,
			EndTime:   now.Add(time.Hour),
			Attendees: []Attendee{
				{
					Name:      "Charlie",
					Email:     "charlie@example.com",
					Role:      "required",
					Status:    "no-response",
					InvitedBy: "organizer@example.com",
					InvitedAt: now,
				},
			},
		}
		require.NoError(t, event.Create(db))

		// Update attendee status and add response
		event.Attendees[0].Status = "declined"
		event.Attendees[0].ResponseAt = &responseTime
		event.Attendees[0].Note = "Conflict with another meeting"
		
		err := event.Update(db)
		require.NoError(t, err)

		// Verify update
		retrieved, err := GetEventByID(db, event.ID)
		require.NoError(t, err)
		assert.Equal(t, "declined", retrieved.Attendees[0].Status)
		assert.Equal(t, "Conflict with another meeting", retrieved.Attendees[0].Note)
		assert.NotNil(t, retrieved.Attendees[0].ResponseAt)
	})

	t.Run("AutoDeclineExpiredInvitations", func(t *testing.T) {
		now := time.Now()
		pastDeadline := now.Add(-time.Hour) // Deadline was 1 hour ago
		
		// Create event with attendee that has expired response deadline
		event := &Event{
			Title:     "Meeting with Expired Invite",
			StartTime: now.Add(24 * time.Hour), // Future meeting
			EndTime:   now.Add(25 * time.Hour),
			Attendees: []Attendee{
				{
					Name:               "Bob",
					Email:              "bob@example.com",
					Role:               "required",
					Status:             "no-response",
					InvitedBy:          "organizer@example.com",
					InvitedAt:          now.Add(-48 * time.Hour), // Invited 2 days ago
					ResponseRequiredBy: &pastDeadline,           // Deadline was 1 hour ago
				},
				{
					Name:               "Alice",
					Email:              "alice@example.com", 
					Role:               "optional",
					Status:             "accepted",
					InvitedBy:          "organizer@example.com",
					InvitedAt:          now.Add(-24 * time.Hour),
					ResponseRequiredBy: nil, // No deadline
				},
			},
		}
		require.NoError(t, event.Create(db))

		// Process expired invitations
		err := ProcessExpiredInvitations(db)
		require.NoError(t, err)

		// Verify that Bob's invitation was auto-declined
		retrieved, err := GetEventByID(db, event.ID)
		require.NoError(t, err)
		
		var bob *Attendee
		var alice *Attendee
		for i := range retrieved.Attendees {
			if retrieved.Attendees[i].Email == "bob@example.com" {
				bob = &retrieved.Attendees[i]
			}
			if retrieved.Attendees[i].Email == "alice@example.com" {
				alice = &retrieved.Attendees[i]
			}
		}
		
		require.NotNil(t, bob)
		require.NotNil(t, alice)
		
		// Bob should be auto-declined
		assert.Equal(t, "declined", bob.Status)
		assert.NotNil(t, bob.ResponseAt)
		assert.Equal(t, "Auto-declined: Response deadline exceeded", bob.Note)
		
		// Alice should remain unchanged
		assert.Equal(t, "accepted", alice.Status)
	})
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	err = CreateEventTable(db)
	require.NoError(t, err)

	return db
}