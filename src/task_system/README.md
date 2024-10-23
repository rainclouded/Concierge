# Task System Microservice
Microservice to access and manage hotel tasks.

## How to Run Server

### Locally
1. Go to directory: `./src/task_system/task_system_server`
2. run command: `dotnet watch run`

### Via Docker Compose
1. Go to directory: `./docker-compose`
2. run command: `docker-compose -f docker-compose.dev.yaml up --build`

## Model

| Field Name   | Description|
|--------------|---------------------------------------|
| **Id**       | Unique value that identifies the task.|
| **Title**    | Type of service/task.|
| **Description** | Detailed description of the task, outlining what needs to be done.|
| **RoomId**   | Unique identifier of the room associated with the task.|
| **AssigneeId** | Unique identifier of the user assigned to the task.|
| **Status**   | Status of the task (e.g., Pending, In Progress, Completed).|
| **CreatedAt**| Day and time indicating when the task was created.       

## Endpoints
### GET /tasks
get all tasks

### GET /tasks/id
get specific amenity by id

### POST /tasks
create amenity

### PUT /tasks/id
update amenity with changed information

### DELETE /tasks/id
delete specified amenity from database

