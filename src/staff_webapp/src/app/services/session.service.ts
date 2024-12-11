import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ISesssionKey } from '../models/session-key.model';
import { ApiResponse } from '../models/apiresponse.model';
import { Observable } from 'rxjs/internal/Observable';
import { ILogin } from '../models/login.models';
import { of, tap } from 'rxjs';
import { ISessionData } from '../models/session-data.model';

@Injectable({
  providedIn: 'root',
})
export class SessionService {
  apiUrl = `/sessions`;
  private sessionCache: ISessionData | null = null;

  constructor(private http: HttpClient) {}

  postSession(credentials: ILogin): Observable<ApiResponse<ISesssionKey>> {
    return this.http.post<ApiResponse<ISesssionKey>>(`${this.apiUrl}`, credentials)
  }

  getSessionMe(): Observable<ApiResponse<ISessionData>> {
    if (this.sessionCache) {
      return of({
        data: this.sessionCache,
        timestamp: new Date().toISOString(),
      } as ApiResponse<ISessionData>);
    }
    return this.http.get<ApiResponse<ISessionData>>(`${this.apiUrl}/me`).pipe(
      tap((response) => {
        if (response.data) {
          this.sessionCache = response.data;
        }
      })
    );
  }

  // Getter for sessionPermissionList
  get sessionPermissionList(): string[] | null {
    return this.sessionCache?.sessionData.SessionPermissionList || null;
  }

  // Getter for accountName
  get accountName(): string | null {
    return this.sessionCache?.sessionData.accountName || null;
  }
}
