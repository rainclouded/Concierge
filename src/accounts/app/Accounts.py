from flask import Flask, jsonify, request
import argparse
import os

app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"


attempt_count = 0  # TODO remove state from server

def get_port():
    parser = argparse.ArgumentParser(
        description="A Flask app that returns accounts information"
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
                f"Invalid PORT environment variable: {port}. Using default port {DEFAULT_PORT}."
            )

    return DEFAULT_PORT


@app.route("/accounts/", methods=['GET'])
def index():
    response = {"message": "Hello, World From Accounts", "status": "success"}
    return jsonify(response)


@app.route("/accounts/login_attempts", methods=["POST"])
def login():
    global attempt_count
    response = {
        "key": None,
        "account_id": None,
        "message": f"Login Fail #{attempt_count}: Invalid Credentials",
        "status": "error",
    }
    data = request.get_json()
    if validate_password(data["username"], data["password"]):
        response["message"] = f"Welcome, {data['username']}!"
        response["status"] = "ok"
    else:
      attempt_count+=1
    return response


def validate_password(username, password):
    return username == "admin" and password == "admin"


if __name__ == "__main__":
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)
