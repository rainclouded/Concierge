# Hotel Services App - Guest Webapp

## Prerequisites

Ensure you have the following installed on your machine:

- Node.js
- npm (Node Package Manager)
- Docker (for containerization)

## Run Locally

1. Install dependencies

```
npm install
```

2. Run React App

Run the development server using the following command:

```
npm run dev
```

Once the server is running, you can access the application in your web browser at:

[http://localhost:5173/](http://localhost:5173/)


## Docker Setup

Run Docker Compose

Execute the following command to start the application:

```
docker compose up
```

This will build the necessary containers and start the application.

Once it is running, you can access the application in your web browser at:

[http://localhost:8081/](http://localhost:8081/)


## Tests

THE INCIDENT REPORTS INTEGRATION TESTS HAVE SIDE EFFECTS!
 - Creates a new Incident Report. Either reset that server or remove the new report manually before running any other service's incident_report_integration tests.

### Integration tests - GUI

The integration tests rely on a fresh environment using stub API data

Run the entire app from the root directory using:

```
docker compose -f ./docker-compose/docker-compose.yaml up --build
```
then cd into the guest_webapp directory. To open the Cypress UI and run the integration tests run:
```
npx cypress open
```

If you get a "Cannot find package 'cypress'" error, run 
```
npm install cypress --save-dev
```

### Integration tests - Headless

The integration tests rely on a fresh environment using stub API data

Run the entire app from the root directory using

```
docker compose -f ./docker-compose/docker-compose.yaml up --build
```
then cd into the guest_webapp directory. To open run the integration tests use:
```
npx cypress run
```
