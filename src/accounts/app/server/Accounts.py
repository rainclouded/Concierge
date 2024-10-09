"""
Module for the account server
"""
import argparse
import os
import app.Configs as cfg
from flask import Flask, jsonify, request
from flask_cors import CORS
from app.dto.UserObject import UserObject as User
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService


app = Flask(__name__)
CORS(app)
database = DatabaseController(cfg.create_database())
auth = AuthenticationManager(database)
user_service = UserService(database)

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
