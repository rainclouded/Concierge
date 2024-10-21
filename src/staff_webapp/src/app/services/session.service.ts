import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ISesssionKey } from '../models/session-key.model';
import { ApiResponse } from '../models/apiresponse.model';
import { Observable } from 'rxjs/internal/Observable';
import { ILogin } from '../models/login.models';

@Injectable({
  providedIn: 'root',
})
export class SessionService {
  apiUrl = 'http://localhost:8089/sessions';

  constructor(private http: HttpClient) {}

  postSession(credentials: ILogin): Observable<ApiResponse<ISesssionKey>> {
    return this.http.post<ApiResponse<ISesssionKey>>(`${this.apiUrl}`, credentials)
  }
}
