# Concierge
![Concierge Logo](/docs/images/logo.png)

# Team 6
Databases: MongoDB and MySQL\
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
├── Docs/
|   └──sprint_0/
|   └──sprint_1/
└── src/
    ├── accounts/
    ├── api-help/
    ├── guest_webapp
    ├── permissions
    ├── sessions
    ├── staff_webapp
    └── webapp
```
## Branch Naming:
Before you branch, create an issue for what you are working on.
We'll use the issue # in the branch name and snake_case.

the format is:
```
{branch-type}/{feature-name}/{issue-#}
ex. feature/account_server/9354093
```

## Supporting Documents

### Sprint 0:
- [Sprint 0 Worksheet](/docs/sprint_0/sprint_0_worksheet.md)
### Sprint 1:
- [Sprint 1 Worksheet](/docs/sprint_1/sprint_1_worksheet.md)
(In progress)

### Project:
- [Getting Started (Docker Compose)](docs/project_docs/getting_started.md)

### Technical details
- [Project Technical Overview](/docs/README.md)
- [Concierge Internal Meetings](docs/project_docs/meeting_agenda.md)
- [API Roadmap](/src/api-help/redoc-static.html) (Best to open in a browser)
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

#### Web Application Servers
- [guest webapp](/src/guest_webapp/README/md)
- [staff webapp](/src/staff_webapp/README.md)
