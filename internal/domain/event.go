package domain

type Event struct {
	ID          int64  `json: "id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"email" binding:"required"`
	Date        string `json:"password" binding:"required"`
}
