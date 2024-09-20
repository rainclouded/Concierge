from flask import Flask, jsonify
import argparse
import os

app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = 'PERMISSIONS_PORT'

def get_port():
    parser = argparse.ArgumentParser(description='A Flask app that returns permission information')
    parser.add_argument('-p', '--port', type=int, help='Port number for the server to listen on')
    args = parser.parse_args()

    if args.port and (0 <= args.port <= 65536):
        return args.port

    port = os.getenv(ENVIRONMENT_VAR_NAME_PORT)
    if port and (0 <= port <= 65536):
        try:
            return int(port)
        except ValueError:
            print(f"Invalid PORT environment variable: {port}. Using default port {DEFAULT_PORT}.")
    
    return DEFAULT_PORT

@app.route('/permissions/', methods=['GET'])
def index():
    response = {
        "message": "Hello, World From Permissions",
        "status": "success"
    }
    return jsonify(response)

if __name__ == '__main__':
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host='0.0.0.0', port=port)
