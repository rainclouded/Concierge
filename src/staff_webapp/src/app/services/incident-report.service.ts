import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { IIncidentReport } from '../models/incident-report.model';
import { Observable } from 'rxjs';
import { ApiResponse } from '../models/apiresponse.model';
import { BASE_API_URL } from '../constants/constants'

@Injectable({
  providedIn: 'root'
})
export class IncidentReportService {
  apiUrl = `${BASE_API_URL}/incident_reports/`;

  constructor(private http: HttpClient) { }

  // GET request
  getAllReports(): Observable<ApiResponse<IIncidentReport[]>> {
    return this.http.get<ApiResponse<IIncidentReport[]>>(`${this.apiUrl}`);
  }

  getReport(id: number): Observable<ApiResponse<IIncidentReport>> {
    return this.http.get<ApiResponse<IIncidentReport>>(`${this.apiUrl}${id}`);
  }

  // POST request
  addReport(amenity: IIncidentReport) {
    return this.http.post<ApiResponse<IIncidentReport>>(`${this.apiUrl}`, amenity);
  }

  // PUT request
  updateReport(id: number, amenity: IIncidentReport): Observable<ApiResponse<IIncidentReport>> {
    return this.http.put<ApiResponse<IIncidentReport>>(`${this.apiUrl}${id}`, amenity);
  }

  // DELETE request
  deleteReport(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}${id}`);
  }
}
