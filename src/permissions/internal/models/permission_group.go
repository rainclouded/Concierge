package models

type PermissionGroup struct {
	ID          int           `json:"groupId"`
	Name        string        `json:"groupName" binding:"required"`
	Description string        `json:"groupDescription" binding:"required"`
	Permissions []*Permission `json:"groupPermissions"`
	Members     []int         `json:"groupMembers"`
}

type PermissionGroupRequest struct {
	TemplateID    int             `json:"templateId"`
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

func (group *PermissionGroup) Merge(groupReq *PermissionGroupRequest) {
	if groupReq.Name != "" {
		group.Name = groupReq.Name
	}

	if groupReq.Description != "" {
		group.Description = groupReq.Description
	}

	if groupReq.Permissions != nil {
		for _, permission := range groupReq.Permissions {
			found := false
			for _, p := range group.Permissions {
				if p.ID == permission.ID {
					p.Value = permission.State
				}
			}
			if !found {
				group.Permissions = append(group.Permissions, &Permission{
					ID:    permission.ID,
					Value: permission.State,
				})
			}
		}
	}

	if groupReq.Members != nil {
		group.Members = groupReq.Members
	}

	if groupReq.MembersRemove != nil {
		for _, removeId := range groupReq.MembersRemove {
			for i, memberId := range group.Members {
				if memberId == removeId {
					group.Members = append(group.Members[:i], group.Members[i+1:]...)
				}
			}
		}
	}
}
