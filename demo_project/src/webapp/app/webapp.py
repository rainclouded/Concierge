from flask import Flask, jsonify, render_template, request
import requests
import argparse
import os

app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = 'WEBAPP_PORT'
ACCOUNTS_SERVICE_ADDRESS = 'http://users:8080' #raw python -p 50000:'localhost:50001';  k8s 'users-service'

def get_port():
    parser = argparse.ArgumentParser(description='A Flask app that renders and returns webapp pages')
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

@app.route('/')
def index():
   # file_path = os.path.join(os.path.dirname(__file__), 'app', 'login.html')
   return render_template('login.html')

# @app.route('/login', methods=['POST'])
# def login():
#    external_response = requests.post(f'{ACCOUNTS_SERVICE_ADDRESS}/login', json=({'username':request.json['password'], 'password':request.json['password']}))
#    return external_response.json()

if __name__ == '__main__':
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host='0.0.0.0', port=port)
