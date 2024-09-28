import datetime
from incident_reports.incident_reports_server.application import Services
from incident_reports.incident_reports_server.model.Models import IncidentReport, Severity, Status

class IncidentReportValidator:
    def __init__(self) -> None:
        _incident_report_persistence = None;

    def validate_incident_report_parameters(self, incident_report: IncidentReport) -> bool:
        if not isinstance(incident_report.ID, int) or incident_report.ID <= 0:
            return False
        
        if not isinstance(incident_report.severity, Severity):
            return False
        
        if not isinstance(incident_report.status, Status):
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
    
    def validate_new_incident_report(self, new_incident_report: IncidentReport) -> bool:
        self._incident_report_persistence = Services.get_incident_report_persistence()

        if(not self.validate_incident_report_parameters(new_incident_report)):
            return False
        
        #check if id exists or not
        return not self.get_incident_report_by_id(new_incident_report.ID) #we don't want to override an existing id
    
    def validate_existing_incident_report(self, new_incident_report: IncidentReport) -> bool:
        self._incident_report_persistence = Services.get_incident_report_persistence()

        if(not self.validate_incident_report_parameters(new_incident_report)):
            return False
        
        #check if id exists or not
        return self.get_incident_report_by_id(new_incident_report.ID) #want to make sure incident report exists in database
           
    
