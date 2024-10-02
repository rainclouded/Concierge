import unittest
from flask import json
from incident_reports_server.application.services import Services
from incident_reports_server.controllers.incident_reports_controller import create_app
from incident_reports_server.models.models import Severity, Status, IncidentReport

class incident_report_unit_test(unittest.TestCase):
    valid_incident_report = IncidentReport(
            severity=Severity.HIGH,  
            status=Status.OPEN,       
            title="Fire Alarm Malfunction",
            description="The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
            filing_person_id=303,    
            reviewer_id=403          
        )
    
    def setUp(self):
        Services.clear()
        self._incident_report_persistence = Services.get_incident_report_persistence()

        self.app = create_app(self._incident_report_persistence).test_client()
        self.app.testing = True

    def test_get_incident_reports(self):
        response = self.app.get('/incident_reports')
        
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.json['data'])
        
    def test_get_incident_report_by_id_valid_id_should_not_return_none(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.json['data'])
        self.assertDictEqual(response.json['data'], incident_report.to_dict())
        pass

    def test_get_incident_report_by_id_invalid_id_returns_not_found(self):
        response = self.app.get("/incident_reports/-1")
        self.assertEqual(response.status_code, 404)

    def test_add_incident_report_valid_incident_report_is_successful(self):
        response = self.app.post("/incident_reports", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        self.assertEqual(response.status_code, 201) 

    def test_add_incident_report_valid_incident_report_able_to_be_fetched(self):
        response = self.app.post("/incident_reports", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        
        incident_report = response.data
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")

    def test_add_incident_report_invalid_incident_report_fails(self):
        pass

    def test_add_incident_report_invalid_incident_report_not_fetchable(self):
        pass

    def test_add_incident_report_duplicate_incident_report_fails(self):
        pass

    def test_add_incident_report_null_incident_report_fails(self):
        pass

    def test_update_incident_report_valid_incident_report_is_successful(self):
        pass

    def test_update_incident_report_valid_incident_report_updated_incident_report_fetched(self):
        pass

    def test_update_incident_report_invalid_incident_report_fails(self):
        pass

    def test_update_incident_report_invalid_incident_report_not_fetchable(self):
        pass

    def test_update_incident_report_non_existing_incident_report_fails(self):
        pass

    def test_update_incident_report_null_incident_report_fails(self):
        pass

    def test_delete_incident_report_valid_incident_report_is_successful(self):
        pass

    def test_delete_incident_report_invalid_incident_report_fails(self):
        pass

    def test_delete_incident_report_invalid_incident_report_not_fetchable(self):
        pass

if __name__ == '__main__':
    unittest.main()