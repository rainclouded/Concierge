
"""
Testing module for FlaskApp
"""
import unittest
import datetime
import jwt
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
        self.db_url = getenv('DB_URI')
        self.service_url = 'http://localhost:8080/accounts'#'http://localhost:50001/accounts'
        self.tests_passed = 0
        self.tests_run = 0

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
            database_client.server_info()
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
        response = requests.post(self.service_url, json=staff_data, timeout=10)
        print(response.text)
        #login
        #delete
        #create
        #update



