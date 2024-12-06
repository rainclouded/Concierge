package models

import "fmt"

// PermissionTemplate represents a template for a permission with an ID, name, and default state.
// It is used to define permission types and their default states in a system.
type PermissionTemplate struct {
	ID      int    `json:"permissionId"`           // Unique identifier for the permission template
	Name    string `json:"permissionName" binding:"required"`  // Name of the permission template, required field
	Default bool   `json:"permissionDefaultState" binding:"required"` // The default state of the permission (enabled/disabled), required field
}

// Permission represents a specific permission assigned to an account or user.
// It contains the ID, name, and current state (enabled or disabled) of the permission.
type Permission struct {
	ID    int    `json:"permissionId"`    // Unique identifier for the permission
	Name  string `json:"permissionName" binding:"required"`  // Name of the permission, required field
	Value bool   `json:"permissionState" binding:"required"` // The state of the permission (enabled or disabled), required field
}

// PermissionPostRequest is the structure used for creating a new permission.
// It contains only the name of the permission being created.
type PermissionPostRequest struct {
	Name string `json:"permissionName" binding:"required"` // The name of the permission to be created, required field
}

// DeepCopy creates a new instance of the Permission with the same values as the original permission.
// This is useful for ensuring that any modifications to the copied permission do not affect the original.
func (permission *Permission) DeepCopy() *Permission {
	// Create a new Permission instance with the same values as the original
	return &Permission{
		ID:    permission.ID,
		Name:  permission.Name,
		Value: permission.Value,
	}
}

// String provides a string representation of the Permission.
// This is useful for debugging or logging purposes to print the details of the permission.
func (p Permission) String() string {
	// Return a formatted string representing the Permission
	return fmt.Sprintf("Permission{ID: %d, Name: %s, Value: %t}", p.ID, p.Name, p.Value)
}
