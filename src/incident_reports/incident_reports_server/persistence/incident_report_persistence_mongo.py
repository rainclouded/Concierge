import pymongo
from typing import List

from incident_reports_server.models.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence

class IncidentReportPersistenceMongo(IIncidentReportPersistence):
    """
    This class provides persistence logic for incident reports, using MongoDB as the backend.
    It allows for basic CRUD operations such as creating, reading, updating, and deleting incident reports.
    """
    
    def __init__(self, db_connection_string: str):
        """
        Initializes the MongoDB client and prepares the database and collection for incident reports.
        
        Arguments:
            db_connection_string -- The MongoDB connection string to connect to the database.
        """
        self.db_client = pymongo.MongoClient(db_connection_string)

        self.concierge_db = self.db_client["concierge"]
        self.ir_collection = self.concierge_db["incident_reports"]
        
        # Initialize DB with sample data if the collection is empty
        if self.ir_collection.count_documents({}) == 0:
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
        Retrieves incident reports based on various filters such as severity, status, and creation date.
        
        Arguments:
            severities -- A list of Severity values to filter by (optional).
            statuses -- A list of Status values to filter by (optional).
            beforeDate -- A datetime object to filter reports created before this date (optional).
            afterDate -- A datetime object to filter reports created after this date (optional).
        
        Returns:
            List[IncidentReport]: A list of IncidentReport objects that match the filters.
        """
        query = {}

        # Apply severity filter if provided
        if severities:
            query["severity"] = {"$in": [severity.value for severity in severities]}

        # Apply status filter if provided
        if statuses:
            query["status"] = {"$in": [status.value for status in statuses]}

        # Apply beforeDate filter if provided
        if beforeDate:
            query["created_at"] = {"$lte": beforeDate}

        # Apply afterDate filter if provided
        if afterDate:
            query.setdefault("created_at", {})
            query["created_at"]["$gte"] = afterDate

        # Fetch incident reports that match the query
        incident_reports = list(self.ir_collection.find(query))

        return [IncidentReport.from_dict(report) for report in incident_reports]
    
    def get_incident_report_by_id(self, id: int) -> IncidentReport:
        """
        Retrieves a specific incident report by its unique ID.
        
        Arguments:
            id -- The unique ID of the incident report to retrieve.
        
        Returns:
            IncidentReport: The incident report matching the provided ID, or None if not found.
        """
        incident_report = self.ir_collection.find_one({"id": id})
        
        if incident_report:
            return IncidentReport.from_dict(incident_report)
        else:
            return None
    
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        """
        Creates a new incident report and inserts it into the database.
        
        Arguments:
            incident_report -- The IncidentReport object to insert into the database.
        
        Returns:
            IncidentReport: The created IncidentReport object with an assigned ID.
        """
        # Get the next available ID based on the highest existing ID
        max_id_document = self.ir_collection.find_one(sort=[("id", -1)])
        next_id = (max_id_document["id"] + 1) if max_id_document else 1
        
        incident_report.set_id(next_id) 
        
        result = self.ir_collection.insert_one(incident_report.to_dict())
        
        return incident_report
    
    def update_incident_report(self, id: int, incident_report: IncidentReport) -> IncidentReport:
        """
        Updates an existing incident report with the provided data.
        
        Arguments:
            id -- The ID of the incident report to update.
            incident_report -- The updated IncidentReport object.
        
        Returns:
            IncidentReport: The updated IncidentReport object.
        
        Raises:
            ValueError: If the incident report with the specified ID is not found.
        """
        incident_report.set_id(id)
        result = self.ir_collection.update_one({"id": id}, {"$set": incident_report.to_dict()})
        
        if result.matched_count > 0:
            return self.get_incident_report_by_id(id)
        else:
            raise ValueError("Incident Report not found!") 
    
    def delete_incident_report(self, id: int) -> None:
        """
        Deletes an incident report from the database by its unique ID.
        
        Arguments:
            id -- The unique ID of the incident report to delete.
        
        Raises:
            ValueError: If no incident report with the specified ID is found.
        """
        result = self.ir_collection.delete_one({"id": id})
        
        if result.deleted_count == 0:
            raise ValueError("Incident Report not found!")
    
    def clear(self) -> None:
        """
        Clears all incident reports from the collection. This operation is irreversible.
        """
        self.ir_collection.delete_many({})
        self.db_client.close()
