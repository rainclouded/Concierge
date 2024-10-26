package database

// import (
// 	"concierge/permissions/internal/models"
// 	"fmt"
// )

// type MariaDatabase struct {
// 	permissions []*models.Permission
// 	groups      []*models.PermissionGroup
// }

// // //
// // // NOT IMPLEMENTED DO NOT USE !!!
// // //
// func NewMariaDB() *MariaDatabase {
// 	return nil
// }

// func (db *MariaDatabase) GetPermissions() ([]*models.Permission, error) {
// 	return db.permissions, nil
// }

// func (db *MariaDatabase) GetPermissionById(permissionId int) (*models.Permission, error) {
// 	for _, permission := range db.permissions {
// 		if permission.ID == permissionId {
// 			return permission, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("permission with ID %d not found", permissionId)
// }

// func (db *MariaDatabase) CreatePermission(newPermission *models.Permission) error {
// 	db.permissions = append(db.permissions, newPermission)
// 	return nil
// }

// func (db *MariaDatabase) UpdatePermission(updatedPermission *models.Permission) error {
// 	for i, permission := range db.permissions {
// 		if permission.ID == updatedPermission.ID {
// 			db.permissions[i] = updatedPermission
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("update failed, permission not found with ID %d", updatedPermission.ID)
// }

// func (db *MariaDatabase) GetPermissionGroups() ([]*models.PermissionGroup, error) {
// 	return db.groups, nil
// }

// func (db *MariaDatabase) GetPermissionGroupById(groupId int) (*models.PermissionGroup, error) {
// 	for _, group := range db.groups {
// 		if group.ID == groupId {
// 			return group, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("permission Group with ID %d not found", groupId)
// }

// func (db *MariaDatabase) CreatePermissionGroup(newGroup *models.PermissionGroup) error {
// 	db.groups = append(db.groups, newGroup)
// 	return nil
// }

// func (db *MariaDatabase) UpdatePermissionGroup(updatedGroup *models.PermissionGroupRequest) error {

// 	return fmt.Errorf("update failed, permission group not found with ID")
// }

// func (db *MariaDatabase) GetGroupMembers(groupId int) ([]int, error) {
// 	group, err := db.GetPermissionGroupById(groupId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return group.Members, nil
// }

// func (db *MariaDatabase) AddMemberToGroup(groupId int, accountId int) error {
// 	group, err := db.GetPermissionGroupById(groupId)
// 	if err != nil {
// 		return err
// 	}

// 	group.Members = append(group.Members, accountId)
// 	return nil
// }

// func (db *MariaDatabase) RemoveMemberFromGroup(groupId int, accountId int) error {
// 	group, err := db.GetPermissionGroupById(groupId)
// 	if err != nil {
// 		return err
// 	}

// 	for i, member := range group.Members {
// 		if member == accountId {
// 			group.Members = append(group.Members[:i], group.Members[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("remove Failed, Account %d is not a member of group %d", groupId, accountId)
// }

// func (db *MariaDatabase) GetPermissionForAccountId(accountId int) ([]*models.Permission, error) {
// 	userGroups := []*models.PermissionGroup{}
// 	for _, group := range db.groups {
// 		for _, member := range group.Members {
// 			if member == accountId {
// 				userGroups = append(userGroups, group)
// 			}
// 		}
// 	}

// 	userPermission := []*models.Permission{}
// 	for _, permission := range db.permissions {
// 		permissionState := false
// 		for _, group := range userGroups {
// 			if db.GetGroupPermissionState(group, permission.ID) {
// 				permissionState = true
// 			}
// 		}
// 		userPermission = append(userPermission, &models.Permission{ID: permission.ID, Name: permission.Name, Value: permissionState})
// 	}

// 	return userPermission, nil
// }

// func (db *MariaDatabase) GetGroupPermissionState(group *models.PermissionGroup, permissionId int) bool {
// 	for _, permission := range group.Permissions {
// 		if permission.ID == permissionId {
// 			return true
// 		}
// 	}
// 	return false
// }
