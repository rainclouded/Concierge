from flask import Flask, jsonify, request
import requests
import argparse
import os

app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "SESSIONS_PORT"
ACCOUNTS_SERVICE_ADDRESS = "http://accounts:8080/accounts"


def get_port():
    parser = argparse.ArgumentParser(
        description="A Flask app that returns Session information"
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


@app.route("/sessions/", methods=["GET", "POST"])
def index():
    if request.method == "GET":
        response = {"message": "Hello, World From Sessions", "status": "success"}
        return jsonify(response)
    elif request.method == "POST":
        username = request.json["username"] if "username" in request.json else ""
        password = request.json["password"] if "password" in request.json else ""

        login_result_Request = requests.post(
            f"{ACCOUNTS_SERVICE_ADDRESS}/login_attempts",
            json={"username": username, "password": password},
        )
        login_result = login_result_Request.json()
        status_code = 200

        if "status" in login_result and login_result["status"] == "ok":
            login_result["session_key"] = (
                create_session_key()
            )  # should be returned as a cookie
        else:
            status_code = 400

        return login_result, status_code

    else:
        return jsonify({"status": "error", "message": "Request not supported"})


def create_session_key():  # TODO use persistent volume & db
    return "defaul-tsessi-ionkey-output"


if __name__ == "__main__":
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)
