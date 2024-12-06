from typing import List
from incident_reports_server.models.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence

class IncidentReportPersistenceStub(IIncidentReportPersistence):
    """
    This class implements the IIncidentReportPersistence interface using an in-memory list.
    It is useful for testing purposes where no actual database is needed.
    """
    
    def __init__(self) -> None:
        """
        Initializes the stub persistence layer. This method will also reset the list of incident reports
        and populate it with some initial sample data.
        """
        self.reset()

    def reset(self) -> None:
        """
        Resets the in-memory list of incident reports. Initializes the list with sample data 
        and resets the next available ID to 1.
        """
        self._incident_reports = []
        self._nextId = 1
        
        # Sample incident reports for initial data
        self.create_incident_report(
            IncidentReport(
                severity=Severity.LOW,
                status=Status.OPEN,
                title="Room Maintenance Request",
                description="Guest reported a leaky faucet in Room 203.",
                filing_person_id=301,
                reviewer_id=401
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.MEDIUM,
                status=Status.IN_PROGRESS,
                title="Lost Property",
                description="A guest has reported a lost wallet in the lobby area.",
                filing_person_id=302,
                reviewer_id=402
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.HIGH,
                status=Status.RESOLVED,
                title="Fire Alarm Malfunction",
                description="The fire alarm went off during dinner service; it was a false alarm.",
                filing_person_id=303,
                reviewer_id=403
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.CRITICAL,
                status=Status.CLOSED,
                title="Food Poisoning Incident",
                description="Multiple guests reported food poisoning after dining at the hotel restaurant.",
                filing_person_id=304,
                reviewer_id=404
            )
        )

    def get_incident_reports(self, severities=None, statuses=None, beforeDate=None, afterDate=None) -> List[IncidentReport]:
        """
        Retrieves incident reports from the in-memory list, optionally filtered by severity, status, 
        or creation date.
        
        Arguments:
            severities -- A list of Severity values to filter by (optional).
            statuses -- A list of Status values to filter by (optional).
            beforeDate -- A datetime object to filter reports created before this date (optional).
            afterDate -- A datetime object to filter reports created after this date (optional).
        
        Returns:
            List[IncidentReport]: A list of filtered IncidentReport objects.
        """
        filtered_reports = self._incident_reports
        
        # Apply severity filter if provided
        if severities:
            filtered_reports = [report for report in filtered_reports if report.severity in severities]

        # Apply status filter if provided
        if statuses:
            filtered_reports = [report for report in filtered_reports if report.status in statuses]

        # Apply beforeDate filter if provided
        if beforeDate:
            filtered_reports = [report for report in filtered_reports if report.created_at <= beforeDate]

        # Apply afterDate filter if provided
        if afterDate:
            filtered_reports = [report for report in filtered_reports if report.created_at >= afterDate]
        
        return filtered_reports

    def get_incident_report_by_id(self, id: int) -> IncidentReport:
        """
        Retrieves an incident report by its unique ID from the in-memory list.
        
        Arguments:
            id -- The ID of the incident report to retrieve.
        
        Returns:
            IncidentReport: The incident report matching the provided ID, or None if not found.
        """
        for incident_report in self._incident_reports:
            if incident_report.id == id:
                return incident_report
            
        return None 

    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        """
        Creates a new incident report, assigns it a unique ID, and adds it to the in-memory list.
        
        Arguments:
            incident_report -- The IncidentReport object to create.
        
        Returns:
            IncidentReport: The created IncidentReport object with a unique ID.
        """
        incident_report.set_id(self._nextId)
        self._nextId += 1  # Increment the next ID for future reports

        self._incident_reports.append(incident_report)
        return incident_report

    def update_incident_report(self, id: int, updated_incident_report: IncidentReport) -> IncidentReport:
        """
        Updates an existing incident report in the in-memory list with new data.
        
        Arguments:
            id -- The ID of the incident report to update.
            updated_incident_report -- The new data to update the incident report with.
        
        Returns:
            IncidentReport: The updated IncidentReport object.
        
        Raises:
            ValueError: If the incident report with the specified ID is not found.
        """
        incident_report = self.get_incident_report_by_id(id)

        if incident_report is not None:
            incident_report.update(updated_incident_report)
            return incident_report
        else:
            raise ValueError("Incident report not found!")

    def delete_incident_report(self, id: int) -> None:
        """
        Deletes an incident report by its unique ID from the in-memory list.
        
        Arguments:
            id -- The ID of the incident report to delete.
        
        Raises:
            ValueError: If the incident report with the specified ID is not found.
        """
        for incident_report in self._incident_reports:
            if incident_report.id == id:
                self._incident_reports.remove(incident_report)
                return
            
        raise ValueError("Incident report not found!")

    def clear(self) -> None:
        """
        Resets the in-memory list of incident reports. This method is useful for clearing all data,
        especially in a testing environment.
        """
        self.reset()
