import datetime
from incident_reports_server.model.models import IncidentReport, Severity, Status

class IncidentReportValidator:
    @staticmethod
    def validate_incident_report_parameters(self, incident_report: IncidentReport) -> bool:
        status_values = [status.value for status in Status]
        severity_values = [severity.value for severity in Severity]
        
        if not isinstance(incident_report.status, Status) or IncidentReport.status not in status_values:
            return False
        
        if not isinstance(incident_report.severity, Severity) or IncidentReport.severity not in severity_values:
            return False
        
        if not isinstance(incident_report.title, str) or not incident_report.title.strip():
            return False
        
        if not isinstance(incident_report.description, str) or not incident_report.description.strip():
            return False
        
        if not isinstance(incident_report.created_at, datetime):
            return False
        
        if not isinstance(incident_report.filing_person_ID, int) or incident_report.filing_person_ID <= 0:
            return False
        
        if not isinstance(incident_report.reviewer_ID, int) or incident_report.reviewer_ID <= 0:
            return False
        
        return True
    
