import datetime
from incident_reports.incident_reports_server.model import Severity, Status


class IncidentReport:
    def __init__(self, ID: int, severity: Severity, status: Status, title: str, description: str, 
                created_at: datetime, filingPersonID: int, reviewerID: int):
        self.ID = ID
        self.severity = severity
        self.status = status
        self.title = title
        self.description = description
        self.created_at = created_at
        self.updated_at = created_at
        self.filingPersonID = filingPersonID
        self.reviewerID = reviewerID
        
    def update_incident_report(self, updatedReport: 'IncidentReport'):
        self.severity = updatedReport.severity
        self.status = updatedReport.status
        self.title = updatedReport.title
        self.description = updatedReport.description
        self.reviewerID = updatedReport.reviewerID
        self.updated_at = datetime.now()
        
    
        