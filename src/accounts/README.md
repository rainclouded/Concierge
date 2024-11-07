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
Or run the services (for testing):
```
docker compose -f ./docker-compose.dev.yaml up
```

Then clean up:
```
docker compose -f ./docker-compose/docker-compose(.dev).yaml down
```

Local testing:
cd into the accounts directory and first install requirements.

```
python3 -m pip install requirements.txt
```
Next run the tests
```
python3 -m tests
```

## Required Permissions
- canDeleteGuestsAccounts
- canDeleteStaffAccounts
- canEditStaffAccounts
- canEditGuestAccounts

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

    /accounts/login_attempt - Login a user
     body: {'username' : '<username>', 'password' : '<password>'}
     returns: a message containing if the login was successful else error

   /accounts/delete - Update a user account
     body: {'username' : '<username>'}
     headers: {'X-Api-Key' : <key>}
     returns: a message containing if the delete was successful else error


### Put
    /accounts/update - Update a user account
     body: {'username' : '<username>'}
     headers: {'X-Api-Key' : <key>}
     returns: a message containing if the update was successful else error

## Architecture

-   The server/account module runs the api server which manages http requests.
-   The authentication module handles authentication of user credentials (ensures username, password pairs match existing accounts).
-   The database module maintains interfaces/controllers to the database
-   The user_service module facilitates the creation/deletion of user accounts
-   The validation manager provides data validation for user credentials

Below is an image of the interactions between the modules.

![account_architecture](/src/accounts/images/account_diagram.png)