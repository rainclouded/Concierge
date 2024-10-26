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
database = Services.get_database()
auth = Services.get_authentication()
user_service = Services.get_user_service()
permissions = Services.get_permissions()

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"


def start_service():
    """Run the service"""
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)


def get_port()->int:
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
    response = {
        "message" : "Could not create user",
        "status" : "error"
    }

    data = request.get_json()
    new_user  = User(**{
        'username' : data["username"],
        'type' : data['type']
    })

    created_user = None
    new_password = None

    if data['type'] == cfg.GUEST_TYPE:
        created_user, new_password = user_service.create_new_guest(new_user)
    else:
        new_password = data["password"]

        created_user = user_service.create_new_staff(new_user, new_password)

    if created_user:
        return jsonify({
            "message" : f"User created successfully. password: {new_password}",
            "status" : "success",

        })
    return jsonify(response)


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
    permisson_key = request.headers.get('X-Api-Key')
    user_to_delete = data['change_request']

    if permissions.can_delete_user(permisson_key, permissions.get_user_type(user_to_delete)):
        if user_service.delete_user(data['change_request']):
            response["message"] = f"{data['change_request']} Successfully deleted!"
            response["status"] = "ok"
    else:
        response = {
            "message": "Action not permitted",
            "status": "forbidden"
        }
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
    permisson_key = request.headers.get('X-Api-Key')
    user_to_change = data['change_request']

    try:
        match user_service.get_user_type(user_to_change):
            case cfg.GUEST_TYPE:
                if permissions.can_update_user(permisson_key, cfg.GUEST_TYPE):
                    _, new_password = user_service.update_user(data['change_request'])
                    if new_password:
                        response["message"] =\
                        (
                            f"{data['change_request']} Successfully updated.!" +
                            f"{new_password}"
                        )
                        response["status"] = "ok"
            case cfg.STAFF_TYPE:
                response = {
                    "message": "Cannot delete user",
                    "status": "error",
                }

    except LookupError:
        response = {
            "message": "User not found",
            "status": "error",
        }


    return response
