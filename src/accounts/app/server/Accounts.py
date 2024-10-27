"""
Module for the account server
"""
import argparse
import os
import app.Configs as cfg
from app.core.Services import Services
from app.dto.UserObject import UserObject as User
from flask import Flask, jsonify, request
from flask_cors import CORS

app = Flask(__name__)
CORS(app)
database = None
auth = None
user_service = None
permissions = None

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"


def start_service():
    """Run the service"""
    #pylint: disable=global-statement
    global database, auth, user_service, permissions
    database = Services.get_database()
    auth = Services.get_authentication()
    user_service = Services.get_user_service()
    permissions = Services.get_permissions()
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)


def set_services(new_database=None, new_authentication=None,
                 new_user_service=None, new_permissions=None):
    """Inject various services into the Server
    """
    # pylint: disable=global-statement
    global database, auth, user_service, permissions
    if new_database:
        database = new_database
    if new_authentication:
        auth = new_authentication
    if new_user_service:
        user_service = new_user_service
    if new_permissions:
        permissions = new_permissions


def get_port() -> int:
    """Get the port the server should run on
        Returns: port number
    """
    parser = argparse.ArgumentParser(
        description="The authorization microservice"
    )
    parser.add_argument(
        "-p", "--port", type=int, help="Port number for the server to listen on"
    )
    args = parser.parse_args()

    if args.port and (0 <= args.port <= 65536):
        return args.port

    port = os.getenv(ENVIRONMENT_VAR_NAME_PORT)
    if port and (0 <= port <= 65536):
        try:
            return int(port)
        except ValueError:
            print(
                f"Invalid PORT environment variable: {port}."
                + f" Using default port {DEFAULT_PORT}."
            )

    return DEFAULT_PORT


@app.route("/accounts", methods=['GET'])
def index():
    """
    Route to the index page
    """
    response = {
        "message": "You have contacted the accounts",
        "status": "success"
    }
    return jsonify(response)


@app.route("/accounts", methods=["POST"])
def create():
    """
    Route to the account_creation
    """
    print(user_service)
    try:
        response = {
            "message": "Could not create user",
            "status": "error"
        }

        data = request.get_json()
        new_user = User(**{
            'username': data["username"],
            'type': data['type']
        })
        print(1)
        created_user = None
        new_password = None
        print(data)
        if data['type'] == cfg.GUEST_TYPE:
            created_user, new_password = user_service.create_new_guest(new_user)
        else:
            new_password = data["password"]

            created_user = user_service.create_new_staff(new_user, new_password)
        print(44)
        print(created_user)
        print(2)
        if created_user:
            print(3)
            return jsonify({
                "message": f"User created successfully. password: {new_password}",
                "status": "success",

            })
        return jsonify(response), 401
    except Exception as e:
        return jsonify(response), 401


@app.route("/accounts/login_attempt", methods=["POST"])
def login():
    """
    Route to login
    """
    response = {
        "message": "Login Fail - Invalid Credentials",
        "status": "error",
    }
    data = request.get_json()
    if auth.authenticate_user_login(data["username"], data["password"]):
        response["message"] = f"Welcome, {data['username']}!"
        response["status"] = "ok"
        return response

    return response, 401


@app.route("/accounts/delete", methods=["POST"])
def delete():
    """
    Route to delete a user
    """
    response = {
        "message": "Deletion could not be completed.",
        "status": "error",
    }
    data = request.get_json()
    token = request.headers.get('X-Api-Key')

    user_to_delete = data['username']
    user_type = user_service.get_user_type(user_to_delete)

    if (
        user_type == cfg.GUEST_TYPE
        and permissions.can_delete_guest(
            token
        )
    ) or (
        user_type == cfg.STAFF_TYPE
        and permissions.can_delete_staff(
            token
        )
    ):
        if user_service.delete_user(data['username']):
            response["message"] = f"{data['username']} Successfully deleted!"
            response["status"] = "ok"
        else:
            return jsonify(response), 401
    else:
        return jsonify({
            "message": "Action not permitted",
            "status": "forbidden"
        }), 403
    
    return response


@app.route("/accounts/update", methods=["PUT"])
def update():
    """
    Route to update a guest user account
    """
    response = {
        "message": "Update could not be completed.",
        "status": "error",
    }

    data = request.get_json()
    token = request.headers.get('X-Api-Key')
    user_to_change = data['username']

    if (
        user_service.get_user_type(user_to_change) == cfg.GUEST_TYPE
        and permissions.can_update_guest(token)
    ):
        _, new_password = user_service.update_user(
            User(**{
                'username': data["username"],
                'type': cfg.GUEST_TYPE
                }
            )
        )
        if new_password:
            response["message"] = (
                f"{data['username']} Successfully updated!" +
                f" password: {new_password}"
            )
            response["status"] = "ok"
        return jsonify(response), 200

    return jsonify(response), 401
