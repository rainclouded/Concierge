from flask import Flask, jsonify, request
import argparse
import os
from . import AuthenticationManager


app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"

auth = AuthenticationManager()

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
    response = {
        "message": "You have contacted the accounts", 
        "status": "success"
        }
    return jsonify(response)


@app.route("/accounts/create_account", methods=["POST"])
def create():
    response = {
        "message" : "Could not create user - {reason}",
        "status" : "error"
    }

    data = request.get_json()
    new_user  = {
        'username' : data["username"],
        'password' : data["password"],
        'type' : data['type']
    }
    if new_user['type'] == 'guest':
        pass
    else:
        if (
            auth.validate_staff_password(new_user["password"]) 
            and auth.validate_new_user(new_user)
        ):
            pass



@app.route("/accounts/login_attempt", methods=["POST"])
def login():
    response = {
        "user_id": None,
        "message": f"Login Fail - Invalid Credentials",
        "status": "error",
    }
    data = request.get_json()
    if auth.validate_staff_login(data["username"], data["password"]):
        response["message"] = f"Welcome, {data['username']}!"
        response["status"] = "ok"

    return response



if __name__ == "__main__":
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)
