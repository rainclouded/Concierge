
"""
Module for integration tests
"""
import json
import os
import time
import requests

class APIProfiler():
    """
    Integration tests for accounts service
    """

    def __init__(self):
        self.json_tests = os.path.join(
            os.path.dirname(os.path.abspath(__file__)),
            'requests.json'
        )
        self.log_file = os.path.join(
            os.path.dirname(os.path.abspath(__file__)),
            'logs.txt'
        )
        self.timeout = 3

        self.write_logs(
            self.profile(
                self.parse_calls()
            )
        )

    def create_call(self, http_request:dict):
        """
        Create the http call
        """
        return type(
            'http_request',
            (),
            {
                'request_type': http_request['type'],
                'url' : http_request['url'],
                'headers' : http_request['headers'] \
                    if 'headers' in http_request else None,
                'body' : http_request['body'] \
                    if 'body' in http_request else None
            }
        )


    def parse_calls(self):
        """
        Parse all of the requests
        """
        api_calls = []
        with open(self.json_tests,'r', encoding='utf-8') as test_file:
            request_data = json.load(test_file)
            api_calls = [
                self.create_call(request)
                for request in request_data['requests']
            ]
  
        return api_calls


    def profile(self, http_requests:list):
        """
        Call each of the requests
        """
        logs = []
        for http_request in http_requests:
            log = f'Sending {http_request.request_type} to {http_request.url}'
            print(log)
            logs.append(log)
            start_time = time.perf_counter()
            try:
                response = getattr(requests, http_request.request_type)(
                    http_request.url,
                    json= http_request.body if http_request.body else {},
                    headers = http_request.headers if http_request.headers else {},
                    timeout = self.timeout
                )
                end_time = time.perf_counter()
                log = f'Response took {end_time-start_time} seconds: ' +(
                    "Request OK\n"
                    if response.ok
                    else f"Request failed. Code {response.status_code} - "+
                        f"{response.text}"
                    )

            except Exception as e:
                print(e)
                log = 'An exception was raised processing this call'
            print(log)
            logs.append(log)
        return logs
 

    def write_logs(self, logs:list):
        with open(self.log_file, 'w', encoding='utf-8') as log_file:
            log_file.writelines(log+'\n' for log in logs)
