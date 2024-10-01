# Concierge


![Concierge Logo](/Docs/images/logo.png)

# Team 6

[rainclouded](https://github.com/rainclouded)\
[LeeroyDilim](https://github.com/LeeroyDilim)\
[mykolabesarab](https://github.com/mykolabesarab)\
[nateng98](https://github.com/nateng98)\
[Mondane-Monkeys](https://github.com/Mondane-Monkeys)

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

### Features
[Here](https://github.com/rainclouded/Concierge/issues?q=is%3Aopen+is%3Aissue+label%3AFeature) are our current features.

## Tech Stack:

Front end 1 (Desktop): Angular+TypeScript\
Front end 2 (Dexktop/Mobile): React+JavaScript\
Backend: Microservices written in Python/C# ASP.NET\
Databases: MongoDB and MySQL\
Contained in docker

## Build Instructions:

(See [here](/docker-compose/README.md) for more details)

To get the app up and running, cd into Concierge/docker-compose and run:
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

### Sprint 0:
[Sprint 0 Worksheet](/Docs/Sprint0/sprint_0_worksheet.md)
### Sprint 1:
(In progress)

### Project:

[Docker build instructions](/docker-compose/README.md)

### Technical details
[Concierge Technical Info](/Docs/project-technical-details.md)

### We have style

[Python](https://google.github.io/styleguide/pyguide.html)\
[C#](https://google.github.io/styleguide/csharp-style.html)\
[JavaScript](https://google.github.io/styleguide/jsguide.html)\
[TypeScript](https://google.github.io/styleguide/tsguide.html)