import pymongo

from typing import List

from incident_reports_server.models.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence

class IncidentReportPersistenceMongo(IIncidentReportPersistence):
    def get_incident_reports(self, severities=None, statuses=None, beforeDate=None, afterDate=None) -> List[IncidentReport]:
        pass
    
    def get_incident_report_by_id(self, id:int) -> IncidentReport:
        pass
    
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        pass 
    
    def update_incident_report(self, id: int, incident_report: IncidentReport) -> IncidentReport:
        pass 
    
    def delete_incident_report(self, id:int) -> None:
        pass