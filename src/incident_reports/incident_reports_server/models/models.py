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
    def __init__(self, id=None, severity=None, status=None, title=None, description=None, created_at=None, updated_at=None, filing_person_id=None, reviewer_id=None):
        self.id = id
        self.severity = severity
        self.status = status
        self.title = title
        self.description = description
        self.created_at = created_at if created_at else datetime.now()
        self.updated_at = updated_at if updated_at else datetime.now()
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
    
    @classmethod
    def from_dict(cls, data: dict):
        return cls(
            id=data.get("id"),  # Get the id from the dictionary
            severity=Severity(data["severity"]),
            status=Status(data["status"]),
            title=data["title"],
            description=data["description"],
            created_at=datetime.fromisoformat(data["created_at"]),
            updated_at=datetime.fromisoformat(data["updated_at"]),
            filing_person_id=data["filing_person_id"],
            reviewer_id=data.get("reviewer_id")  # Defaults to None if not present
        )
        
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
    UNAUTHORIZED = "Your account does not have permission to do this operation."

class PermissionNames:
    VIEW_IR =   "canViewAll"
    EDIT_IR =   "canEditAll"
    CREATE_IR = "canCreate"
    DELETE_IR = "canDelete"
