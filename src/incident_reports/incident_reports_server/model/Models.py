import datetime
from enum import Enum

class Severity(Enum):
    LOW = "Low"
    Medium = "Medium"
    High = "High"
    Critical = "Critical"
    
    from enum import Enum

class Status(Enum):
    OPEN = "Open"
    CLOSED = "Closed"
    RESOLVED = "Resolved"
    IN_PROGRESS = "In Progress"
    
class IncidentReport:
    def __init__(self, id: int, severity: Severity, status: Status, title: str, description: str, filing_person_Id: int, reviewer_Id: int) -> None:
        self.id = id
        self.severity = severity
        self.status = status
        self.title = title
        self.description = description
        self.created_at = datetime.now()
        self.updated_at = self.created_at
        self.filing_person_Id = filing_person_Id
        self.reviewer_Id = reviewer_Id
        
    def update(self, updated_report: 'IncidentReport') -> None:
        self.severity = updated_report.severity
        self.status = updated_report.status
        self.title = updated_report.title
        self.description = updated_report.description
        self.reviewerID = updated_report.reviewerID
        self.updated_at = datetime.now()
        
class IncidentReportResponse:
    def __init__(self, message : str, data) -> None:
        self.message = message
        self.data = data
        self.timeStamp = datetime.now()
        
    