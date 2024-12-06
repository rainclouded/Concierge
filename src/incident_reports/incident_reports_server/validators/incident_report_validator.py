from datetime import datetime
from incident_reports_server.models.models import IncidentReport, Severity, Status

class IncidentReportValidator:
    # List of valid status values from the Status enum
    status_values = [status.value for status in Status]
    
    # List of valid severity values from the Severity enum
    severity_values = [severity.value for severity in Severity]
        
    @staticmethod
    def validate_incident_report_parameters(incident_report: IncidentReport) -> bool:
        """
        Validates the parameters of an IncidentReport object to ensure all fields are correctly typed
        and within valid ranges.

        Arguments:
            incident_report -- The IncidentReport instance to validate.
        
        Returns:
            bool: True if all parameters are valid, False otherwise.
        """
        
        # Validate the 'status' field - must be a valid Status enum value
        if not isinstance(incident_report.status, Status) or incident_report.status.value not in IncidentReportValidator.status_values:
            return False
        
        # Validate the 'severity' field - must be a valid Severity enum value
        if not isinstance(incident_report.severity, Severity) or incident_report.severity.value not in IncidentReportValidator.severity_values:
            return False
        
        # Validate the 'title' field - must be a non-empty string
        if not isinstance(incident_report.title, str) or not incident_report.title.strip():
            return False
        
        # Validate the 'description' field - must be a non-empty string
        if not isinstance(incident_report.description, str) or not incident_report.description.strip():
            return False
        
        # Validate the 'created_at' field - must be a datetime object
        if not isinstance(incident_report.created_at, datetime):
            return False
        
        # Validate the 'filing_person_id' field - must be a positive integer
        if not isinstance(incident_report.filing_person_id, int) or incident_report.filing_person_id <= 0:
            return False
        
        # Validate the 'reviewer_id' field - must be a positive integer
        if not isinstance(incident_report.reviewer_id, int) or incident_report.reviewer_id <= 0:
            return False
        
        # If all checks pass, return True
        return True
    
    @staticmethod
    def validate_incident_report_severity(severity: str) -> bool:
        """
        Validates the severity value to ensure it matches one of the predefined severity levels.

        Arguments:
            severity -- The severity level as a string.
        
        Returns:
            bool: True if the severity is valid, False otherwise.
        """
        # Check if the severity is in the list of valid severity values
        if severity not in IncidentReportValidator.severity_values:
            return False
        
        return True
    
    @staticmethod
    def validate_incident_report_status(status: str) -> bool:
        """
        Validates the status value to ensure it matches one of the predefined status levels.

        Arguments:
            status -- The status level as a string.
        
        Returns:
            bool: True if the status is valid, False otherwise.
        """
        # Check if the status is in the list of valid status values    
        if status not in IncidentReportValidator.status_values:
            return False
        
        return True
