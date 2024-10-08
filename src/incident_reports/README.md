# Incident Reports Microservice
Microservice to access and manage incident reports for hotels.

## How to Run Server
### Locally
1. Go to directory: `./src/incident_reports`
2. Run command: `python -m incident_reports_server.controllers.incident_reports_controller`

### Via Docker Compose
1. Go to directory: `./docker-compose`
2. Run command: `docker-compose -f docker-compose.yaml build`
3. Run command: `docker-compose -f docker-compose.yaml up`

## How to Run Unit Tests
1. Go to directory: `./src/incident_reports/`
2. Run command: `python -m unittest discover -s .\incident_reports_tests\ -p "*.py"`

## Model

### id
* Unique value that identifies the incident report

### title
* Title of the incident report

### description
* Detailed description of the incident

### created_at
* Timestamp when the incident report was created

### updated_at
* Timestamp when the incident report was last updated

### filing_person_id
* Account ID of the person filing the report

### reviewer_id
* Account ID of the person reviewing the report

## Endpoints
### GET /incident_reports
Retrieve all incident reports.

#### Query Parameters
- **severities**: (optional) A comma-separated list of severity values to filter the reports. 
  - **Valid Values**: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`
- **statuses**: (optional) A comma-separated list of status values to filter the reports. 
  - **Valid Values**: `OPEN`, `CLOSED`, `RESOLVED`, `IN_PROGRESS`
- **beforeDate**: (optional) A date string to filter reports created before this date.
- **afterDate**: (optional) A date string to filter reports created after this date.

### GET /incident_reports/<id>
Retrieve a specific incident report by ID.

### POST /incident_reports
Create a new incident report.

### PUT /incident_reports/<id>
Update an existing incident report with changed information.

### DELETE /incident_reports/<id>
Delete a specified incident report from the database.

## Environment Variables
* `forProduction_DB`
  * Used to determine whether a stub or real implementation of a data layer is used.
* `forProduction_Auth`
  * Used to determine whether a fake or real implementation of a permission validator is used.