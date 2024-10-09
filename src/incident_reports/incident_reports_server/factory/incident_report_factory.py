import json
from datetime import datetime
from incident_reports_server.validators.incident_report_validator import IncidentReportValidator
from incident_reports_server.models.models import IncidentReport, Severity, Status

class IncidentReportFactory:

    #converts json object to incident report object
    @staticmethod
    def create_incident_report(incident_report_JSON : json) -> IncidentReport:
        required_keys = ["severity","status","title","description","filing_person_id","reviewer_id"]
        
        if incident_report_JSON is None:
            raise ValueError("No JSON object was passed!")
        
        #check if all required keys are present
        missing_keys = [key for key in required_keys if key not in incident_report_JSON]
        if missing_keys:
            raise KeyError(f"Values are missing in incident report: {','.join(missing_keys)}")
        
        #create report instance with added incident report
        result = IncidentReport(
            severity=Severity[incident_report_JSON['severity'].upper()], 
            status=Status[incident_report_JSON['status'].upper()],        
            title=incident_report_JSON['title'],
            description=incident_report_JSON['description'],
            filing_person_id=incident_report_JSON['filing_person_id'],
            reviewer_id=incident_report_JSON['reviewer_id']  
        )
        
        #if id was passed, set id to report
        if "id" in incident_report_JSON and incident_report_JSON["id"]:
            result.set_id(incident_report_JSON["id"])
        
        #if id was passed, set id to report
        if "created_at" in incident_report_JSON and incident_report_JSON["created_at"]:
            result.set_created_at(datetime.strptime(incident_report_JSON["created_at"], '%Y-%m-%dT%H:%M:%S.%f'))
            
        return result
    
    #convert status string into status enum. can pass multiple by separating values with comma
    @staticmethod
    def create_severity(severities: str) -> list:
        severity_list = []
        
        for severity in severities.split(','):
            severity = severity.strip()
            
            if not IncidentReportValidator.validate_incident_report_severity(severity.upper()):
                raise ValueError("Invalid severity passed!")
            
            severity_list.append(Severity[severity.upper()])
            
        return severity_list
    
    #convert status string into status enum. can pass multiple by separating values with comma
    @staticmethod
    def create_status(statuses: str) -> list:
        status_list = []
        
        for status in statuses.split(','):
            status = status.strip()
            
            if not IncidentReportValidator.validate_incident_report_status(status.upper()):
                raise ValueError("Invalid status passed!")
            
            status_list.append(Status[status.upper()])
            
        return status_list

    @staticmethod
    def create_date(date: str):        
        return datetime.strptime(date, "%Y-%m-%d") 

    
        
 
        
        
        
        
     