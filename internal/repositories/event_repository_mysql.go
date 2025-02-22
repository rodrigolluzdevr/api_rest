package repositories

import (
	"database/sql"
	"fmt"
	"rodrigolluzdevr/api_rest/internal/db"
	"rodrigolluzdevr/api_rest/internal/domain"
)

// EventRepository defines the methods to interect with the database
type EventRepository interface {
	CreateEvent(event *domain.Event) error
}

// EventRepositoryMySQL implement EventRepository to MySQL
type EventRepositoryMySQL struct {
	db *sql.DB
}

// NewEventRepository creates a new event repository
func NewEventRepository() EventRepository {
	return &EventRepositoryMySQL{
		db: db.GetDB(), // geting the database connection
	}
}

// CreateEvent insert an event in database
func (r *EventRepositoryMySQL) CreateEvent(event *domain.Event) error {
	query := `INSERT INTO events (name, description, date) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, event.Name, event.Description, event.Date)
	if err != nil {
		return fmt.Errorf("erro ao inserir evento: %v", err)
	}
	return nil
}
