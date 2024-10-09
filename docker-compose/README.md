# Docker Compose

## Running the services
- Ensure Docker Engine is running, and that Docker Compose is installed on your machine
- run `docker compose -f docker-compose.yaml build`
    - Note, ensure the relative path to docker-compose.yaml is provided in the above command
- Browse to localhost:8089
- Or, Browse to the endpoint you want to test, as defined in the docker-compose.yaml


## Adding a service to the environment
- Create a Dockerfile for your service.
- Create a docker-compose.yaml file in your service.
- Test it locally in straight Docker. Ensure endpoints are accessible.
- In docker-compose.yaml, add an entry to to "services:"
- Give your service a *meaningful* name.
    - This name is how services can find each other within Docker Compose.
- Reference your server's compose file.
    - add an `extends:` tag
    - Inside that tag, include:
      - `file: <path/to/your-file>`
      - `service: <your-service-name>`
- Add your service to the [nginx config file](configs/nginx.conf) to make your service public
