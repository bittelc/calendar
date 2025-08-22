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
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	err = CreateEventTable(db)
	require.NoError(t, err)

	return db
}