from datetime import datetime
from enum import Enum

class Severity(Enum):
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"
    CRITICAL = "CRITICAL"
    
class Status(Enum):
    OPEN = "OPEN"
    CLOSED = "CLOSED"
    RESOLVED = "RESOLVED"
    IN_PROGRESS = "IN_PROGRESS"
    
class IncidentReport:
    def __init__(self, severity: Severity, status: Status, title: str, description: str, filing_person_id: int, reviewer_id: int) -> None:
        self.severity = severity
        self.status = status
        self.title = title
        self.description = description
        self.created_at = datetime.now()
        self.updated_at = self.created_at
        self.filing_person_id = filing_person_id
        self.reviewer_id = reviewer_id
        
    def update(self, updated_report: 'IncidentReport') -> None:
        self.severity = updated_report.severity
        self.status = updated_report.status
        self.title = updated_report.title
        self.description = updated_report.description
        self.reviewer_id = updated_report.reviewer_id
        self.updated_at = datetime.now()

    def set_id(self, id : int) -> None:
        self.id = id
        
    def set_created_at(self, created_at : datetime) -> None:
        self.created_at = created_at
        
    def to_dict(self): 
        return {
            "id": getattr(self, 'id', None),
            "severity": self.severity.value,
            "status": self.status.value,        
            "title": self.title,
            "description": self.description,
            "created_at": self.created_at.isoformat(),
            "updated_at": self.updated_at.isoformat(),
            "filing_person_id": self.filing_person_id,
            "reviewer_id": self.reviewer_id,
         }
        
class IncidentReportResponse:
    def __init__(self, message : str, data):
        self.message = message
        self.data = data
        self.time_stamp = datetime.now()
        
    def to_dict(self):
        return {
            "message": self.message,
            "data": self.data,
            "timestamp": self.time_stamp
        }
        
class ResponseMessages:
    GET_INCIDENT_REPORTS_SUCCESS = "Incident reports retrieved successfully."
    GET_INCIDENT_REPORT_SUCCESS = "Incident report retrieved successfully."
    CREATE_INCIDENT_REPORT_SUCCESS = "Incident report created successfully."
    UPDATE_INCIDENT_REPORT_SUCCESS = "Incident report was updated successfully."
    DELETE_INCIDENT_REPORT_SUCCESS = "Incident report deleted successfully."
    GET_INCIDENT_REPORTS_FAILED = "We had trouble fetching your incident reports."
    GET_INCIDENT_REPORT_FAILED = "Incident report with specified id not found."
    INVALID_INCIDENT_REPORT_PASSED = "Bad Request. Incident report with invalid parameters was passed."
    INVALID_PARAMETERS_PASSED = "Bad Request. Invalid parameters was passed."

    