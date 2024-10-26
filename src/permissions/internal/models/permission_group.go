package models

import (
	"fmt"
	"strings"
)

type PermissionGroup struct {
	ID          int           `json:"groupId"`
	Name        string        `json:"groupName" binding:"required"`
	Description string        `json:"groupDescription" binding:"required"`
	Permissions []*Permission `json:"groupPermissions"`
	Members     []int         `json:"groupMembers"`
}

type PermissionGroupRequest struct {
	Name          string          `json:"groupName"`
	Description   string          `json:"groupDescription"`
	Permissions   []*PermissionId `json:"groupPermissions"`
	Members       []int           `json:"groupMembers"`
	MembersRemove []int           `json:"removeGroupMembers"`
}

type PermissionId struct {
	ID    int  `json:"permissionId"`
	State bool `json:"state"`
}

func (group *PermissionGroup) DeepCopy() *PermissionGroup {
	newGroup := &PermissionGroup{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		Permissions: []*Permission{},
		Members:     group.Members,
	}

	for _, permission := range group.Permissions {
		newGroup.Permissions = append(newGroup.Permissions, permission.DeepCopy())
	}

	return newGroup
}

func (pg PermissionGroup) String() string {
	var permissions []string
	for _, perm := range pg.Permissions {
		permissions = append(permissions, perm.String())
	}

	return fmt.Sprintf(
		"PermissionGroup{ID: %d, Name: %s, Description: %s, Permissions: [%s], Members: %v}",
		pg.ID,
		pg.Name,
		pg.Description,
		strings.Join(permissions, ", "),
		pg.Members,
	)
}
