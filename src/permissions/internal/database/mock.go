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
		{ID: 0, Name: "canEditAll", Value: true},
		{ID: 1, Name: "canViewAll", Value: true},
		{ID: 2, Name: "canDelete", Value: true},
		{ID: 3, Name: "canCreate", Value: true},
	}
	db := &MockDatabase{
		permissions: permissions,
		groups: []*models.PermissionGroup{
			{
				ID:          1,
				Name:        "admin",
				Description: "Has all permissions",
				Permissions: []*models.Permission{
					permissions[0].DeepCopy(),
					permissions[1].DeepCopy(),
					permissions[2].DeepCopy(),
					permissions[3].DeepCopy(),
				},
				Members: []int{0, 1, 2},
			},
			{
				ID:          2,
				Name:        "editor",
				Description: "Can edit and view",
				Permissions: []*models.Permission{
					permissions[0].DeepCopy(),
					permissions[1].DeepCopy(),
				},
				Members: []int{3},
			},
			{
				ID:          3,
				Name:        "viewer",
				Description: "Can only view",
				Permissions: []*models.Permission{
					permissions[1].DeepCopy(),
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
			return permission.DeepCopy(), nil
		}
	}

	return nil, fmt.Errorf("permission with ID %d not found", permissionId)
}

func (db *MockDatabase) CreatePermission(permissionName string) (*models.Permission, error) {
	permission := &models.Permission{ID: db.getMaxPermissoinId(), Name: permissionName, Value: true}
	for _, p := range db.permissions {
		if p.Name == permissionName {
			return nil, fmt.Errorf("conflict")
		}
	}

	db.permissions = append(db.permissions, permission)
	return permission, nil
}

func (db *MockDatabase) UpdatePermission(updatedPermission *models.Permission) error {
	for i, permission := range db.permissions {
		if permission.ID == updatedPermission.ID {
			db.permissions[i] = updatedPermission
			return nil
		}
	}

	return fmt.Errorf("update failed, permission not found with ID %d", updatedPermission.ID)
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

	return nil, fmt.Errorf("permission Group with ID %d not found", groupId)
}

func (db *MockDatabase) CreatePermissionGroup(newGroup *models.PermissionGroupRequest) error {
	newGroupObj := models.PermissionGroup{
		ID:          db.getMaxGroupId(),
		Name:        newGroup.Name,
		Description: newGroup.Description,
	}

	for _, p := range newGroup.Permissions {
		permRef, err := db.GetPermissionById(p.ID)
		if err != nil {
			return fmt.Errorf("could not find permission %d", p.ID)
		}
		permRef = permRef.DeepCopy()
		permRef.Value = p.State

		newGroupObj.Permissions = append(newGroupObj.Permissions, permRef)
	}

	newGroupObj.Members = append(newGroupObj.Members, newGroup.Members...)

	db.groups = append(db.groups, &newGroupObj)
	return nil
}

func (db *MockDatabase) UpdatePermissionGroup(id int, groupReq *models.PermissionGroupRequest) error {
	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		return err
	}
	group = group.DeepCopy()

	if groupReq.Name != "" {
		group.Name = groupReq.Name
	}

	if groupReq.Description != "" {
		group.Description = groupReq.Description
	}

	if groupReq.Name != "" {
		group.Name = groupReq.Name
	}

	if groupReq.Permissions != nil {
		newPermSlice := []*models.Permission{}
		for _, newPermission := range groupReq.Permissions {
			newPermObj, err := db.GetPermissionById(newPermission.ID)
			if err != nil {
				return err
			}
			newPermObj.Value = newPermission.State
			newPermSlice = append(newPermSlice, newPermObj.DeepCopy())
		}

		for _, newPerm := range newPermSlice {
			found := false
			for _, p := range group.Permissions {
				if p.ID == newPerm.ID {
					p.Value = newPerm.Value
					found = true
				}
			}
			if !found {
				group.Permissions = append(group.Permissions, newPerm)
			}
		}
	}

	if groupReq.Members != nil && groupReq.MembersRemove != nil {
		for _, addM := range groupReq.Members {
			for _, removeM := range groupReq.MembersRemove {
				if addM == removeM {
					return fmt.Errorf("cannot add and remove the same group member: %d", addM)
				}
			}
		}
	}

	if groupReq.Members != nil {
		for _, addM := range groupReq.Members {
			found := false
			for _, m := range group.Members {
				if m == addM {
					found = true
				}
			}

			if !found {
				group.Members = append(group.Members, addM)
			}
		}
	}

	if groupReq.MembersRemove != nil {
		for _, removeM := range groupReq.MembersRemove {
			for index, m := range group.Members {
				if m == removeM {
					group.Members = append(group.Members[:index], group.Members[index+1:]...)
				}
			}
		}
	}

	for i, v := range db.groups {
		if v.ID == group.ID {
			db.groups[i] = group
		}
	}
	return nil
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

	return fmt.Errorf("remove Failed, Account %d is not a member of group %d", groupId, accountId)
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

// Testing setup methods
func (db *MockDatabase) ClearPermissions() {
	db.permissions = []*models.Permission{}
}
