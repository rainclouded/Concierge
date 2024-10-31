package database

import (
	"concierge/permissions/internal/models"
)

type Database interface {
	GetPermissions() ([]*models.Permission, error)
	GetPermissionById(int) (*models.Permission, error)
	CreatePermission(string) (*models.Permission, error)
	UpdatePermission(*models.Permission) error
	GetPermissionGroups() ([]*models.PermissionGroup, error)
	GetPermissionGroupById(int) (*models.PermissionGroup, error)
	CreatePermissionGroup(*models.PermissionGroupRequest) error
	UpdatePermissionGroup(int, *models.PermissionGroupRequest) error
	GetGroupMembers(int) ([]int, error)
	AddMemberToGroup(int, int) error
	RemoveMemberFromGroup(int, int) error
	GetPermissionForAccountId(int) ([]*models.Permission, error)
}
