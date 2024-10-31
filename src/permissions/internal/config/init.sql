
CREATE DATABASE IF NOT EXISTS permissions_db;

CREATE TABLE IF NOT EXISTS permissions_db.Permissions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS permissions_db.PermissionGroups (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS permissions_db.GroupMembers (
    groupId INT NOT NULL,
    memberId INT NOT NULL,
    PRIMARY KEY (groupId, memberId),
    FOREIGN KEY (groupId) REFERENCES permissions_db.PermissionGroups(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS permissions_db.GroupPermissions (
    groupId INT NOT NULL,
    permission_id INT NOT NULL,
    value BOOLEAN NOT NULL,
    PRIMARY KEY (groupId, permission_id),
    FOREIGN KEY (groupId) REFERENCES permissions_db.PermissionGroups(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions_db.Permissions(id) ON DELETE CASCADE
);

INSERT INTO permissions_db.Permissions (name) VALUES
('canViewPermissionGroups'),
('canEditPermissionGroups'),
('canViewPermissions'),
('canEditPermissions');

INSERT INTO permissions_db.PermissionGroups (name, description) VALUES
('admin', 'Has all permissions'),
('editor', 'Can edit and view'),
('viewer', 'Can only view');

INSERT INTO permissions_db.GroupPermissions (groupId, permission_id, value) VALUES
(1, 1, true),
(1, 2, true),
(1, 3, true),
(1, 4, true),
(2, 1, true),
(2, 3, true);

INSERT INTO permissions_db.GroupMembers (groupId, memberId) VALUES
(1, 0),
(1, 1),
(1, 2),
(2, 0),
(2, 1),
(3, -1),
(3, 4),
(3, 5);