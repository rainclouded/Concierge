from flask import Flask, jsonify, request, make_response
import argparse
import os

from incident_reports_server.application.services import Services
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory
from incident_reports_server.models.models import IncidentReportResponse
from incident_reports_server.validators.incident_report_validator import IncidentReportValidator

DEFAULT_PORT = 8080
ENVIRONMENT_VAR_NAME_PORT = "INCIDENT_REPORTS_PORT"

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

def create_app(persistence=None):
    app = Flask(__name__)
    
    _incident_report_persistence = persistence or Services.get_incident_report_persistence()

    @app.route("/incident_reports/", methods=["GET"])
    def get_incident_reports() -> IncidentReportResponse:
        incident_reports = _incident_report_persistence.get_incident_reports()

        if incident_reports is None:
            return jsonify(IncidentReportResponse("Incident reports failed to be retrieved", None).to_dict()), 404

        incident_reports_data = [report.to_dict() for report in incident_reports]
        return jsonify(IncidentReportResponse("Incident reports retrieved successfully", incident_reports_data).to_dict()), 200

    @app.route("/incident_reports/<int:id>", methods=["GET"])
    def get_incident_report_by_id(id: int) -> IncidentReportResponse:
        incident_report = _incident_report_persistence.get_incident_report_by_id(id)

        if incident_report is None:
            return jsonify(IncidentReportResponse("Incident report with specified id not found", None).to_dict()), 404

        return jsonify(IncidentReportResponse("Incident report retrieved successfully", incident_report.to_dict()).to_dict()), 200

    @app.route("/incident_reports/", methods=["POST"])
    def create_incident_report() -> IncidentReportResponse:
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception as e:
            return jsonify(IncidentReportResponse(str(e), None).to_dict()), 400

        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            return jsonify(IncidentReportResponse("Invalid incident report passed!", None).to_dict()), 400

        incident_report = _incident_report_persistence.create_incident_report(incident_report)

        uri = f"{request.scheme}://{request.host}/incident_reports/{incident_report.id}"
        
        response_data = IncidentReportResponse("Incident Report successfully created.", incident_report.to_dict()).to_dict()
        response = make_response(jsonify(response_data), 201)
        response.headers["Location"] = uri

        return response

    @app.route("/incident_reports/<int:id>", methods=["PUT"])
    def update_incident_report(id: int) -> IncidentReportResponse:        
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception as e:
            return jsonify(IncidentReportResponse(str(e), None).to_dict()), 400

        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            return jsonify(IncidentReportResponse("Invalid incident report passed!", None).to_dict()), 400

        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            return jsonify(IncidentReportResponse("Incident report not found.", None).to_dict()), 404
        
        incident_report = _incident_report_persistence.update_incident_report(id, incident_report)

        return jsonify(IncidentReportResponse("Incident Report successfully updated.", incident_report.to_dict()).to_dict()), 200

    @app.route("/incident_reports/<int:id>", methods=["DELETE"])
    def delete_incident_report(id: int) -> IncidentReportResponse:
        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            return jsonify(IncidentReportResponse("Incident report not found.", None).to_dict()), 404

        _incident_report_persistence.delete_incident_report(id)
        return jsonify(IncidentReportResponse("Incident report deleted successfully", None).to_dict()), 200

    return app

if __name__ == "__main__":
    port = get_port()
    print(f"Starting server on port {port}...")
    app = create_app()
    app.run(host="0.0.0.0", port=port)