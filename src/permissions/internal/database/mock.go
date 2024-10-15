package database

import (
	"concierge/permissions/internal/models"
	"fmt"
)

type MockDatabase struct {
	permissions []*models.Permission
	groups      []*models.PermissionGroup
}

func NewMockDB() *MockDatabase {
	var permissions = []*models.Permission{
		{ID: 0, Name: "canEditAll", Value: false},
		{ID: 1, Name: "canViewAll", Value: false},
		{ID: 2, Name: "canDelete", Value: false},
		{ID: 3, Name: "canCreate", Value: true},
	}
	db := &MockDatabase{
		permissions: permissions,
		groups: []*models.PermissionGroup{
			{
				ID:          1,
				Name:        "admin",
				Description: "Has all permissions",
				Permissions: permissions,
				Members:     []int{0, 1, 2},
			},
			{
				ID:          2,
				Name:        "editor",
				Description: "Can edit and view",
				Permissions: []*models.Permission{
					permissions[0],
					permissions[1],
				},
				Members: []int{3},
			},
			{
				ID:          3,
				Name:        "viewer",
				Description: "Can only view",
				Permissions: []*models.Permission{
					permissions[1],
				},
				Members: []int{-1, 4, 5},
			},
		},
	}

	return db
}

func (db *MockDatabase) GetPermissions() ([]*models.Permission, error) {
	return db.permissions, nil
}

func (db *MockDatabase) GetPermissionById(permissionId int) (*models.Permission, error) {
	for _, permission := range db.permissions {
		if permission.ID == permissionId {
			return permission, nil
		}
	}

	return nil, fmt.Errorf("Permission with ID %d not found", permissionId)
}

func (db *MockDatabase) CreatePermission(newPermission *models.Permission) error {
	newPermission.ID = db.getMaxPermissoinId()
	db.permissions = append(db.permissions, newPermission)
	return nil
}

func (db *MockDatabase) UpdatePermission(updatedPermission *models.Permission) error {
	for i, permission := range db.permissions {
		if permission.ID == updatedPermission.ID {
			db.permissions[i] = updatedPermission
			return nil
		}
	}

	return fmt.Errorf("Update failed, permission not found with ID %d", updatedPermission.ID)
}

func (db *MockDatabase) GetPermissionGroups() ([]*models.PermissionGroup, error) {
	return db.groups, nil
}

func (db *MockDatabase) GetPermissionGroupById(groupId int) (*models.PermissionGroup, error) {
	for _, group := range db.groups {
		if group.ID == groupId {
			return group, nil
		}
	}

	return nil, fmt.Errorf("Permission Group with ID %d not found", groupId)
}

func (db *MockDatabase) CreatePermissionGroup(newGroup *models.PermissionGroup) error {
	newGroup.ID = db.getMaxGroupId()
	db.groups = append(db.groups, newGroup)
	return nil
}

func (db *MockDatabase) UpdatePermissionGroup(updatedGroup *models.PermissionGroup) error {
	for i, group := range db.groups {
		if group.ID == updatedGroup.ID {
			db.groups[i] = updatedGroup
			return nil
		}
	}

	return fmt.Errorf("Update failed, permission group not found with ID %d", updatedGroup.ID)
}

func (db *MockDatabase) GetGroupMembers(groupId int) ([]int, error) {
	group, err := db.GetPermissionGroupById(groupId)
	if err != nil {
		return nil, err
	}

	return group.Members, nil
}

func (db *MockDatabase) AddMemberToGroup(groupId int, accountId int) error {
	group, err := db.GetPermissionGroupById(groupId)
	if err != nil {
		return err
	}

	group.Members = append(group.Members, accountId)
	return nil
}

func (db *MockDatabase) RemoveMemberFromGroup(groupId int, accountId int) error {
	group, err := db.GetPermissionGroupById(groupId)
	if err != nil {
		return err
	}

	for i, member := range group.Members {
		if member == accountId {
			group.Members = append(group.Members[:i], group.Members[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Remove Failed, Account %d is not a member of group %d", groupId, accountId)
}

func (db *MockDatabase) getMaxPermissoinId() int {
	if len(db.permissions) == 0 {
		return 0
	}

	max := 1
	for _, permission := range db.permissions {
		if permission.ID > max {
			max = permission.ID
		}
	}

	return max + 1
}

func (db *MockDatabase) getMaxGroupId() int {
	if len(db.groups) == 0 {
		return 0
	}

	max := 1
	for _, group := range db.groups {
		if group.ID > max {
			max = group.ID
		}
	}

	return max + 1
}

func (db *MockDatabase) GetPermissionForAccountId(accountId int) ([]*models.Permission, error) {
	userGroups := []*models.PermissionGroup{}
	for _, group := range db.groups {
		for _, member := range group.Members {
			if member == accountId {
				userGroups = append(userGroups, group)
			}
		}
	}

	userPermission := []*models.Permission{}
	for _, permission := range db.permissions {
		permissionState := false
		for _, group := range userGroups {
			if db.GetGroupPermissionState(group, permission.ID) {
				permissionState = true
			}
		}
		userPermission = append(userPermission, &models.Permission{ID: permission.ID, Name: permission.Name, Value: permissionState})
	}

	return userPermission, nil
}

func (db *MockDatabase) GetGroupPermissionState(group *models.PermissionGroup, permissionId int) bool {
	for _, permission := range group.Permissions {
		if permission.ID == permissionId {
			return true
		}
	}
	return false
}
