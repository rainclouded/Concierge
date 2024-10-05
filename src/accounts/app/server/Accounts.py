import argparse
import os
import app.Configs as cfg
from flask import Flask, jsonify, request
from app.dto.UserObject import UserObject as User
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService

app = Flask(__name__)
auth = user_service = None

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"


def start_service(database:DatabaseController):
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)
    auth = AuthenticationManager(database)
    user_service = UserService(database)


def get_port():
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


@app.route("/accounts/", methods=['GET'])
def index():
    """
    Route the index page
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
        'password' : data["password"],
        'type' : data['type']
    })

    if (
        user_service.create_new_guest(new_user)
        if new_user.type == cfg.GUEST_TYPE
        else user_service.create_new_staff(new_user)
        ):
        return jsonify({
            "message" : "User created successfully",
            "status" : "success"

        })
    return jsonify(response)

@app.route("/accounts/login_attempt", methods=["POST"])
def login():
    """
    Route to login
    """
    response = {
        "username": None,
        "message": "Login Fail - Invalid Credentials",
        "status": "error",
    }
    data = request.get_json()
    if auth.authenticate_user_login(data["username"], data["password"]):
        response["message"] = f"Welcome, {data['username']}!"
        response["status"] = "ok"

    return response
