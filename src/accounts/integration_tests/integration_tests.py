
"""
Testing module for FlaskApp
"""
import unittest
import datetime
import jwt
import json
import os
import re
import requests
import traceback
from os import getenv
from pymongo import MongoClient
class IntegrationTests():
    """
    Tests for FlaskApp
    """
    tests = []
    class TestFailureError(Exception):
        def __init__(self, error_message):
            super().__init__(error_message)

    def __init__(self):
        self.json_tests = os.path.join(
            os.path.dirname(os.path.abspath(__file__)),
            'tests.json'
        )
        self.db_url = getenv('DB_URI')
        self.service_url = 'http://localhost:8080/accounts'#'http://localhost:50001/accounts'
        self.tests_passed = 0
        self.tests_run = 0
        self.timeout = 3

    def update_stats(self, test_passed):
        if test_passed:
            self.tests_passed += 1
        self.tests_run += 1

    def print_stats_and_exit(self):
        percentage = 100*(
            self.tests_passed/self.tests_run
            if self.tests_run
            else 0
        )
        print(
            f"Tests ran: {self.tests_run}\n"+
            f"Tests passed: {self.tests_passed}\n"+
            f"Pass percentage: {percentage}%\n"
            )
        exit(0 if percentage == 100 else 1)

    def run_all(self):
        self.exit_on_failure(self.test_health_check)
        self.integration_tests()
        self.print_stats_and_exit()


    def exit_on_failure(self, func):
        try:
            func()
        except self.TestFailureError as e:
            print(f"{traceback.format_exc()}\n{e}\nHealth check failed. Exiting.")
            exit(0)

    #start of thes

    def test_health_check(self):
        try:
            database_client = MongoClient(self.db_url, serverSelectionTimeoutMS=5000)
            database = database_client['accounts']
            database_client.server_info()
            collections = database.list_collection_names()
            for collection in collections:
                database[collection].drop()
            collection = database.create_collection('accounts')
            collection = database['accounts']
            collection.create_index([("username", 1)])

        except ConnectionError:
            raise self.TestFailureError("Database could not be accessed")
        
        self.update_stats(True)

        try:
            database_health = requests.get(self.service_url, timeout=5)
            database_health.raise_for_status()

        except requests.HTTPError:
            raise self.TestFailureError("Accounts Service could not be accessed")     
        
        self.update_stats(True)

    def integration_tests(self):
        with open(self.json_tests,'r', encoding='utf-8') as test_file:
            test_data = json.load(test_file)
            for test_package in test_data['test_packages']:
                self.update_stats(self.call_test(test_package))
        #create
        staff_data = {
            'username' : 'staffman1',
            'type' : 'staff',
            'password' : 'LongPassword99'
        }
        guest_data = {
            'username' : '409',
            'type' : 'guest'
        }
        #response = requests.post(self.service_url, json=staff_data, timeout=10)
        #print(response.text)
        #login
        #delete
        #create
        #update
    def validate(self, response, expected_data, expected_code):
        
        expected_data = expected_data.replace('*', '.*')    
        regex_data = re.compile(expected_data)
        match_found = regex_data.search(response.text)

        print(f"""
            regex data: {regex_data}
            compare: {response.text}
            match found: {match_found}\n
            """)
        return bool(match_found) and expected_code == response.status_code
            
            
    def call_test(self, test_data:dict):
        headers = {}
        
        expected_response_number = test_data.get('response_number')
        expected_response_data = test_data.get('response_data')
        request_endpoint = self.service_url + test_data.get('endpoint', '')
        description = test_data.get('desc')
        request_message = test_data.get('request_message', {})
        request_type = test_data.get('request_type', None)

        if 'headers' in test_data:
            for header in test_data['headers']:
                headers[header['name']] = header['value']
        
        response = None
        match request_type:
            case 'get':
                response = requests.get(request_endpoint, headers=headers)
            case 'post':
                response = requests.post(request_endpoint, json=request_message, headers=headers)
            case 'put':
                response = requests.put(request_endpoint, json=request_message, headers=headers)
            case _:
                pass
        print(f"Testing: {description}")

        return self.validate(response, expected_response_data, expected_response_number)
        
