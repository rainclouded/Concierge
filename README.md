# Concierge


![Concierge Logo](/Docs/images/logo.png)

# Team 6

![rainclouded](https://avatars.githubusercontent.com/u/rainclouded) [rainclouded](https://github.com/rainclouded)\
![LeeroyDilim](https://avatars.githubusercontent.com/u/LeeroyDilim) [LeeroyDilim](https://github.com/LeeroyDilim)\
![mykolabesarab](https://avatars.githubusercontent.com/u/mykolabesarab) [mykolabesarab](https://github.com/mykolabesarab)\
![nateng98](https://avatars.githubusercontent.com/u/nateng98) [nateng98](https://github.com/nateng98)\
![Mondane-Monkeys](https://avatars.githubusercontent.com/u/Mondane-Monkeys) [Mondane-Monkeys](https://github.com/Mondane-Monkeys)

## Table of contents
1. [Vision Statement](#vision-statement)
2. [Build Instructions](#build-instructions)
3. [Endpoints](#endpoints)
4. [Enviornment Variables](#environment-variables)
5. [Project Structure](#project-structure)
6. [Supporting Documents](#supporting-documents)

## Vision Statement:

Concierge is an easy to use hotel administration system which allows guests to communicate with hotel operators from anywhere and provide comprehensive administration tools for hotel employees.

### Complete summary: 

To enhance the hotel stay experience, providing a platform that allows guests to request and inquire services and amenities and for employees to manage such services and hotel operations, ensuring convenience and consolidation for all members.

Concierge is a platform that is catered to both hotel guests and employees. We will provide hotel management suite that provides both hotel administration and guests with a cohesive and easy to use interface for requesting and inquiring hotel services. Additionally, Concierge includes features tailored for employees, such as dashboards to create and manage incident reports, and tools for efficiently assigning and tracking tasks to ensure smooth hotel operations.

Concierge stands out by modernizing the hotel experience, aligning it with today’s mobile-driven world. While most of our daily needs can be easily managed through smartphones, many hotels still rely on outdated systems like landline phones for service requests and inquiries. Concierge transforms this process, making it simple and seamless for both guests and employees to request and manage services through their mobile devices. By streamlining operations and minimizing frustrations, the platform not only enhances guest satisfaction but also drives increased sales through a more efficient, user-friendly experience.


## Build Instructions:

(See [here](/docker-compose/README.md) for more details)

To get the app up and running, cd into Concierge/docker-compose and run:
```bash
docker compose -f docker-compose.yaml build
```


## Endpoints

### Accounts:
#### GET:
```
/accounts - List all user accounts
```
#### POST:
```
/accounts - Create a new account
/accounts/login_attempt - Attempt to log user into account
```
### Permissions:
#### GET:
```
/permissions - List all permissions
/permissions/{permission-id}/sessions/{session-key}  - Get the permissions of the associated key
/permission-groups - Get all the permission groups
/permission-groups/{group-id} - Get the permissions and details of the associated group
/permission-groups/{group-id}/accounts - Get all accounts in the corresponding group
```

#### POST:
```
/permissions - Create a new permission
/permission-groups - Create a new permission group
/permission-groups/{group-id}/accounts - Add account to permission group
```
#### PUT:
```
/permission-groups/{group-id} - Edit a permission group

```

### Sessions:
#### GET:
```
/sessions/{session-key} - Get the details of this session
```
#### POST:
```
/sessions - Login to account
```

### Incident Reports:
#### GET:
```
/incident-reports - View incident reports
```

#### POST:
```
/incident-reports - Create an incident report
```

### Amenities:
#### GET:
```
/amenities - List all of the amenities
```
#### POST:
```
/amenities - Create a new amenity
```
#### PUT:
```
/amenities/{amenity-id} - Update an existing amenity
```

### Tasks:
#### GET:
```
/tasks - List all tasks
/tasks/{task-id} - Get info about specified task
/task-templates - List existing task templates
/task-template/{task-id} - Get info about specified task template
/guest-templates - List all of the templates for guests
```
#### POST:
```
/tasks - Create a new task
/task-templates - Create a new task template
/guest-templates - Create a new guest template
```

#### PUT:
```
/tasks/{task-id} - Update specified task
/task-template/{task-id} - Update specified task template
```

### Reservations:
#### GET:
```
/reservations - Get all current reservations
```
#### POST:
```
/reservations - Create a new reservation
```

#### PUT:
```
/reservations/{reservation-id} - Update the specified reservation
```
## Environment variables

There are a few environment variables that can be set:
```bash
ACCOUNTS_PORT #The port to bind the Accounts service to
SESSIONS_PORT #The port to bind the Sessions serivce to
PERMISSIONS_PORT #The port to bind the Permissions service to
WEBAPP_PORT #The port to bind the WEBAPP service to
```

## Project Structure
```
Concierge/
├── docker-compose/
|   └──configs/
├── Docs/
|   └──Sprint0/
└── src/
    ├── accounts/
    ├── api-help/
    ├── guest_webapp     
    ├── permissions
    ├── sessions
    ├── staff_webapp
    └── webapp
```


## Supporting Documents

[Sprint 0 Worksheet](/Docs/Sprint0/sprint_0_worksheet.md)

[Docker build instructions](/docker-compose/README.md)


