import datetime
from incident_reports.incident_reports_server.model.Models import IncidentReport, Severity, Status
from incident_reports.incident_reports_server.persistence import IIncidentReportPersistence

class IncidentReportPersistenceStub(IIncidentReportPersistence):
    def __init__(self):
        self._incident_reports = [
            IncidentReport(
                ID=1,
                severity=Severity.LOW,
                status=Status.OPEN,
                title="Room Maintenance Request",
                description="Guest reported a leaky faucet in Room 203.",
                created_at=datetime.datetime(2024, 9, 1, 10, 30),
                filing_person_ID=301,  # Filing person could be a staff member
                reviewer_ID=401  # Reviewer could be a maintenance supervisor
            ),
            IncidentReport(
                ID=2,
                severity=Severity.MEDIUM,
                status=Status.IN_PROGRESS,
                title="Lost Property",
                description="A guest has reported a lost wallet in the lobby area.",
                created_at=datetime.datetime(2024, 9, 10, 14, 15),
                filing_person_ID=302,
                reviewer_ID=402
            ),
            IncidentReport(
                ID=3,
                severity=Severity.HIGH,
                status=Status.RESOLVED,
                title="Fire Alarm Malfunction",
                description="The fire alarm went off during dinner service; it was a false alarm.",
                created_at=datetime.datetime(2024, 9, 15, 9, 0),
                filing_person_ID=303,
                reviewer_ID=403
            ),
            IncidentReport(
                ID=4,
                severity=Severity.CRITICAL,
                status=Status.CLOSED,
                title="Food Poisoning Incident",
                description="Multiple guests reported food poisoning after dining at the hotel restaurant.",
                created_at=datetime.datetime(2024, 9, 20, 18, 45),
                filing_person_ID=304,
                reviewer_ID=404
            )
    ]

    def get_incident_reports(self):
        return self._incident_reports
    
    def get_incident_report_by_id(self, id:int):
        for incident_report in self._incident_reports:
            if incident_report.ID == id:
                return incident_report
            
        return None 
    
    def create_incident_report(self, incident_report: IncidentReport):
        self._incident_reports.insert(incident_report)
    
    def update_incident_report(self, id: int, updated_incident_report: IncidentReport):
        incident_report = self.get_incident_report_by_id(id)

        if(incident_report != None):
            incident_report.update(None)
        else:
            return ValueError

                
    def delete_incident_report(self, id:int):
        for incident_report in self._incident_reports:
            if incident_report.ID == id:
                self._incident_reports.remove(incident_report)
                return
            
        return ValueError
            