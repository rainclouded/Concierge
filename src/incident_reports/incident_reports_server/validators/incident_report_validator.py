from datetime import datetime
from incident_reports_server.models.models import IncidentReport, Severity, Status

class IncidentReportValidator:
    status_values = [status.value for status in Status]
    severity_values = [severity.value for severity in Severity]
        
    @staticmethod
    def validate_incident_report_parameters(incident_report: IncidentReport) -> bool:
        #validate each report variable in this class, check if type matches
        
        #check if enum value matches
        if not isinstance(incident_report.status, Status) or incident_report.status.value not in IncidentReportValidator.status_values:
            return False
        
        #check if enum value matches
        if not isinstance(incident_report.severity, Severity) or incident_report.severity.value not in IncidentReportValidator.severity_values:
            return False
        
        #check for empty string
        if not isinstance(incident_report.title, str) or not incident_report.title.strip():
            return False
        
        #check for empty string
        if not isinstance(incident_report.description, str) or not incident_report.description.strip():
            return False
        
        if not isinstance(incident_report.created_at, datetime):
            return False
        
        #check for negative id
        if not isinstance(incident_report.filing_person_id, int) or incident_report.filing_person_id <= 0:
            return False
        
        #check for negative id
        if not isinstance(incident_report.reviewer_id, int) or incident_report.reviewer_id <= 0:
            return False
        
        return True
    
    @staticmethod
    def validate_incident_report_severity(severity : str) -> bool:    
        #check if value is in the range of stated enum values 
        if severity not in IncidentReportValidator.severity_values:
            return False
        
        return True
    
    @staticmethod
    def validate_incident_report_status(severity : str) -> bool:  
        #check if value is in the range of stated enum values    
        if severity not in IncidentReportValidator.status_values:
            return False
        
        return True
    
    
    
    
    
