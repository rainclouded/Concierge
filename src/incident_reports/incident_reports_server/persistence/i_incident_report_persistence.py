from abc import ABC, abstractmethod
from typing import List

from incident_reports_server.models.models import IncidentReport

class IIncidentReportPersistence(ABC):
    @abstractmethod
    def get_incident_reports(self, severities=None, statuses=None, beforeDate=None, afterDate=None) -> List[IncidentReport]:
        pass
    
    @abstractmethod
    def get_incident_report_by_id(self, id:int) -> IncidentReport:
        pass
    
    @abstractmethod
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        pass 
    
    @abstractmethod
    def update_incident_report(self, id: int, incident_report: IncidentReport) -> IncidentReport:
        pass 
    
    @abstractmethod
    def delete_incident_report(self, id:int) -> None:
        pass
    
    