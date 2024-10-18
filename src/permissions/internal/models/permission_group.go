package models

type PermissionGroup struct {
	ID          int           `json:"group-id"`
	Name        string        `json:"group-name" binding:"required"`
	Description string        `json:"group-description" binding:"required"`
	Permissions []*Permission `json:"group-permissions"`
	Members     []int         `json:"group-members"`
}
