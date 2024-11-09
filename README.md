[![Build Docker](https://github.com/rainclouded/Concierge/actions/workflows/docker-image.yml/badge.svg)](https://github.com/rainclouded/Concierge/actions/workflows/docker-image.yml)
[![Publish Docker image](https://github.com/rainclouded/Concierge/actions/workflows/push-docker.yaml/badge.svg)](https://github.com/rainclouded/Concierge/actions/workflows/push-docker.yaml)
# Team 6: Concierge

![Concierge Logo](/docs/images/logo.png)


- [@Mondane-Monkeys](https://github.com/Mondane-Monkeys)
- [@LeeroyDilim](https://github.com/LeeroyDilim)
- [@nateng98](https://github.com/nateng98)
- [@mykolabesarab](https://github.com/mykolabesarab)
- [@rainclouded](https://github.com/rainclouded)

## Vision Statement:

Concierge is an easy to use hotel administration system which allows guests to communicate with hotel operators from anywhere and provide comprehensive administration tools for hotel employees.

### More detailed summary: 

To enhance the hotel stay experience, providing a platform that allows guests to request and inquire services and amenities and for employees to manage such services and hotel operations, ensuring convenience and consolidation for all members.

 

Concierge is a platform that is catered to both hotel guests and employees. We will provide hotel management suite that provides both hotel administration and guests with a cohesive and easy to use interface for requesting and inquiring hotel services. Additionally, Concierge includes features tailored for employees, such as dashboards to create and manage incident reports, and tools for efficiently assigning and tracking tasks to ensure smooth hotel operations.

 

Concierge stands out by modernizing the hotel experience, aligning it with today’s mobile-driven world. While most of our daily needs can be easily managed through smartphones, many hotels still rely on outdated systems like landline phones for service requests and inquiries. Concierge transforms this process, making it simple and seamless for both guests and employees to request and manage services through their mobile devices. By streamlining operations and minimizing frustrations, the platform not only enhances guest satisfaction but also drives increased sales through a more efficient, user-friendly experience.

## Core Features
### Task System 
* Allows for hotel guests to create and manage hotel service requests through a user-friendly dashboard. Guests can track the status of their requests, while staff can view incomplete tasks, claim responsibility, and process them in real time.

### Accounts
* Provides secure access for both staff and guests to the hotel management system. Users log in with their room code and gain access to request services. While staff can create accounts and manage permissions of each accounts, while users have the option to customize their settings.

### Amenities: 
* Offers guests detailed information about available hotel amenities, including operating hours and description, displayed on a dashboard. Staff can create, update, or delete amenity listings to ensure that guests have up-to-date information.

### Incident Reports 
* Enables hotel managers to monitor past and current incidents through a dedicated dashboard that displays details such as severity, status and description. Managers can update reports in real time, ensuring accurate information is displayed as an incident case progresses.

# Team 6
Databases: MongoDB, Postgres, and MySQL (Maria)\
Contained in docker \
Orchestrated with Docker Compose 

![Architecture diagram](/docs/sprint_0/images/block_diagram.jpg)

## Quick Build Instructions:
(See [here](/docker-compose/README.md) for more details)

To get the app up and running, use the following script from the project root directory:
```bash
docker compose -f ./docker-compose/docker-compose.yaml build
```
Then once that has been completed run:
```bash
docker compose -f ./docker-compose/docker-compose.yaml up
```
Then once you are done run:
```bash
docker compose -f ./docker-compose/docker-compose.yaml down
```


## Project Structure
```
Concierge/
├── docker-compose/
|   └──configs/
├── docs/
|    └──images/
|   └──sprint_0/
|   └──sprint_1/
|   └──sprint_2/
└── src/
    ├── accounts/
    ├── amenities/
    ├── api_help/
    ├── guest_webapp
    ├── incident_reports/     
    ├── permissions
    ├── task_system
    ├── staff_webapp
    └── system_tests
```
## Branch Naming:
Before you branch, create an issue for what you are working on.
We'll use the issue # in the branch name and snake_case.

the format is:
```
{branch-type}/{feature-name}/{issue-#}
ex. feature/account_server/123
```

## Supporting Documents

### Sprint 0:
- [Sprint 0 Worksheet](/docs/sprint_0/sprint_0_worksheet.md)
### Sprint 1:
- [Sprint 1 Worksheet](/docs/sprint_1/sprint_1_worksheet.md)
### Sprint 1:
- [Sprint 1 Worksheet](/docs/sprint_2/sprint_2_worksheet.md)

### Project:
- [Getting Started (Docker Compose)](docs/project_docs/getting_started.md)

### Technical details
- [Project Technical Overview](/docs/README.md)
- [Concierge Internal Meetings](docs/project_docs/meeting_agenda.md)
- [API Roadmap](/src/api_help/redoc-static.html) (Best to open in a browser)
- [Microservice Architecture Overview](docs/project_docs/microservice_overview.md)
- [Testing Plan](docs/project_docs/testing_plan.md)

### We have style
- [Python](https://google.github.io/styleguide/pyguide.html)
- [C#](https://google.github.io/styleguide/csharp-style.html)
- [JavaScript](https://google.github.io/styleguide/jsguide.html)
- [TypeScript](https://google.github.io/styleguide/tsguide.html)

### Other important (more specific) information

#### API Servers
- [accounts](/src/accounts/README.md)
- [amenities](src/amenities/README.md)
- [incident_reports](src/incident_reports)
- [permissions](src/permissions)
- [task_system](src/task_system)

#### Web Application Servers
- [guest webapp](/src/guest_webapp/README.md)
- [staff webapp](/src/staff_webapp/README.md)

