import json
from datetime import datetime
from incident_reports_server.validators.incident_report_validator import IncidentReportValidator
from incident_reports_server.models.models import IncidentReport, Severity, Status

class IncidentReportFactory:

    @staticmethod
    def create_incident_report(incident_report_JSON: json) -> IncidentReport:
        """
        Converts a JSON object to an IncidentReport object.
        
        Arguments:
            incident_report_JSON -- The JSON object containing the incident report data.
        
        Returns:
            IncidentReport: The created IncidentReport object.

        Raises:
            ValueError: If no JSON object is passed.
            KeyError: If required keys are missing from the JSON object.
        """
        required_keys = ["severity", "status", "title", "description", "filing_person_id", "reviewer_id"]
        
        if incident_report_JSON is None:
            raise ValueError("No JSON object was passed!")
        
        missing_keys = [key for key in required_keys if key not in incident_report_JSON]
        if missing_keys:
            raise KeyError(f"Values are missing in incident report: {','.join(missing_keys)}")
        
        result = IncidentReport(
            severity=Severity[incident_report_JSON['severity'].upper()], 
            status=Status[incident_report_JSON['status'].upper()],        
            title=incident_report_JSON['title'],
            description=incident_report_JSON['description'],
            filing_person_id=incident_report_JSON['filing_person_id'],
            reviewer_id=incident_report_JSON['reviewer_id']  
        )
        
        if "id" in incident_report_JSON:
            result.set_id(incident_report_JSON["id"])
        
        return result
    
    @staticmethod
    def create_severity(severities: str) -> list:
        """
        Converts a string of severities into a list of Severity enum values.
        
        Arguments:
            severities -- A comma-separated string representing severities.
        
        Returns:
            list: A list of Severity enum values.
        
        Raises:
            ValueError: If any severity is invalid.
        """
        severity_list = []
        
        for severity in severities.split(','):
            severity = severity.strip()
            
            if not IncidentReportValidator.validate_incident_report_severity(severity.upper()):
                raise ValueError("Invalid severity passed!")
            
            severity_list.append(Severity[severity.upper()])
            
        return severity_list
    
    @staticmethod
    def create_status(statuses: str) -> list:
        """
        Converts a string of statuses into a list of Status enum values.
        
        Arguments:
            statuses -- A comma-separated string representing statuses.
        
        Returns:
            list: A list of Status enum values.
        
        Raises:
            ValueError: If any status is invalid.
        """
        status_list = []
        
        for status in statuses.split(','):
            status = status.strip()
            
            if not IncidentReportValidator.validate_incident_report_status(status.upper()):
                raise ValueError("Invalid status passed!")
            
            status_list.append(Status[status.upper()])
            
        return status_list

    @staticmethod
    def create_date(date: str):
        """
        Converts a string representation of a date into a datetime object.
        
        Arguments:
            date -- A string representing the date in the format "YYYY-MM-DD".
        
        Returns:
            datetime: The corresponding datetime object.
        """
        return datetime.strptime(date, "%Y-%m-%d")
