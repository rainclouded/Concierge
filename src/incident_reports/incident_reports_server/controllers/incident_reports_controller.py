from flask import Flask, jsonify, request, make_response
import argparse
import os

from incident_reports_server.application.services import Services
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory
from incident_reports_server.models.models import IncidentReportResponse, ResponseMessages
from incident_reports_server.validators.incident_report_validator import IncidentReportValidator

def create_app(persistence=None):
    app = Flask(__name__)
    
    _incident_report_persistence = persistence or Services.get_incident_report_persistence()
    #_permission_validator = Services.get_permission_validator()

    @app.route("/incident_reports/", methods=["GET"])
    def get_incident_reports() -> IncidentReportResponse:
        severity_list = None
        status_list = None
        beforeDate = None
        afterDate = None
        
        # Get query parameters
        severity_args = request.args.get('severity')
        statuses_args = request.args.get('status')
        beforeDate_args = request.args.get('beforeDate')
        afterDate_args = request.args.get('afterDate')
        
        #convert strings to model objects
        try:
            if severity_args:
                severity_list = IncidentReportFactory.create_severity(severity_args)
                
            if statuses_args:
                status_list = IncidentReportFactory.create_status(statuses_args)
                
            if beforeDate_args:
                beforeDate = IncidentReportFactory.create_date(beforeDate_args)
                
            if afterDate_args:
                afterDate = IncidentReportFactory.create_date(afterDate_args)
        except ValueError:
            response = IncidentReportResponse(ResponseMessages.INVALID_PARAMETERS_PASSED, None).to_dict()
            return jsonify(response), 400
                        
        incident_reports = _incident_report_persistence.get_incident_reports(severities=severity_list, statuses=status_list, 
                                                                             beforeDate=beforeDate, afterDate=afterDate)

        if incident_reports is None:
            response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORTS_FAILED, None).to_dict()
            return jsonify(response), 404

        #convert list to JSON object
        incident_reports_data = [report.to_dict() for report in incident_reports]
        
        response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORTS_SUCCESS, incident_reports_data).to_dict()
        return jsonify(response), 200
    
    @app.route("/incident_reports/<int:id>", methods=["GET"])
    def get_incident_report_by_id(id: int) -> IncidentReportResponse:
        incident_report = _incident_report_persistence.get_incident_report_by_id(id)

        if incident_report is None:
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404

        incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_SUCCESS, incident_report.to_dict()).to_dict()
        return jsonify(incident_response), 200

    @app.route("/incident_reports/", methods=["POST"])
    def create_incident_report() -> IncidentReportResponse:
        #TODO: validate session call
        #if not _permission_validator.validate_permissions(permission, sessionKey)

        #convert JSON to incident report object
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception:
            response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED,None).to_dict()
            return jsonify(response), 400

        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED,None).to_dict()
            return jsonify(response), 400

        incident_report = _incident_report_persistence.create_incident_report(incident_report)

        #create URI that points to the newly created report
        uri = f"{request.scheme}://{request.host}/incident_reports/{incident_report.id}"
        
        response_data = IncidentReportResponse(ResponseMessages.CREATE_INCIDENT_REPORT_SUCCESS,incident_report.to_dict()).to_dict()
        response = make_response(jsonify(response_data), 201)
        response.headers["Location"] = uri

        return response

    @app.route("/incident_reports/<int:id>", methods=["PUT"])
    def update_incident_report(id: int) -> IncidentReportResponse:   
        #TODO: validate session call
        #if not _permission_validator.validate_permissions(permission, sessionKey)

        #convert JSON to incident report object
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception as e:
            incident_response = IncidentReportResponse(str(e), None).to_dict()
            return jsonify(incident_response), 400

        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            incident_response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED, None).to_dict()
            return jsonify(incident_response), 400

        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404
        
        incident_report = _incident_report_persistence.update_incident_report(id, incident_report)

        incident_response = IncidentReportResponse(ResponseMessages.UPDATE_INCIDENT_REPORT_SUCCESS, incident_report.to_dict()).to_dict()
        return jsonify(incident_response), 200

    @app.route("/incident_reports/<int:id>", methods=["DELETE"])
    def delete_incident_report(id: int) -> IncidentReportResponse:
        #TODO: validate session call
        #if not _permission_validator.validate_permissions(permission, sessionKey)

        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404

        _incident_report_persistence.delete_incident_report(id)

        incident_response = IncidentReportResponse(ResponseMessages.DELETE_INCIDENT_REPORT_SUCCESS, None).to_dict()
        return jsonify(incident_response), 200
    
    return app

if __name__ == "__main__":
    app = create_app()
    app.run(host="0.0.0.0", port=8080, debug=True)