package domain

type USER struct {
	ID       int64  `json: "id"`
	NAME     string `json:"name" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}
