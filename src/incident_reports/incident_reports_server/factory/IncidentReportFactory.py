import json
from incident_reports.incident_reports_server.model.Models import IncidentReport

class IncidentReportFactory:
    def __init__(self) -> None:
        pass

    #converts json object to incident report object
    def create_incident_report(incident_report_JSON : json) -> IncidentReport:
        pass 