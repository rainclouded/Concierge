import pymongo

from typing import List

from incident_reports_server.models.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory

class IncidentReportPersistenceMongo(IIncidentReportPersistence):
    def __init__(self, db_connection_string: str):
        self.db_client = pymongo.MongoClient(db_connection_string)
        self.concierge_db = self.db_client["concierge"]
        self.ir_collection = self.concierge_db["incident_reports"]
        
    def get_incident_reports(self, severities=None, statuses=None, beforeDate=None, afterDate=None) -> List[IncidentReport]:
        query = {}
        
        if severities:
            query["severity"] = {"$in": [severity.value for severity in severities]}

        if statuses:
            query["status"] = {"$in": [status.value for status in statuses]}

        if beforeDate:
            query["created_at"] = {"$lte": beforeDate}

        if afterDate:
            query.setdefault("created_at", {})
            query["created_at"]["$gte"] = afterDate

        incident_reports = list(self.ir_collection.find(query))

        return [IncidentReportFactory.create_incident_report(report) for report in incident_reports]

    
    def get_incident_report_by_id(self, id:int) -> IncidentReport:
        incident_report = self.ir_collection.find_one({"id" : id})
        
        if incident_report:
            return IncidentReportFactory.create_incident_report(incident_report)
        else:
            return None
    
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        result = self.ir_collection.insert_one(incident_report.to_dict())
        incident_report.set_id(result.inserted_id)
        
        return incident_report 
    
    def update_incident_report(self, id: int, incident_report: IncidentReport) -> IncidentReport:
        result = self.ir_collection.update_one({"id" : id}, {"$set": incident_report.to_dict()})
        
        if result.matched_count > 0:
            return self.get_incident_report_by_id(id)
        else:
            raise ValueError("Incident Report not found!") 
    
    def delete_incident_report(self, id:int) -> None:
        result = self.ir_collection.delete_one({"id": id})
        
        if result.deleted_count == 0:
            raise ValueError("Incident Report not found!")
    