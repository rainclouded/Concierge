package models

type Account struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"username"`
}

type LoginAttempt struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountResponse struct {
	Message     string   `json:"message" binding:"required"`
	Status      string   `json:"status" binding:"required"`
	AccountData *Account `json:"data"`
}
