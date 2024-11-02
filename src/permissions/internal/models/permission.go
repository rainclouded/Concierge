package models

import "fmt"

type PermissionTemplate struct {
	ID      int    `json:"permissionId"`
	Name    string `json:"permissionName" binding:"required"`
	Default bool   `json:"permissionDefaultState" binding:"required"`
}

type Permission struct {
	ID    int    `json:"permissionId"`
	Name  string `json:"permissionName" binding:"required"`
	Value bool   `json:"permissionState" binding:"required"`
}

type PermissionPostRequest struct {
	Name string `json:"permissionName" binding:"required"`
}

func (permission *Permission) DeepCopy() *Permission {
	return &Permission{
		ID:    permission.ID,
		Name:  permission.Name,
		Value: permission.Value,
	}
}

func (p Permission) String() string {
	return fmt.Sprintf("Permission{ID: %d, Name: %s, Value: %t}", p.ID, p.Name, p.Value)
}
