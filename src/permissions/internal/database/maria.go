package database

import (
	"concierge/permissions/internal/models"
	"database/sql"
	"fmt"
	"sort"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDB struct {
	db         *sql.DB
	dataSource string
	forTesting bool
}

func NewMariaDB(dataSourceName string, forTesting bool) (*MariaDB, error) {
	return &MariaDB{db: nil, dataSource: dataSourceName, forTesting: forTesting}, nil
}

func (m *MariaDB) Close() error {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			return err
		}

		m.db = nil
		return nil
	}

	return fmt.Errorf("connection already closed")
}

func (m *MariaDB) setupConnection() error {
	if m.db == nil {
		db, err := sql.Open("mysql", m.dataSource)
		if err != nil {
			m.db = nil
			return err
		}
		m.db = db
	}
	if err := m.db.Ping(); err != nil {
		m.db = nil
		return err
	}
	return nil
}

func (m *MariaDB) GetPermissions() ([]*models.Permission, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Query("SELECT id, name FROM Permissions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		var p models.Permission
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		permissions = append(permissions, &p)
	}
	return permissions, nil
}

func (m *MariaDB) GetPermissionById(id int) (*models.Permission, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	var p models.Permission
	row := m.db.QueryRow("SELECT id, name FROM Permissions WHERE id = ?", id)
	err = row.Scan(&p.ID, &p.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permission not found")
		}
		return nil, err
	}
	// if p.ID == nil {
	// 	return nil, fmt.Errorf("Permission not found")
	// }
	return &p, nil
}

func (m *MariaDB) CreatePermission(name string) (*models.Permission, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := tx.Exec("INSERT INTO Permissions (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if !m.forTesting {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return &models.Permission{ID: int(id), Name: name}, nil
}

func (m *MariaDB) UpdatePermission(p *models.Permission) error {
	err := m.setupConnection()
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE Permissions SET name = ? WHERE id = ?", p.Name, p.ID)

	if !m.forTesting {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

func (m *MariaDB) GetPermissionGroups() ([]*models.PermissionGroup, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	groupQuery := `
SELECT 
	pg.id AS groupId,
	pg.name AS groupName,
	pg.description AS groupDescription
FROM 
	PermissionGroups pg;
	`

	rows, err := m.db.Query(groupQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groupsMap := make(map[int]*models.PermissionGroup)

	for rows.Next() {
		var g models.PermissionGroup
		if err := rows.Scan(&g.ID, &g.Name, &g.Description); err != nil {
			return nil, err
		}
		groupsMap[g.ID] = &g
	}

	memberQuery := `
SELECT 
	gm.groupId,
	gm.memberId
FROM 
	GroupMembers gm;
	`

	memberRows, err := m.db.Query(memberQuery)
	if err != nil {
		return nil, err
	}
	defer memberRows.Close()

	for memberRows.Next() {
		var groupId, memberId int
		if err := memberRows.Scan(&groupId, &memberId); err != nil {
			return nil, err
		}
		if group, exists := groupsMap[groupId]; exists {
			group.Members = append(group.Members, memberId)
		}
	}

	permissionQuery := `
SELECT 
	gp.groupId,
	p.id AS permissionId,
	p.name AS permissionName,
	gp.value as permissionValue
FROM 
	GroupPermissions gp
LEFT JOIN 
	Permissions p ON gp.permission_id = p.id;
	`

	permissionRows, err := m.db.Query(permissionQuery)
	if err != nil {
		return nil, err
	}
	defer permissionRows.Close()

	for permissionRows.Next() {
		var groupId, permissionId int
		var permissionName string
		var permissionValue bool
		if err := permissionRows.Scan(&groupId, &permissionId, &permissionName, &permissionValue); err != nil {
			return nil, err
		}
		if group, exists := groupsMap[groupId]; exists && permissionId != 0 {
			perm := &models.Permission{
				ID:    permissionId,
				Name:  permissionName,
				Value: permissionValue,
			}
			group.Permissions = append(group.Permissions, perm)
		}
	}

	var groups []*models.PermissionGroup
	for _, group := range groupsMap {
		groups = append(groups, group)
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].ID < groups[j].ID
	})

	return groups, nil
}
func (m *MariaDB) GetPermissionGroupById(id int) (*models.PermissionGroup, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	groupQuery := `
SELECT 
	pg.id AS groupId,
	pg.name AS groupName,
	pg.description AS groupDescription
FROM 
	PermissionGroups pg
WHERE 
	pg.id = ?;
		`

	var group models.PermissionGroup
	err = m.db.QueryRow(groupQuery, id).Scan(&group.ID, &group.Name, &group.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permission group not found")
		}
		return nil, err
	}

	memberQuery := `
SELECT 
	gm.memberId
FROM 
	GroupMembers gm
WHERE 
	gm.groupId = ?
ORDER BY 
	gm.memberId;
		`

	memberRows, err := m.db.Query(memberQuery, id)
	if err != nil {
		return nil, err
	}
	defer memberRows.Close()

	for memberRows.Next() {
		var memberId int
		if err := memberRows.Scan(&memberId); err != nil {
			return nil, err
		}
		group.Members = append(group.Members, memberId)
	}

	permissionQuery := `
SELECT 
	p.id AS permissionId,
	p.name AS permissionName,
	gp.value AS permissionValue
FROM 
	GroupPermissions gp
LEFT JOIN 
	Permissions p ON gp.permission_id = p.id
WHERE 
	gp.groupId = ?
ORDER BY 
	p.id;
		`

	permissionRows, err := m.db.Query(permissionQuery, id)
	if err != nil {
		return nil, err
	}
	defer permissionRows.Close()

	for permissionRows.Next() {
		var permissionId int
		var permissionName string
		var permissionValue bool
		if err := permissionRows.Scan(&permissionId, &permissionName, &permissionValue); err != nil {
			return nil, err
		}
		perm := &models.Permission{
			ID:    permissionId,
			Name:  permissionName,
			Value: permissionValue,
		}
		group.Permissions = append(group.Permissions, perm)
	}

	return &group, nil
}

func (m *MariaDB) CreatePermissionGroup(req *models.PermissionGroupRequest) error {
	err := m.setupConnection()
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec("INSERT INTO PermissionGroups (name, description) VALUES (?, ?)", req.Name, req.Description)
	if err != nil {
		tx.Rollback()
		return err
	}

	groupId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, permission := range req.Permissions {
		_, err := tx.Exec("INSERT IGNORE INTO GroupPermissions (groupId, permission_id, value) VALUES (?, ?, ?)", groupId, permission.ID, permission.State)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, memberId := range req.Members {
		_, err := tx.Exec("INSERT IGNORE INTO GroupMembers (groupId, memberId) VALUES (?, ?)", groupId, memberId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if !m.forTesting {
		return tx.Commit()
	} else {
		return tx.Rollback()
	}
}

func (m *MariaDB) UpdatePermissionGroup(id int, req *models.PermissionGroupRequest) error {
	err := m.setupConnection()
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	if len(req.Name) > 0 {
		_, err = tx.Exec("UPDATE PermissionGroups SET name = ?, description = ? WHERE id = ?", req.Name, req.Description, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if req.Permissions != nil {
		for _, permission := range req.Permissions {
			_, err := tx.Exec("INSERT IGNORE INTO GroupPermissions (groupId, permission_id, value) VALUES (?, ?, ?)", id, permission.ID, permission.State)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if req.MembersRemove != nil {
		for _, memberId := range req.MembersRemove {
			_, err := tx.Exec("DELETE FROM GroupMembers WHERE groupId = ? AND memberId = ?", id, memberId)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if req.Members != nil {
		for _, memberId := range req.Members {
			_, err := tx.Exec("INSERT IGNORE INTO GroupMembers (groupId, memberId) VALUES (?, ?)", id, memberId)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if !m.forTesting {
		return tx.Commit()
	} else {
		return tx.Rollback()
	}
}

func (m *MariaDB) GetGroupMembers(groupId int) ([]int, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Query("SELECT member_id FROM GroupMembers WHERE group_id = ?", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []int
	for rows.Next() {
		var memberId int
		if err := rows.Scan(&memberId); err != nil {
			return nil, err
		}
		members = append(members, memberId)
	}
	return members, nil
}

func (m *MariaDB) AddMemberToGroup(groupId int, memberId int) error {
	err := m.setupConnection()
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT IGNORE INTO GroupMembers (group_id, member_id) VALUES (?, ?)", groupId, memberId)

	if !m.forTesting {
		err = tx.Commit()
	}
	return err
}

func (m *MariaDB) RemoveMemberFromGroup(groupId int, memberId int) error {
	err := m.setupConnection()
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM GroupMembers WHERE group_id = ? AND member_id = ?", groupId, memberId)
	if !m.forTesting {
		err = tx.Commit()
	}
	return err
}

func (m *MariaDB) GetPermissionForAccountId(accountId int) ([]*models.Permission, error) {
	err := m.setupConnection()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT p.id, p.name, gp.value
		FROM GroupMembers gm
		JOIN GroupPermissions gp ON gm.groupId = gp.groupId
		JOIN Permissions p ON gp.permission_id = p.id
		WHERE gm.memberId = ?`

	rows, err := m.db.Query(query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		var p models.Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.Value); err != nil {
			return nil, err
		}
		permissions = append(permissions, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
