import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiGenPasswordResponse, ApiResponse } from '../models/apiresponse.model';
import { IAccount } from '../models/account.model';
import { BASE_API_URL } from '../constants/constants'

@Injectable({
  providedIn: 'root',
})
export class AccountService {
  apiUrl = `/accounts`;

  constructor(private http: HttpClient) {}

  // GET request - Retrieve all accounts
  getAllAccounts(): Observable<IAccount[]> {
    return this.http.get<IAccount[]>(`${this.apiUrl}`);
  }

  // GET request - Retrieve a single account by username
  getAccount(username: string): Observable<ApiResponse<IAccount>> {
    return this.http.get<ApiResponse<IAccount>>(`${this.apiUrl}/${username}`);
  }

  // POST request - Create a new account
  addAccount(account: IAccount): Observable<ApiResponse<IAccount>> {
    return this.http.post<ApiResponse<IAccount>>(`${this.apiUrl}`, account);
  }

  // POST request - Delete an account by username
  deleteAccount(username: string): Observable<ApiResponse<null>> {
    return this.http.post<ApiResponse<null>>(`${this.apiUrl}/delete`, {
      username,
    });
  }

  // PUT request - Update an existing account
  updateAccount(account: IAccount): Observable<ApiGenPasswordResponse> {
    return this.http.put<ApiGenPasswordResponse>(
      `${this.apiUrl}/update`,
      account
    );
  }
}
