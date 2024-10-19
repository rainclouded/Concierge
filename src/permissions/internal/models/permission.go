package models

type PermissionTemplate struct {
	ID      int    `json:"permission-id"`
	Name    string `json:"permission-name" binding:"required"`
	Default bool   `json:"permission-default-state" binding:"required"`
}

type Permission struct {
	ID    int    `json:"permission-id"`
	Name  string `json:"permission-name" binding:"required"`
	Value bool   `json:"permission-state" binding:"required"`
}

func (permission *Permission) DeepCopy() *Permission {
	return &Permission{
		ID:    permission.ID,
		Name:  permission.Name,
		Value: permission.Value,
	}
}
