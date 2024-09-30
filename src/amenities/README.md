# Amenities Microservice
Microservice to access and manage hotel amenities.

## How to Run Server
### Locally
1. Go to directory: ./src/amenities/amenities_server
2. run command: dotnet run

### Via Docker Compose
1. Go to directory: ./docker-compose
2. run command: docker-compose -f docker-compose.yaml up

## How to Run Unit Tests
1. Go to directory: ./src/amenities/amenities_test
2. run command: dotnet test

## Endpoints
### GET /amenities
get all amenities

### GET /amenities/id
get specific amenity by id

### POST /amenities
create amenity

### PUT /amenities
update amenity with changed information

### DELETE /amenities/id
delete specified amenity from database

## Environment Variables
 * forProduction
	* used to determine whether a stub or real implementation of a data layer is returned
