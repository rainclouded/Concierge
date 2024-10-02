import json
from incident_reports_server.models.models import IncidentReport, Severity, Status

class IncidentReportFactory:

    #converts json object to incident report object
    @staticmethod
    def create_incident_report(incident_report_JSON : json) -> IncidentReport:
        required_keys = ["severity","status","title","description","filing_person_id","reviewer_id"]
        
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
        
 
        
        
        
        
     