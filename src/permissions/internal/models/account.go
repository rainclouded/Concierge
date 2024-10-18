package models

type Account struct {
	ID   int    `json:"account-id" binding:"required"`
	Name string `json:"account-name"`
}

type LoginAttempt struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
