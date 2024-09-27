# Docker Compose

## Running the services
- Ensure Docker Engine is running, and that Docker Compose is installed on your machine
- run `docker compose -f docker-compose.yaml build`
    - Note, ensure the relative path to docker-compose.yaml is provided in the above command
- Browse to localhost:8089
- Or, Browse to the endpoint you want to test, as defined in the docker-compose.yaml


## Adding a service to the environment
- Create a Dockerfile for your service. 
- Test it locally in straight Docker. Ensure endpoints are accessible
- In docker-compose.yaml, add an entry to to "services:"
- Give your service a *meaningful* name.
    - This name is how services can find each other within Docker Compose
- Give your service proprties
    - "build:" should point to your Dockerfile
    - "ports:" opens ports to your service. 
        - Note the format, [external port]:[internal port]
        - Browsing from your host to [external port] will send a request to your service at [insternal port]
    - Other configurations include 
        - environment variables, 
        - docker image (instead of build)
        - "depends_on" ensures another service is built and run before your service
        - Many other config tags exist, feel free to look into them
