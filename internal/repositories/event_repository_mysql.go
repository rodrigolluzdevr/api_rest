package repositories

import (
	"database/sql"
	"fmt"
	"rodrigolluzdevr/api_rest/internal/db"
	"rodrigolluzdevr/api_rest/internal/domain"
)

// EventRepository define os métodos para interagir com o banco de dados
type EventRepository interface {
	CreateEvent(event *domain.Event) error
}

// EventRepositoryMySQL implementa EventRepository para MySQL
type EventRepositoryMySQL struct {
	db *sql.DB
}

// NewEventRepository cria um novo repositório de eventos
func NewEventRepository() EventRepository {
	return &EventRepositoryMySQL{
		db: db.GetDB(), // Obtendo a conexão do banco
	}
}

// CreateEvent insere um evento no banco de dados
func (r *EventRepositoryMySQL) CreateEvent(event *domain.Event) error {
	query := `INSERT INTO events (name, description, date) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, event.Name, event.Description, event.Date)
	if err != nil {
		return fmt.Errorf("erro ao inserir evento: %v", err)
	}
	return nil
}
