package database

import (
	"concierge/permissions/internal/constants"
	"concierge/permissions/internal/models"
	"fmt"
)

// MockDatabase is a mock implementation of a database that holds permissions and permission groups
// for testing purposes. It simulates operations such as creating, retrieving, updating, and deleting permissions and permission groups.
// Args:
//    None
// Returns:
//    None
type MockDatabase struct {
	permissions []*models.Permission
	groups      []*models.PermissionGroup
}

// NewMockDB creates a new instance of MockDatabase with predefined permissions and permission groups
// Args:
//    None
// Returns:
//    *MockDatabase: Returns a new instance of MockDatabase populated with test data
func NewMockDB() *MockDatabase {
	var permissions = []*models.Permission{
		{ID: 1, Name: constants.CanViewPermissionGroups, Value: true},
		{ID: 2, Name: constants.CanEditPermissionGroups, Value: true},
		{ID: 3, Name: constants.CanViewPermissions, Value: true},
		{ID: 4, Name: constants.CanEditPermissions, Value: true},
		{ID: 5, Name: constants.CanViewAmenities, Value: true},
		{ID: 6, Name: constants.CanEditAmenities, Value: true},
		{ID: 7, Name: constants.CanDeleteAmenities, Value: true},
		{ID: 8, Name: constants.CanDeleteGuestsAccounts, Value: true},
		{ID: 9, Name: constants.CanDeleteStaffAccounts, Value: true},
		{ID: 10, Name: constants.CanEditStaffAccounts, Value: true},
		{ID: 11, Name: constants.CanEditGuestAccounts, Value: true},
		{ID: 12, Name: constants.CanViewIncidentReports, Value: true},
		{ID: 13, Name: constants.CanEditIncidentReports, Value: true},
		{ID: 14, Name: constants.CanCreateIncidentReports, Value: true},
		{ID: 15, Name: constants.CanDeleteIncidentReports, Value: true},
		{ID: 16, Name: constants.CanViewTasks, Value: true},
		{ID: 17, Name: constants.CanCreateTasks, Value: true},
		{ID: 18, Name: constants.CanEditTasks, Value: true},
		{ID: 19, Name: constants.CanDeleteTasks, Value: true},
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
				Description: "Can edit and view most data",
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
					permissions[0].DeepCopy(),
				},
				Members: []int{-1, 4, 5},
			},
		},
	}

	return db
}

// GetPermissions retrieves all permissions stored in the mock database.
// Args:
//    None
// Returns:
//    []*models.Permission: A slice of all permission objects stored in the mock database.
//    error: Returns any error that occurs during the operation.
func (db *MockDatabase) GetPermissions() ([]*models.Permission, error) {
	return db.permissions, nil
}

// GetPermissionById retrieves a permission by its ID from the mock database.
// Args:
//    permissionId: The ID of the permission to retrieve.
// Returns:
//    *models.Permission: The permission with the specified ID.
//    error: Returns an error if the permission is not found.
func (db *MockDatabase) GetPermissionById(permissionId int) (*models.Permission, error) {
	for _, permission := range db.permissions {
		if permission.ID == permissionId {
			return permission.DeepCopy(), nil
		}
	}

	return nil, fmt.Errorf("permission with ID %d not found", permissionId)
}

// CreatePermission creates a new permission in the mock database.
// Args:
//    permissionName: The name of the new permission to create.
// Returns:
//    *models.Permission: The newly created permission object.
//    error: Returns an error if a permission with the same name already exists.
func (db *MockDatabase) CreatePermission(permissionName string) (*models.Permission, error) {
	permission := &models.Permission{ID: db.getMaxPermissionId(), Name: permissionName, Value: true}
	for _, p := range db.permissions {
		if p.Name == permissionName {
			return nil, fmt.Errorf("conflict")
		}
	}

	db.permissions = append(db.permissions, permission)
	return permission, nil
}

// UpdatePermission updates an existing permission in the mock database.
// Args:
//    updatedPermission: A pointer to the updated permission object.
// Returns:
//    error: Returns an error if the permission with the specified ID is not found.
func (db *MockDatabase) UpdatePermission(updatedPermission *models.Permission) error {
	for i, permission := range db.permissions {
		if permission.ID == updatedPermission.ID {
			db.permissions[i] = updatedPermission
			return nil
		}
	}

	return fmt.Errorf("update failed, permission not found with ID %d", updatedPermission.ID)
}

// GetPermissionGroups retrieves all permission groups stored in the mock database.
// Args:
//    None
// Returns:
//    []*models.PermissionGroup: A slice of all permission group objects stored in the mock database.
//    error: Returns any error that occurs during the operation.
func (db *MockDatabase) GetPermissionGroups() ([]*models.PermissionGroup, error) {
	return db.groups, nil
}

// GetPermissionGroupById retrieves a permission group by its ID from the mock database.
// Args:
//    groupId: The ID of the permission group to retrieve.
// Returns:
//    *models.PermissionGroup: The permission group with the specified ID.
//    error: Returns an error if the permission group is not found.
func (db *MockDatabase) GetPermissionGroupById(groupId int) (*models.PermissionGroup, error) {
	for _, group := range db.groups {
		if group.ID == groupId {
			return group, nil
		}
	}

	return nil, fmt.Errorf("permission Group with ID %d not found", groupId)
}

func (db *MockDatabase) GetPermissionGroupsByAccount(accountId int) ([]*models.PermissionGroup, error) {
	var groupsWithAcc []*models.PermissionGroup

	for _, group := range db.groups {
		if contains(group.Members, accountId) {
			groupsWithAcc = append(groupsWithAcc, group)
		}
	}

	return groupsWithAcc, nil
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

// UpdatePermissionGroup updates an existing permission group in the mock database.
// Args:
//    id: The ID of the permission group to update.
//    groupReq: The request object containing the new details for the group.
// Returns:
//    error: Returns an error if the update fails.
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

// GetGroupMembers retrieves the list of member IDs for a given permission group.
// Args:
//    groupId: The ID of the permission group.
// Returns:
//    []int: A slice of member IDs belonging to the specified group.
//    error: Returns an error if the group is not found.
func (db *MockDatabase) GetGroupMembers(groupId int) ([]int, error) {
	group, err := db.GetPermissionGroupById(groupId)
	if err != nil {
		return nil, err
	}

	return group.Members, nil
}

// AddMemberToGroup adds a member to a given permission group.
// Args:
//    groupId: The ID of the permission group.
//    accountId: The ID of the account to add as a member.
// Returns:
//    error: Returns an error if the account cannot be added to the group.
func (db *MockDatabase) AddMemberToGroup(groupId int, accountId int) error {
	group, err := db.GetPermissionGroupById(groupId)
	if err != nil {
		return err
	}

	group.Members = append(group.Members, accountId)
	return nil
}

// RemoveMemberFromGroup removes a member from a given permission group.
// Args:
//    groupId: The ID of the permission group.
//    accountId: The ID of the account to remove from the group.
// Returns:
//    error: Returns an error if the account cannot be removed from the group.
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

// getMaxPermissionId calculates the maximum permission ID from the existing permissions.
// Args:
//    None
// Returns:
//    int: The next available ID for a new permission.
func (db *MockDatabase) getMaxPermissionId() int {
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

// getMaxGroupId calculates the maximum group ID from the existing groups.
// Args:
//    None
// Returns:
//    int: The next available ID for a new group.
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

// GetPermissionForAccountId retrieves the permissions for a given account ID based on their group memberships.
// Args:
//    accountId: The ID of the account to retrieve permissions for.
// Returns:
//    []*models.Permission: A slice of permissions for the specified account.
//    error: Returns an error if the account is not found or permissions cannot be determined.
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

// GetGroupPermissionState checks if a permission is assigned to a group.
// Args:
//    group: The permission group to check.
//    permissionId: The ID of the permission.
// Returns:
//    bool: Returns true if the permission is assigned to the group, otherwise false.
func (db *MockDatabase) GetGroupPermissionState(group *models.PermissionGroup, permissionId int) bool {
	for _, permission := range group.Permissions {
		if permission.ID == permissionId {
			return true
		}
	}
	return false
}

// ClearPermissions clears all permissions from the mock database.
// Args:
//    None
// Returns:
//    None
func (db *MockDatabase) ClearPermissions() {
	db.permissions = []*models.Permission{}
}

func contains(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
