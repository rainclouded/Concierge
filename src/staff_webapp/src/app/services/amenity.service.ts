import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiResponse } from '../models/apiresponse.model';
import { IAmenity } from '../models/amenity.model';

@Injectable({
  providedIn: 'root'
})
export class AmenityService {
  apiUrl = 'http://localhost:8089/amenities/'

  constructor(private http: HttpClient) { }

  // GET request
  getAllAmenities(): Observable<ApiResponse<IAmenity[]>> {
    return this.http.get<ApiResponse<IAmenity[]>>(`${this.apiUrl}`);
  }

  getAmenity(id: number): Observable<ApiResponse<IAmenity>> {
    return this.http.get<ApiResponse<IAmenity>>(`${this.apiUrl}${id}`);
  }

  // POST request
  addAmenity(amenity: IAmenity) {
    return this.http.post<ApiResponse<IAmenity>>(`${this.apiUrl}`, amenity);
  }

  // PUT request
  updateAmenity(id: number, amenity: IAmenity): Observable<ApiResponse<IAmenity>> {
    return this.http.put<ApiResponse<IAmenity>>(`${this.apiUrl}${id}`, amenity);
  }

  // DELETE request
  deleteAmenity(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}${id}`);
  }
}
