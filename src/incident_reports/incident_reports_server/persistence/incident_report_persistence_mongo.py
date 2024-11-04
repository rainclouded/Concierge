import pymongo
import os 
from typing import List

from incident_reports_server.models.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence

class IncidentReportPersistenceMongo(IIncidentReportPersistence):
    def __init__(self, db_connection_string: str):
        self.db_client = pymongo.MongoClient(db_connection_string)

        self.concierge_db = self.db_client["concierge"]
        self.ir_collection = self.concierge_db["incident_reports"]
        
        #initialize db if empty
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
        query = {}
        
        #get filter parameters if any
        if severities:
            query["severity"] = {"$in": [severity.value for severity in severities]}

        if statuses:
            query["status"] = {"$in": [status.value for status in statuses]}

        if beforeDate:
            query["created_at"] = {"$lte": beforeDate}

        if afterDate:
            query.setdefault("created_at", {})
            query["created_at"]["$gte"] = afterDate

        #then get incident reports that match the requirements of the query
        incident_reports = list(self.ir_collection.find(query))

        return [IncidentReport.from_dict(report) for report in incident_reports]

    
    def get_incident_report_by_id(self, id:int) -> IncidentReport:
        incident_report = self.ir_collection.find_one({"id" : id})
        
        if incident_report:
            return IncidentReport.from_dict(incident_report)
        else:
            return None
    
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        #get highest id in collection
        max_id_document = self.ir_collection.find_one(sort=[("id", -1)])
        next_id = (max_id_document["id"] + 1) if max_id_document else 1
        
        incident_report.set_id(next_id) 
        
        result = self.ir_collection.insert_one(incident_report.to_dict())
        
        return incident_report
    
    def update_incident_report(self, id: int, incident_report: IncidentReport) -> IncidentReport:
        incident_report.set_id(id)
        result = self.ir_collection.update_one({"id" : id}, {"$set": incident_report.to_dict()})
        
        if result.matched_count > 0:
            return self.get_incident_report_by_id(id)
        else:
            raise ValueError("Incident Report not found!") 
    
    def delete_incident_report(self, id:int) -> None:
        result = self.ir_collection.delete_one({"id": id})
        
        if result.deleted_count == 0:
            raise ValueError("Incident Report not found!")
    
    def clear(self) -> None:
        self.ir_collection.delete_many({})
        self.db_client.close()