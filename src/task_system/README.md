# Task System Microservice
Microservice to access and manage hotel tasks.

## How to Run Server

### Locally
1. Go to directory: `./src/task_system/task_system_server`
2. run command: `dotnet watch run`

### Via Docker Compose
1. Open Docker
2. Go to directory: `./docker-compose`
3. run command: `docker-compose -f docker-compose.dev.yaml up --build`

## How to Test

### Run Unit Test
1. Go to directory: `./src/task_system/task_system_test`
2. run command: `dotnet test`

### Run Integration Test
1. Open Docker
2. Run command from anywhere: `docker run -e POSTGRES_PASSWORD=sa -d -p 50021:5432 postgres`
3. Go to directory: `./src/task_system/task_system_integration_test`
4. run command: `dotnet test`

### Run All Tests
1. Open Docker
2. Run command from anywhere: `docker run -e POSTGRES_PASSWORD=sa -d -p 50017:5432 postgres`
3. Go to directory: `./src/task_system`
4. run command: `dotnet test`

## Model

| Field Name   | Description|
|--------------|---------------------------------------|
| **Id**       | Unique value that identifies the task.|
| **TaskType**    | Type of service/task.|
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

## Environment Variables
* `DB_IMPLEMENTATION`
	* used to determine whether a stub or real implementation of a data layer is returned
	* Value: `POSTGRES`, Attempts to connect to a POSTGRESQL database
	* Value: `MOCK`, Uses in-memory storage
* `DB_HOST`
   * Configure Postgress Connection string
* `DB_PORT`
   * Configure Postgress Connection string
* `DB_USERNAME`
   * Configure Postgress Connection string
* `DB_PASSWORD`
   * Configure Postgress Connection string