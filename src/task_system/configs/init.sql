CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    description TEXT,
    roomId NUMERIC,
    assigneeId NUMERIC,
    status VARCHAR(100),
    createdAt DATETIME
);

INSERT INTO tasks (title, description, description, roomId, assigneeId, status, createdAt) VALUES
('Maintenance', 'There is a leak in the bathroom sink that needs urgent attention.', 101, 1, 2, 'In Progress', '2024-10-10 10:30:00'),
('Maintenance', 'Some light bulbs are out in the hallway. Please replace them.', 102, 3, 2, 'Pending', '2024-10-12 09:00:00'),
('Maintenance', 'The conference room needs to be cleaned before the meeting.', 201, 4, 5, 'Completed', '2024-10-13 15:00:00'),
('Maintenance', 'Ensure that all fire alarms have functioning batteries.', 103, 2, 3, 'Pending', '2024-10-15 11:00:00'),
('Maintenance', 'Organize the storage area to make supplies easily accessible.', 104, 1, 4, 'In Progress', '2024-10-18 14:00:00');