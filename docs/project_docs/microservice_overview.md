# Microservice Overview

[API Documentation (Open in a browser)](/src/api_help/redoc-static.html)

## Architecture Diagram
![](/docs/project_docs/images/sprint1_arch_diag.drawio.svg)

### Description
The database will be MongoDB and MySQL.\
The backend will be microservices orchestrated by Kubernetes with Python (and maybe C#) webservers hosted on Azure.\
The first front end will be for desktop and will utilize Angular 18 and TypeScript.\
The second from end will be for mobile and desktop and will utilize React and JavaScript.

#### Separation of layers

We will have a 3 layer architecture all encapsulated within microservices.

All layers will be hosted on azure separated into their own microservice component.

The first layer is the UI layer which consists of our two front ends. The second layer is our business/logic layer. This layer handles the logic/processing of the communications between the differet layers. Lastly is out data layer which consists of all of our databases storing all user data.


## Nginx
We use this to route our user requests to the appropriate microservice for processing. 

## Accounts
Handles all `/accounts` endpoints.

Dedicated to managing both guest and staff account information. Also used in authenticating app permissions. 

[Read more.](/src/accounts/README.md)

## Amenities
Handles all `/amenities` endpoints.

Dedicated to managing hotel amenity information. 

[Read more.](/src/amenities/README.md)

## Incident Reports
Handles all `/incident_report` endpoints.

Dedicated to managing incident reports made by hotel staff. 

[Read more.](/src/accounts/README.md)

## Guest Web App
The guest front end of Concierge. Responsible in displaying an appealing interface for our guests to access hotel information and request for services. 

[Read more.](/src/guest_webapp/README.md)

## Staff Webapp
The staff front end of Concierge. Responsible in displaying an functional interface for our staff to update hotel information and process guest requests. 

[Read more.](/src/staff_webapp/README.md)

