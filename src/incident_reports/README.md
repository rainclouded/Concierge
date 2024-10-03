how to run

locally
in src/incident_reports
python -m incident_reports_server.controllers.incident_reports_controller

docker-compose

tests
python -m unittest discover -s .\incident_reports_tests\ -p "*.py"