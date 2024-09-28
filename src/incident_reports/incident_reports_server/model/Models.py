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
    def __init__(self, ID: int, severity: Severity, status: Status, title: str, description: str, 
                created_at: datetime, filing_person_ID: int, reviewer_ID: int):
        self.ID = ID
        self.severity = severity
        self.status = status
        self.title = title
        self.description = description
        self.created_at = created_at
        self.updated_at = created_at
        self.filing_person_ID = filing_person_ID
        self.reviewer_ID = reviewer_ID
        
    def update(self, updated_report: 'IncidentReport'):
        self.severity = updated_report.severity
        self.status = updated_report.status
        self.title = updated_report.title
        self.description = updated_report.description
        self.reviewerID = updated_report.reviewerID
        self.updated_at = datetime.now()
    