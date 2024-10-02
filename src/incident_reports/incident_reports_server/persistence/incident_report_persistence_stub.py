from typing import List
from incident_reports_server.model.models import IncidentReport, Severity, Status
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence

class IncidentReportPersistenceStub(IIncidentReportPersistence):

    def __init__(self) -> None:
        self._incident_reports = []
        self._nextId = 1
        self.create_incident_report(
            IncidentReport(
                severity=Severity.LOW,
                status=Status.OPEN,
                title="Room Maintenance Request",
                description="Guest reported a leaky faucet in Room 203.",
                filing_person_Id=301, 
                reviewer_Id=401 
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.MEDIUM,
                status=Status.IN_PROGRESS,
                title="Lost Property",
                description="A guest has reported a lost wallet in the lobby area.",
                filing_person_Id=302,
                reviewer_Id=402
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.HIGH,
                status=Status.RESOLVED,
                title="Fire Alarm Malfunction",
                description="The fire alarm went off during dinner service; it was a false alarm.",
                filing_person_Id=303,
                reviewer_Id=403
            )
        )
        self.create_incident_report(
            IncidentReport(
                severity=Severity.CRITICAL,
                status=Status.CLOSED,
                title="Food Poisoning Incident",
                description="Multiple guests reported food poisoning after dining at the hotel restaurant.",
                filing_person_Id=304,
                reviewer_Id=404
            )
        )

    def get_incident_reports(self) -> List[IncidentReport]:
        return self._incident_reports
    
    def get_incident_report_by_id(self, id:int) -> IncidentReport:
        for incident_report in self._incident_reports:
            if incident_report.id == id:
                return incident_report
            
        return None 
    
    def create_incident_report(self, incident_report: IncidentReport) -> IncidentReport:
        incident_report.set_id(self._nextId)
        self._nextId += 1

        self._incident_reports.append(incident_report)
        return incident_report

    def update_incident_report(self, id: int, updated_incident_report: IncidentReport) -> IncidentReport:
        incident_report = self.get_incident_report_by_id(id)

        if(incident_report != None):
            incident_report.update(updated_incident_report)
            return incident_report
        else:
            return ValueError

                
    def delete_incident_report(self, id:int) -> None:
        for incident_report in self._incident_reports:
            if incident_report.id == id:
                self._incident_reports.remove(incident_report)
                return
            
        return ValueError
            