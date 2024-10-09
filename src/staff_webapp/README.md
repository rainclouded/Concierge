# Admin Dashboard - Staff Webapp

## Prerequisites

Ensure you have the following installed on your machine:

- Node.js
- npm (Node Package Manager)
- Docker (for containerization)

## Run Locally

1. Install Angular CLI

If you haven’t already installed Angular CLI globally, run the following command:

```
npm install -g @angular/cli
```

2. Install dependencies

```
npm install
```

3. Run Angular App

Start the development server using the Angular CLI:

```
ng serve
```

Once the server is running, you can access the application in your web browser at:

[http://localhost:4200/](http://localhost:4200/)


## Docker Setup

Run Docker Compose

Execute the following command to start the application:

```
docker compose up
```

This will build the necessary containers and start the application.

Once it is running, you can access the application in your web browser at:

[http://localhost:8082/](http://localhost:8082/)


## Tests

### Unit/Spec test
To run the angular spec tests cd into /staff_webapp and run:

```
ng test
```

### Integration tests - GUI

The integration tests rely on a fresh environment using stub API data

Run the entire app from the root directory using:

```
docker compose -f ./docker-compose/docker-compose.yaml up --build
```
then cd into the staff_webapp directory. To open the Cypress UI and run the integration tests run:
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
then cd into the staff_webapp directory. To open run the integration tests use:
```
npx cypress run
```
