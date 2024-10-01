from typing import List
from incident_reports.incident_reports_server.application import Services
from flask import Flask, jsonify, request
import argparse
import os

from incident_reports.incident_reports_server.factory import IncidentReportFactory
from incident_reports.incident_reports_server.model.Models import IncidentReport, IncidentReportResponse
from incident_reports.incident_reports_server.validators import IncidentReportValidator

app = Flask(__name__)

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "ACCOUNTS_PORT"
_incident_report_persistence = None

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

@app.route("/incident_reports/", methods=["GET"])
def get_incident_reports() -> IncidentReportResponse:
    incident_reports = _incident_report_persistence.get_incident_reports()
    
    if incident_reports is None:
        return jsonify(IncidentReportResponse("Incident reports were failed to be retrieved", None)), 404
    
    return jsonify(IncidentReportResponse("Incident reports retrieved successfully",incident_reports)), 200

@app.route("/incident_reports/<int:id>", methods=["GET"])
def get_incident_report_by_id() -> IncidentReportResponse:
    incident_report = _incident_report_persistence.get_incident_reports(id)
    
    if incident_report is None:
        return jsonify(IncidentReportResponse("Incident report with specified id not found", None)), 404
    
    return jsonify(IncidentReportResponse("Incident report retrieved successfully",incident_report)), 200
 

@app.route("/incident_reports/", methods=["POST"])
def create_incident_report() -> IncidentReportResponse:
    #TODO: validate session
    #create object
    try:
        incident_report = IncidentReportFactory.create_incident_report(request.get_json())
    except Exception as e:
        return jsonify(IncidentReportResponse(str(e).message,None)), 400
    
    #validate object
    if not IncidentReportValidator.validate_new_incident_report(incident_report):
        return jsonify(IncidentReportResponse("Invalid incident report passed!",None)), 400

    _incident_report_persistence.create_incident_report(incident_report)
    
    #create uri that points to newly created report
    uri = f"{request.scheme}://{request.host}/amenities/{incident_report.id}"    
    
    return jsonify(IncidentReportResponse("Incident Report succesfully created.", incident_report), 201, {"Location": uri})

@app.route("/incident_reports/<int:id>", methods=["PUT"])
def update_incident_report() -> IncidentReportResponse:
    #TODO: validate session
    
    #create object
    try:
        incident_report = IncidentReportFactory.create_incident_report(request.get_json())
    except Exception as e:
        return jsonify(IncidentReportResponse(str(e).message,None)), 400
    
    #validate object
    if not IncidentReportValidator.validate_new_incident_report(incident_report):
        return jsonify(IncidentReportResponse("Invalid incident report passed!",None)), 400
    
    _incident_report_persistence.update_incident_report(incident_report)
     
    return jsonify(IncidentReportResponse("Incident Report succesfully updated.", incident_report), 200)


@app.route("/incident_reports/<int:id>", methods=["DELETE"])
def delete_incident_report() -> IncidentReportResponse:
    incident_report = _incident_report_persistence.get_incident_report_by_id(id)
    
    if incident_report is None:
        return jsonify(IncidentReportResponse("Non-existent incident report requested to be deleted.", None)), 400

    incident_report = _incident_report_persistence.delete_incident_report(id)
    return jsonify(IncidentReportResponse("Incident report deleted successfully", None)), 200
 

if __name__ == "__main__":
    _incident_report_persistence = Services.get_incident_report_persistence()
    port = get_port()
    print(f"Starting server on port {port}...")
    app.run(host="0.0.0.0", port=port)
