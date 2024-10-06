# Microservice Overview

[API Documentation (Open in a browser)](/src/api_help/redoc-static.html)

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

