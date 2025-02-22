package domain

type EVENT struct {
	ID          int64  `json: "id"`
	NAME        string `json:"name" binding:"required"`
	DESCRIPTION string `json:"email" binding:"required"`
	DATE        string `json:"password" binding:"required"`
}
