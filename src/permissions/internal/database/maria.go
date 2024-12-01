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

// NewMariaDB creates and returns a new MariaDB instance, initializing the database connection and testing flag.
// Args:
//     dataSourceName (string): The data source name for the database connection.
//     forTesting (bool): Flag to determine if the database is for testing.
// Returns:
//     *MariaDB: A new instance of MariaDB.
//     error: Any error encountered during initialization (not used here).
func NewMariaDB(dataSourceName string, forTesting bool) (*MariaDB, error) {
	return &MariaDB{db: nil, dataSource: dataSourceName, forTesting: forTesting}, nil
}

// Close closes the database connection if it is open.
// Args:
//     None
// Returns:
//     error: Any error encountered during closing the connection, or nil if successful.
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

// setupConnection sets up the database connection if not already established.
// Args:
//     None
// Returns:
//     error: Any error encountered while establishing the connection.
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

// GetPermissions retrieves all permissions from the database.
// Args:
//     None
// Returns:
//     []*models.Permission: A slice of permission models.
//     error: Any error encountered while fetching the permissions.
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

// GetPermissionById retrieves a permission by its ID from the database.
// Args:
//     id (int): The ID of the permission to retrieve.
// Returns:
//     *models.Permission: The retrieved permission.
//     error: Any error encountered while fetching the permission.
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
	return &p, nil
}

// CreatePermission creates a new permission in the database.
// Args:
//     name (string): The name of the new permission.
// Returns:
//     *models.Permission: The created permission.
//     error: Any error encountered while creating the permission.
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

// UpdatePermission updates the details of an existing permission in the database.
// Args:
//     p (*models.Permission): The permission model with updated details.
// Returns:
//     error: Any error encountered while updating the permission.
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

// GetPermissionGroups retrieves all permission groups, including their members and permissions.
// Args:
//     None
// Returns:
//     []*models.PermissionGroup: A slice of permission group models.
//     error: Any error encountered while fetching the permission groups.
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

// GetPermissionGroupById retrieves a permission group by its ID, including its members and permissions.
// Args:
//     id (int): The ID of the permission group.
// Returns:
//     *models.PermissionGroup: The retrieved permission group.
//     error: Any error encountered while fetching the permission group.
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
// CreatePermissionGroup creates a new permission group in the database, inserts associated permissions
// and members, and commits the transaction. Rolls back the transaction if an error occurs.
// Args:
//    req: A struct containing the details of the permission group, including its name, description,
//         permissions, and members to be added to the group.
// Returns:
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) CreatePermissionGroup(req *models.PermissionGroupRequest) error {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return err  // Return error if connection setup fails
	}

	tx, err := m.db.Begin()  // Begin a new database transaction
	if err != nil {
		return err  // Return error if transaction initialization fails
	}

	// Insert the new permission group into the PermissionGroups table
	result, err := tx.Exec("INSERT INTO PermissionGroups (name, description) VALUES (?, ?)", req.Name, req.Description)
	if err != nil {
		tx.Rollback()  // Rollback transaction if insert fails
		return err
	}

	groupId, err := result.LastInsertId()  // Get the last inserted ID for the new group
	if err != nil {
		tx.Rollback()  // Rollback transaction if retrieving the last insert ID fails
		return err
	}

	// Insert the permissions associated with the new group into GroupPermissions
	for _, permission := range req.Permissions {
		_, err := tx.Exec("INSERT IGNORE INTO GroupPermissions (groupId, permission_id, value) VALUES (?, ?, ?)", groupId, permission.ID, permission.State)
		if err != nil {
			tx.Rollback()  // Rollback transaction if permission insert fails
			return err
		}
	}

	// Insert the members associated with the new group into GroupMembers
	for _, memberId := range req.Members {
		_, err := tx.Exec("INSERT IGNORE INTO GroupMembers (groupId, memberId) VALUES (?, ?)", groupId, memberId)
		if err != nil {
			tx.Rollback()  // Rollback transaction if member insert fails
			return err
		}
	}

	if !m.forTesting {  // If not in testing mode, commit the transaction
		return tx.Commit()
	} else {
		return tx.Rollback()  // Rollback transaction in testing mode
	}
}

// UpdatePermissionGroup updates an existing permission group's name, description, permissions, and members.
// It handles adding and removing permissions and members, and commits the transaction. Rolls back on error.
// Args:
//    id: The ID of the permission group to be updated.
//    req: A struct containing the updated details of the permission group, including its name, description,
//         permissions, members to be added, and members to be removed.
// Returns:
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) UpdatePermissionGroup(id int, req *models.PermissionGroupRequest) error {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return err  // Return error if connection setup fails
	}

	tx, err := m.db.Begin()  // Begin a new database transaction
	if err != nil {
		return err  // Return error if transaction initialization fails
	}

	// Update the permission group's name and description
	if len(req.Name) > 0 {
		_, err = tx.Exec("UPDATE PermissionGroups SET name = ?, description = ? WHERE id = ?", req.Name, req.Description, id)
		if err != nil {
			tx.Rollback()  // Rollback transaction if group update fails
			return err
		}
	}

	// Update or insert permissions for the group
	if req.Permissions != nil {
		for _, permission := range req.Permissions {
			_, err := tx.Exec(`INSERT INTO GroupPermissions (groupId, permission_id, value) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE value = ?`, id, permission.ID, permission.State, permission.State)
			if err != nil {
				tx.Rollback()  // Rollback transaction if permission insert fails
				return err
			}
		}
	}

	// Remove members from the group
	if req.MembersRemove != nil {
		for _, memberId := range req.MembersRemove {
			_, err := tx.Exec("DELETE FROM GroupMembers WHERE groupId = ? AND memberId = ?", id, memberId)
			if err != nil {
				tx.Rollback()  // Rollback transaction if member removal fails
				return err
			}
		}
	}

	// Insert new members into the group
	if req.Members != nil {
		for _, memberId := range req.Members {
			_, err := tx.Exec("INSERT IGNORE INTO GroupMembers (groupId, memberId) VALUES (?, ?)", id, memberId)
			if err != nil {
				tx.Rollback()  // Rollback transaction if member insert fails
				return err
			}
		}
	}

	if !m.forTesting {  // If not in testing mode, commit the transaction
		return tx.Commit()
	} else {
		return tx.Rollback()  // Rollback transaction in testing mode
	}
}

// GetGroupMembers retrieves the list of members in a given permission group.
// Args:
//    groupId: The ID of the permission group whose members are to be retrieved.
// Returns:
//    members: A slice of integers representing member IDs in the group.
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) GetGroupMembers(groupId int) ([]int, error) {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return nil, err  // Return error if connection setup fails
	}

	rows, err := m.db.Query("SELECT member_id FROM GroupMembers WHERE group_id = ?", groupId)  // Query for members in the group
	if err != nil {
		return nil, err  // Return error if query fails
	}
	defer rows.Close()

	var members []int
	for rows.Next() {  // Loop through each member row
		var memberId int
		if err := rows.Scan(&memberId); err != nil {
			return nil, err  // Return error if scanning fails
		}
		members = append(members, memberId)  // Add member ID to the list
	}
	return members, nil  // Return the list of members
}

// AddMemberToGroup adds a member to a given permission group.
// Args:
//    groupId: The ID of the permission group to which the member is to be added.
//    memberId: The ID of the member to be added to the group.
// Returns:
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) AddMemberToGroup(groupId int, memberId int) error {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return err  // Return error if connection setup fails
	}

	tx, err := m.db.Begin()  // Begin a new database transaction
	if err != nil {
		return err  // Return error if transaction initialization fails
	}

	_, err = tx.Exec("INSERT IGNORE INTO GroupMembers (group_id, member_id) VALUES (?, ?)", groupId, memberId)  // Add member to the group
	if !m.forTesting {  // If not in testing mode, commit the transaction
		err = tx.Commit()
	}
	return err  // Return any error that occurred
}

// RemoveMemberFromGroup removes a member from a given permission group.
// Args:
//    groupId: The ID of the permission group from which the member is to be removed.
//    memberId: The ID of the member to be removed from the group.
// Returns:
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) RemoveMemberFromGroup(groupId int, memberId int) error {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return err  // Return error if connection setup fails
	}

	tx, err := m.db.Begin()  // Begin a new database transaction
	if err != nil {
		return err  // Return error if transaction initialization fails
	}

	_, err = tx.Exec("DELETE FROM GroupMembers WHERE group_id = ? AND member_id = ?", groupId, memberId)  // Remove member from the group
	if !m.forTesting {  // If not in testing mode, commit the transaction
		err = tx.Commit()
	}
	return err  // Return any error that occurred
}

// GetPermissionForAccountId retrieves the permissions associated with a specific account (member).
// Args:
//    accountId: The ID of the account for which permissions are to be retrieved.
// Returns:
//    permissions: A slice of pointers to Permission models representing the permissions granted to the account.
//    error: Returns an error if any database operation fails, otherwise returns nil.
func (m *MariaDB) GetPermissionForAccountId(accountId int) ([]*models.Permission, error) {
	err := m.setupConnection()  // Set up the database connection
	if err != nil {
		return nil, err  // Return error if connection setup fails
	}

	query := `
Select p.id, p.name, gp.value from GroupPermissions as gp
LEFT JOIN GroupMembers as gm on gp.groupId = gm.groupId
LEFT JOIN Permissions as p on gp.permission_id = p.Id
WHERE gm.memberId = ?
GROUP BY p.name, p.id, gp.value
ORDER BY p.id;
		`

	rows, err := m.db.Query(query, accountId)  // Query for permissions of the specified account
	if err != nil {
		return nil, err  // Return error if query fails
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {  // Loop through each permission row
		var p models.Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.Value); err != nil {
			return nil, err  // Return error if scanning fails
		}
		if p.Value {  // Only include active permissions
			permissions = append(permissions, &p)
		}
	}

	if err := rows.Err(); err != nil {  // Check for errors after looping through rows
		return nil, err
	}

	return permissions, nil  // Return the list of permissions
}
