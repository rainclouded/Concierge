import { Injectable } from '@angular/core';
import { IPermissionGroup } from '../models/permission-group';
import { ApiResponse } from '../models/apiresponse.model';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class PermissionService {
  apiUrl = 'http://localhost:8089/permission-groups';

  constructor(private http: HttpClient) {}

  getAllPermissionGroups() {
    return this.http.get<ApiResponse<IPermissionGroup[]>>(`${this.apiUrl}`);
  }

  getPermissionsForAccount(accountId: number) {
    return this.http.get<ApiResponse<IPermissionGroup[]>>(
      `${this.apiUrl}?account-id=${accountId}`
    );
  }

  updatePermissionGroup(groupId: number, permissionId: number, state: boolean) {
    console.log(`Group ${groupId} permission ${permissionId} state ${state}`);
    return this.http.patch<ApiResponse<IPermissionGroup[]>>(
      `${this.apiUrl}/${groupId}`,
      {
        templateId: groupId,
        groupPermissions: [{ permissionId: permissionId, state: state }],
      }
    );
  }

  updatePermissionGroupMembers(
    groupId: number,
    addAccounts: number[],
    removeAccounts: number[]
  ) {
    return this.http.patch<ApiResponse<IPermissionGroup[]>>(
      `${this.apiUrl}/${groupId}`,
      {
        'add-accounts': addAccounts,
        'remove-accounts': removeAccounts,
      }
    );
  }
}
