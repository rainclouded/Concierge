# Amenities Microservice
Microservice to access and manage hotel amenities.

## How to Run Server

### Locally
1. Go to directory: `./src/amenities`
2. run command: `dotnet run`

### Via Docker Compose
1. Go to directory: `./docker-compose`
2. run command: `docker-compose -f docker-compose.yaml build`
3. run command: `docker-compose -f docker-compose.yaml up`

## How to Run Unit Tests
1. Go to directory: `./src/amenities/amenities_test`
2. run command: `dotnet test`

## How to Run Integration Tests
1. Go to directory: `./src/amenities/`
2. Run `docker compose -f docker-compose.test.yaml up -d`
3. Go to `src/amenities/amenities_db_integration_test`
4. Run command: `dotnet test`

## Permissions
 - canViewAmenities
 - canEditAmenities
 - canDeleteAmenities

## Model

### Id
* Unique value that identifies the amenity

### Name
* Name of amenity

### Description
* Description of amenity

### StartTime
* Opening time of amenity (constant throughout the week) 

### EndTime
* Closing time of amenity (constant throughout the week)

## Endpoints
### GET /amenities
get all amenities

### GET /amenities/id
get specific amenity by id

### POST /amenities
create amenity

### PUT /amenities/id
update amenity with changed information

### DELETE /amenities/id
delete specified amenity from database

## Environment Variables
* `PERMISSIONS_ENDPOINT`
   * Used to identify the permissions endpoint. Eg, "http://permissions:8080"
* `SESSIONS_ENDPOINT`
   * Used to identify the Sessions endpoint. Eg, "http://sessions:8080"
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


## Unexpected behaviour
- If the server cannot connect to the postgres database when first initialized, it assumes there was a configuration error and uses the default mock database.
  - If the server loses connection to the postgres server after at least 1 success connection, it will return 500 errors until the database is restored
