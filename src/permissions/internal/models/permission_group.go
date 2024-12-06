package models

import (
	"fmt"
	"strings"
)

// PermissionGroup represents a group of permissions that can be assigned to users
// It contains metadata about the group, a list of associated permissions, and group members
type PermissionGroup struct {
	ID          int           `json:"groupId"`            // Unique identifier for the permission group
	Name        string        `json:"groupName" binding:"required"`  // Name of the permission group, required
	Description string        `json:"groupDescription" binding:"required"`  // Description of the permission group, required
	Permissions []*Permission `json:"groupPermissions"`   // List of permissions associated with the group
	Members     []int         `json:"groupMembers"`       // List of user IDs who are members of the group
}

// PermissionGroupRequest represents the structure of a request to create or update a PermissionGroup
// It contains the group's name, description, permissions, and members (including those to be removed)
type PermissionGroupRequest struct {
	Name          string          `json:"groupName"`          // The name of the permission group
	Description   string          `json:"groupDescription"`   // The description of the permission group
	Permissions   []*PermissionId `json:"groupPermissions"`   // The permissions to be assigned to the group
	Members       []int           `json:"groupMembers"`       // List of member IDs to be added to the group
	MembersRemove []int           `json:"removeGroupMembers"` // List of member IDs to be removed from the group
}

// PermissionId represents a permission ID with an associated state (enabled/disabled)
type PermissionId struct {
	ID    int  `json:"permissionId"`  // The unique identifier for the permission
	State bool `json:"state"`         // The state of the permission (enabled/disabled)
}

// DeepCopy creates a deep copy of the PermissionGroup, including its permissions
// This is useful to avoid modifying the original PermissionGroup when making changes
// Returns:
//   *PermissionGroup: A deep copy of the original permission group
func (group *PermissionGroup) DeepCopy() *PermissionGroup {
	// Create a new PermissionGroup instance
	newGroup := &PermissionGroup{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		Permissions: []*Permission{}, // Empty slice for the copied permissions
		Members:     group.Members,   // Shallow copy of the members slice (IDs)
	}

	// Copy each permission from the original group to the new group
	for _, permission := range group.Permissions {
		newGroup.Permissions = append(newGroup.Permissions, permission.DeepCopy())
	}

	return newGroup // Return the deep-copied permission group
}

// String provides a string representation of the PermissionGroup
// This is helpful for debugging or logging the permission group's details
// Returns:
//   string: A formatted string describing the permission group
func (pg PermissionGroup) String() string {
	// Create a slice of string representations for the permissions in the group
	var permissions []string
	for _, perm := range pg.Permissions {
		permissions = append(permissions, perm.String()) // Append each permission's string representation
	}

	// Format and return a string representation of the permission group
	return fmt.Sprintf(
		"PermissionGroup{ID: %d, Name: %s, Description: %s, Permissions: [%s], Members: %v}",
		pg.ID,                          // Permission group ID
		pg.Name,                        // Permission group name
		pg.Description,                 // Permission group description
		strings.Join(permissions, ", "), // Joined string of permission representations
		pg.Members,                     // List of member IDs
	)
}
