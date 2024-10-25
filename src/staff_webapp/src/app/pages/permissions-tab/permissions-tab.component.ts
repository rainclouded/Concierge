import { Component } from '@angular/core';
import { WindowComponent } from '../../components/window/window.component';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { PermissionService } from '../../services/permission.service';
import { IPermissionGroup } from '../../models/permission-group';
import { ApiResponse } from '../../models/apiresponse.model';
import { error } from 'cypress/types/jquery';
import { finalize } from 'rxjs/internal/operators/finalize';
import { IPermission } from '../../models/permission';

@Component({
  selector: 'app-permissions-tab',
  standalone: true,
  imports: [WindowComponent, MatTableModule, MatCheckboxModule],
  templateUrl: './permissions-tab.component.html',
})
export class PermissionsTabComponent {
  isOpenWIndow = false;
  permissionGroups: IPermissionGroup[] = [];
  flatPermissions: { [key: string]: string | boolean | number }[] = [];
  columnNames: string[] = ['cat', 'dog'];

  constructor(private permissionService: PermissionService) {
    this.getAllPermissionGroups();
  }

  getAllPermissionGroups() {
    this.permissionService.getAllPermissionGroups().subscribe({
      next: (response) => {
        this.permissionGroups = response.data;
        this.flatPermissions = this.getFlatPermissions();
        this.columnNames = Object.keys(this.flatPermissions[0]);
      },
    });
  }

  ngOnInit(): void {}

  openWindow() {
    this.isOpenWIndow = true;
  }

  closeWindow() {
    this.isOpenWIndow = false;
  }

  getFlatPermissions(): { [key: string]: string | boolean | number }[] {
    const permissions: IPermission[] = [];
    const flatPermissionGroups: { [key: string]: string | boolean | number }[] =
      [];

    this.permissionGroups.forEach((pGroup) => {
      pGroup.groupPermissions.forEach((permission) => {
        if (!permissions.find(p=>p.permissionId == permission.permissionId)) {
          permissions.push(permission)
        }
      });
    });

    Array.from(permissions, (permission) => {
      const permissionEntry: {
        id: number;
        [key: string]: string | boolean | number;
      } = {
        id: permission.permissionId,
        name: permission.permissionName,
      };
      this.permissionGroups.forEach((group) => {
        const groupPermission = group.groupPermissions.find(
          (p) => p.permissionName === permission.permissionName
        );
        permissionEntry[group.groupName] =
          groupPermission?.permissionState ?? false;
      });
      flatPermissionGroups.push(permissionEntry);
    });
    return flatPermissionGroups;
  }

  togglePermission(row: any, column: string, state: any) {
    console.log(row);
    console.log(column);

    const group = this.permissionGroups.find(
      (group) => group.groupName === column
    );
    const groupId = group?.groupId ?? -1;
    this.permissionService
      .updatePermissionGroup(groupId, row.id, state)
      .pipe(
        finalize(() => {
          this.getAllPermissionGroups();
        })
      )
      .subscribe({
        next: (v) => {},
        error: (err) => {},
      });
  }
}
