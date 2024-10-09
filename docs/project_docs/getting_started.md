# Getting Started

Welcome to Concierge! This guide will help you install Docker Desktop and set up the project on your machine. Follow the steps below to get everything up and running.

---

## Requirements

Before setting up the project, make sure you have the following installed:

1. [Docker Desktop](https://www.docker.com/products/docker-desktop)

---

## 1. Installing Docker Desktop

Docker Desktop is required to run our project using containers. Follow the instructions based on your operating system

### For Windows

1. Visit the [Docker Desktop download page](https://www.docker.com/products/docker-desktop/).
2. Click **Get Docker** and download the executable file for Windows.
3. Run the installer and follow the instructions:
   - Agree to the license terms.
   - Ensure the option to enable WSL 2 is selected for Windows 10/11.
4. After installation, launch Docker Desktop from the Start menu.
5. Then launch the Docker Engine!

## 2. Running our Project

(See [here](/docker-compose/README.md) on more details to add a microservice to our project)

To get the app up and running follow these instructions below:

1. Open terminal and set directory to `Concierge/docker-compose`

2. Run this command to build our project
    
    `docker compose -f ./docker-compose/docker-compose.yaml build`

3. Then once that has been completed, run this command to start project: 

    `docker compose -f ./docker-compose/docker-compose.yaml up`

4. Open browser and head to:
    #### Guest UI:
    [localhost:8089](localhost:8089)
    
    #### Staff UI
    [localhost:8082](localhost:8082)

5. Once finished, run this command to halt project:

    `docker compose -f ./docker-compose/docker-compose.yaml down`
