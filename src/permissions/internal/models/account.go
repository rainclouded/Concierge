package models

type Account struct {
	ID   int    `json:"accountId" binding:"required"`
	Name string `json:"accountName"`
}

type LoginAttempt struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
