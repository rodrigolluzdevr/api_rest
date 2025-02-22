package services

import (
	"errors"
	"rodrigolluzdevr/api_rest/internal/domain"
	"rodrigolluzdevr/api_rest/internal/repositories"
)

// EventService lida com as regras de negócios dos eventos
type EventService struct {
	eventRepo repositories.EventRepository
}

// NewEventService cria uma nova instância do serviço de eventos
func NewEventService() *EventService {
	return &EventService{
		eventRepo: repositories.NewEventRepository(),
	}
}

// CreateEvent cria um novo evento no banco de dados
func (s *EventService) CreateEvent(event *domain.Event) error {
	err := s.eventRepo.CreateEvent(event)
	if err != nil {
		return errors.New("failed to create event")
	}
	return nil
}
