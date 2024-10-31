CREATE TABLE IF NOT EXISTS "Tasks" (
    "Id" SERIAL PRIMARY KEY,
    "TaskType" INTEGER,
    "Description" TEXT,
    "RoomId" NUMERIC,
    "RequesterId" NUMERIC,
    "AssigneeId" NUMERIC,
    "Status" INTEGER,
    "CreatedAt" TIMESTAMP
);

INSERT INTO "Tasks" ("TaskType", "Description", "RoomId", "RequesterId", "AssigneeId", "Status", "CreatedAt") VALUES
(1, 'There is a leak in the bathroom sink that needs urgent attention.', 101, 1, 2, 1, '2024-10-10 10:30:00'),
(1, 'Some light bulbs are out in the hallway. Please replace them.', 102, 3, 2, 0, '2024-10-12 09:00:00'),
(1, 'The conference room needs to be cleaned before the meeting.', 201, 4, 5, 2, '2024-10-13 15:00:00'),
(1, 'Ensure that all fire alarms have functioning batteries.', 103, 2, 3, 0, '2024-10-15 11:00:00'),
(1, 'Organize the storage area to make supplies easily accessible.', 104, 1, 4, 1, '2024-10-18 14:00:00');