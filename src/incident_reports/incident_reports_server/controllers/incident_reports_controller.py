from flask import Flask, jsonify, request, make_response
from flask_cors import CORS
from incident_reports_server.application.services import Services
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory
from incident_reports_server.models.models import IncidentReportResponse, ResponseMessages, PermissionNames
from incident_reports_server.validators.incident_report_validator import IncidentReportValidator

def create_app(persistence=None, permissionValidator=None):
    app = Flask(__name__)
    
    _incident_report_persistence = persistence or Services.get_incident_report_persistence()
    _permission_validator = permissionValidator or Services.get_permission_validator()

    @app.route("/incident_reports/", methods=["GET"])
    def get_incident_reports() -> IncidentReportResponse:
        api_key = request.headers.get('X-API-Key')
        if not _permission_validator.validate_session_key_for_permission_name(api_key, PermissionNames.VIEW_IR):
            return jsonify(IncidentReportResponse(ResponseMessages.UNAUTHORIZED, None).to_dict()), 401

        # Get query parameters
        severity_list = request.args.get('severity')
        status_list = request.args.get('status')
        beforeDate = request.args.get('beforeDate')
        afterDate = request.args.get('afterDate')
        
        #convert strings to model objects
        try:
            if severity_list:
                severity_list = IncidentReportFactory.create_severity(severity_list)
                
            if status_list:
                status_list = IncidentReportFactory.create_status(status_list)
                
            if beforeDate:
                beforeDate = IncidentReportFactory.create_date(beforeDate)
                
            if afterDate:
                afterDate = IncidentReportFactory.create_date(afterDate)
        except ValueError:
            #if invalid parameters were passed, return bad request
            response = IncidentReportResponse(ResponseMessages.INVALID_PARAMETERS_PASSED, None).to_dict()
            return jsonify(response), 400
                        
        #get incident reports with added filters(if any)
        incident_reports = _incident_report_persistence.get_incident_reports(severities=severity_list, statuses=status_list, 
                                                                             beforeDate=beforeDate, afterDate=afterDate)

        if incident_reports is None:
            response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORTS_FAILED, None).to_dict()
            return jsonify(response), 404

        #convert list to JSON object
        incident_reports_data = [report.to_dict() for report in incident_reports]
        
        #return incident report list with 200        
        response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORTS_SUCCESS, incident_reports_data).to_dict()
        return jsonify(response), 200
    
    @app.route("/incident_reports/<int:id>", methods=["GET"])
    def get_incident_report_by_id(id: int) -> IncidentReportResponse:
        api_key = request.headers.get('X-API-Key')
        if not _permission_validator.validate_session_key_for_permission_name(api_key, PermissionNames.VIEW_IR):
            return jsonify(IncidentReportResponse(ResponseMessages.UNAUTHORIZED, None).to_dict()), 401
        
        incident_report = _incident_report_persistence.get_incident_report_by_id(id)

        #check if report with passed id exists
        if incident_report is None:
            #if not found, return not found
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404

        incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_SUCCESS, incident_report.to_dict()).to_dict()
        return jsonify(incident_response), 200

    @app.route("/incident_reports/", methods=["POST"])
    def create_incident_report() -> IncidentReportResponse:
        api_key = request.headers.get('X-API-Key')
        if not _permission_validator.validate_session_key_for_permission_name(api_key, PermissionNames.CREATE_IR):
            return jsonify(IncidentReportResponse(ResponseMessages.UNAUTHORIZED, None).to_dict()), 401

        #convert JSON to incident report object
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception:
            response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED,None).to_dict()
            return jsonify(response), 400

        #validate incident report
        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED,None).to_dict()
            return jsonify(response), 400

        #add incident report to database
        incident_report = _incident_report_persistence.create_incident_report(incident_report)

        #create URI that points to the newly created report
        uri = f"{request.scheme}://{request.host}/incident_reports/{incident_report.id}"
        
        response_data = IncidentReportResponse(ResponseMessages.CREATE_INCIDENT_REPORT_SUCCESS,incident_report.to_dict()).to_dict()
        response = make_response(jsonify(response_data), 201)
        response.headers["Location"] = uri

        return response

    @app.route("/incident_reports/<int:id>", methods=["PUT"])
    def update_incident_report(id: int) -> IncidentReportResponse:   
        api_key = request.headers.get('X-API-Key')
        if not _permission_validator.validate_session_key_for_permission_name(api_key, PermissionNames.EDIT_IR):
            return jsonify(IncidentReportResponse(ResponseMessages.UNAUTHORIZED, None).to_dict()), 401

        #convert JSON to incident report object
        try:
            incident_report = IncidentReportFactory.create_incident_report(request.get_json())
        except Exception as e:
            incident_response = IncidentReportResponse(str(e), None).to_dict()
            return jsonify(incident_response), 400

        #validate incident report
        if not IncidentReportValidator.validate_incident_report_parameters(incident_report):
            incident_response = IncidentReportResponse(ResponseMessages.INVALID_INCIDENT_REPORT_PASSED, None).to_dict()
            return jsonify(incident_response), 400

        #find if incident report with passed id exists
        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404
        
        #make call to database to update incident report
        incident_report = _incident_report_persistence.update_incident_report(id, incident_report)

        incident_response = IncidentReportResponse(ResponseMessages.UPDATE_INCIDENT_REPORT_SUCCESS, incident_report.to_dict()).to_dict()
        return jsonify(incident_response), 200

    @app.route("/incident_reports/<int:id>", methods=["DELETE"])
    def delete_incident_report(id: int) -> IncidentReportResponse:
        api_key = request.headers.get('X-API-Key')
        if not _permission_validator.validate_session_key_for_permission_name(api_key, PermissionNames.DELETE_IR):
            return jsonify(IncidentReportResponse(ResponseMessages.UNAUTHORIZED, None).to_dict()), 401

        #find if incident report with passed id exists
        if _incident_report_persistence.get_incident_report_by_id(id) is None:
            incident_response = IncidentReportResponse(ResponseMessages.GET_INCIDENT_REPORT_FAILED, None).to_dict()
            return jsonify(incident_response), 404

        #make call to database to delete incident report
        _incident_report_persistence.delete_incident_report(id)

        incident_response = IncidentReportResponse(ResponseMessages.DELETE_INCIDENT_REPORT_SUCCESS, None).to_dict()
        return jsonify(incident_response), 200
    
    return app

if __name__ == "__main__":
    app = create_app()
    CORS(app)
    app.run(host="0.0.0.0", port=8080, debug=True)