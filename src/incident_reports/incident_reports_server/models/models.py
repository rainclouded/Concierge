from datetime import datetime
from enum import Enum

class Severity(Enum):
    """
    Enum representing the severity levels of an incident report.
    """
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"
    CRITICAL = "CRITICAL"
    
class Status(Enum):
    """
    Enum representing the possible statuses of an incident report.
    """
    OPEN = "OPEN"
    CLOSED = "CLOSED"
    RESOLVED = "RESOLVED"
    IN_PROGRESS = "IN_PROGRESS"
    
class IncidentReport:
    """
    Represents an incident report.
    
    Attributes:
        id: Unique identifier for the incident report.
        severity: The severity level of the incident.
        status: The current status of the incident.
        title: The title of the incident report.
        description: A detailed description of the incident.
        created_at: Timestamp when the incident report was created.
        updated_at: Timestamp when the incident report was last updated.
        filing_person_id: The ID of the person who filed the incident report.
        reviewer_id: The ID of the person reviewing the incident report.
    """
    def __init__(self, id=None, severity=None, status=None, title=None, description=None, created_at=None, updated_at=None, filing_person_id=None, reviewer_id=None):
        """
        Initializes an IncidentReport object with the provided details.
        
        Arguments:
            id -- The unique identifier for the report (optional).
            severity -- The severity of the incident.
            status -- The current status of the incident.
            title -- The title of the incident report.
            description -- The description of the incident.
            created_at -- The creation timestamp of the incident (optional).
            updated_at -- The last update timestamp of the incident (optional).
            filing_person_id -- The ID of the person who filed the report.
            reviewer_id -- The ID of the person reviewing the report.
        """
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
        """
        Updates the current incident report with the values from another report.
        
        Arguments:
            updated_report -- An IncidentReport object containing the updated data.
        """
        self.severity = updated_report.severity
        self.status = updated_report.status
        self.title = updated_report.title
        self.description = updated_report.description
        self.reviewer_id = updated_report.reviewer_id
        self.updated_at = datetime.now()

    def set_id(self, id: int) -> None:
        """
        Sets the ID for the incident report.
        
        Arguments:
            id -- The ID to set for the incident report.
        """
        self.id = id
        
    def to_dict(self): 
        """
        Converts the incident report to a dictionary representation.
        
        Returns:
            dict: A dictionary containing the incident report details.
        """
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
        """
        Converts a dictionary to an IncidentReport object.
        
        Arguments:
            data -- A dictionary containing the incident report details.
        
        Returns:
            IncidentReport: The corresponding IncidentReport object.
        """
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
    """
    Represents the response for an incident report API request.
    
    Attributes:
        message: The response message indicating the status of the request.
        data: The data (if any) to be returned in the response.
        time_stamp: The timestamp when the response was created.
    """
    def __init__(self, message: str, data):
        """
        Initializes an IncidentReportResponse with a message and data.
        
        Arguments:
            message -- The response message.
            data -- The data to include in the response.
        """
        self.message = message
        self.data = data
        self.time_stamp = datetime.now()
        
    def to_dict(self):
        """
        Converts the response to a dictionary representation.
        
        Returns:
            dict: A dictionary representation of the response.
        """
        return {
            "message": self.message,
            "data": self.data,
            "timestamp": self.time_stamp,
        }
        
class ResponseMessages:
    """
    Contains predefined response messages used in the application.
    """
    GET_INCIDENT_REPORTS_SUCCESS = "Incident reports retrieved successfully."
    GET_INCIDENT_REPORT_SUCCESS = "Incident report retrieved successfully."
    CREATE_INCIDENT_REPORT_SUCCESS = "Incident report created successfully."
    UPDATE_INCIDENT_REPORT_SUCCESS = "Incident report was updated successfully."
    DELETE_INCIDENT_REPORT_SUCCESS = "Incident report deleted successfully."
    GET_INCIDENT_REPORTS_FAILED = "We had trouble fetching your incident reports."
    GET_INCIDENT_REPORT_FAILED = "Incident report with specified id not found."
    INVALID_INCIDENT_REPORT_PASSED = "Bad Request. Incident report with invalid parameters was passed."
    INVALID_PARAMETERS_PASSED = "Bad Request. Invalid parameters were passed."
    UNAUTHORIZED = "Your account does not have permission to do this operation."

class PermissionNames:
    """
    Contains permission names for various operations on incident reports.
    """
    VIEW_IR = "canViewIncidentReports"
    EDIT_IR = "canEditIncidentReports"
    CREATE_IR = "canCreateIncidentReports"
    DELETE_IR = "canDeleteIncidentReports"
