package domain

import "time"

// Event define a entidade do evento no sistema
type Event struct {
	ID          int64     `json:"id"`                             // ID auto-incrementado, não enviado na criação
	Name        string    `json:"name" binding:"required"`        // Nome do evento
	Description string    `json:"description" binding:"required"` // Descrição do evento
	Date        time.Time `json:"date" binding:"required"`        // Data do evento
	CreatedAt   time.Time `json:"created_at"`                     // Data de criação do evento
}
