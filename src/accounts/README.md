# Accounts

This is the accounts service. It handles all of the authentication and validation of user accounts either staff or guest.

## Running the service:

Local:

 cd into the accounts directory and first install requirements.

```
python3 -m pip install requirements.txt
```
Next start the server (ctrl-c to stop)
```
python3 -m accounts
```

Using docker:

 cd into the Concierge/docker-compose directory and first build the project.

```
docker compose -f ./docker-compose.yaml build
```
Then run the services (ctrl-c to stop):
```
docker compose -f ./docker-compose.yaml up
```
Then run the services:
```
docker compose -f ./docker-compose.yaml up
```

Then clean up:
```
docker compose -f ./docker-compose down
```

Local testing:
cd into the accounts directory and first install requirements.

```
python3 -m pip install requirements.txt
```
Next run the tests (ctrl-c to stop)
```
python3 -m tests
```

## Endpoints
Handles all /accounts endpoints

### Get
    /accounts - index page of index
    body: none
    returns: a message if the page was accessed successfully

### Post
     /accounts - Create a new user
     body: {'username' : '<username>', 'password' : '<password>', 'type':'<guest/staff>' }
     password is not required for 'guest' type
     returns: a message containing the password if successful else error

    /accounts/login_attempt - Create a new user
     body: {'username' : '<username>', 'password' : '<password>'}
     returns: a message containing if the login was successful else error
