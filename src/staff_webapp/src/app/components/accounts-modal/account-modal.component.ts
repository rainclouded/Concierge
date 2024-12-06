import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IAccount } from '../../models/account.model';
import { IPermissionGroup } from '../../models/permission-group';
import { PermissionService } from '../../services/permission.service';
import { AccountService } from '../../services/account.service';
import { ConfirmDeletionModalComponent } from './confirm-deletion-modal.component';

@Component({
  selector: 'account-modal',
  standalone: true,
  imports: [CommonModule, FormsModule, ConfirmDeletionModalComponent],
  templateUrl: './account-modal.component.html',
})
export class AccountModalComponent implements OnInit {
  @Input() account!: IAccount; // Account details passed from parent
  @Output() close = new EventEmitter<void>();
  @Output() accountDeleted = new EventEmitter<void>();

  allPermissions: IPermissionGroup[] = []; // All available permission groups
  accountPermissions: number[] = []; // Current permission group IDs for the account
  modifiedPermissions: number[] = []; // Selected permissions (modified locally)
  isPasswordEnabled: boolean = false;
  isShowingPassword: boolean = false;
  copiedToClipboard: boolean = false;
  newPassword: string = '';

  isDeleteModalOpen: boolean = false;

  constructor(
    private permissionService: PermissionService,
    private accountService: AccountService
  ) {}

  ngOnInit(): void {
    this.isShowingPassword = false;
    this.isPasswordEnabled = this.account.type === 'staff';
    this.loadPermissions();
  }

  loadPermissions(): void {
    // Fetch all permission groups
    this.permissionService.getAllPermissionGroups().subscribe({
      next: (response) => {
        this.allPermissions = response.data;
        if (this.account?.id !== undefined) {
          // Fetch account-specific permissions
          this.permissionService
            .getPermissionsForAccount(+this.account.id)
            .subscribe({
              next: (accountPermissionsResponse) => {
                this.accountPermissions = accountPermissionsResponse.data.map(
                  (group) => group.groupId
                );
                this.modifiedPermissions = [...this.accountPermissions]; // Clone for editing
              },
              error: (err) =>
                console.error('Error fetching account permissions:', err),
            });
        }
      },
      error: (err) => console.error('Error fetching permissions:', err),
    });
  }

  togglePermission(groupId: number): void {
    if (this.modifiedPermissions.includes(groupId)) {
      this.modifiedPermissions = this.modifiedPermissions.filter(
        (id) => id !== groupId
      );
    } else {
      this.modifiedPermissions.push(groupId);
    }
  }

  hasPermissionChanges(): boolean {
    return (
      JSON.stringify(this.modifiedPermissions) !==
      JSON.stringify(this.accountPermissions)
    );
  }

  saveChanges(): void {
    const addGroups = this.modifiedPermissions.filter(
      (id) => !this.accountPermissions.includes(id)
    );
    const removeGroups = this.accountPermissions.filter(
      (id) => !this.modifiedPermissions.includes(id)
    );

    // Sequentially update permission groups
    const updates = [
      ...removeGroups.map((groupId) =>
        this.permissionService.updatePermissionGroupMembers(
          groupId,
          [],
          [+this.account.id!]
        )
      ),
      ...addGroups.map((groupId) =>
        this.permissionService.updatePermissionGroupMembers(
          groupId,
          [+this.account.id!],
          []
        )
      ),
    ];

    Promise.allSettled(updates).then((results) => {
      console.log('Permission updates:', results);
      this.close.emit(); // Close the modal after saving
    });
  }

  updatePassword(): void {
    const updatedAccount = {
      ...this.account,
      password: this.newPassword || undefined,
    };

    this.accountService.updateAccount(updatedAccount).subscribe({
      next: (response) => {
        if (this.account.type === 'guest' && response.message) {
          const match = response.message.match(/password:\s*(\d+)/);
          if (match) {
            console.log("I did a tthing2"+match[1])
            this.newPassword = ""+match[1];
            this.isShowingPassword = true;
          }
        }
      },
      error: (err) => console.error('Error updating account:', err),
    });
  }

  copyToClipboard(text: string): void {
    this.copiedToClipboard = true;
    navigator.clipboard.writeText(text).then(() => {
      console.log('Password copied to clipboard');
    }).catch((err) => {
      console.error('Error copying password to clipboard:', err);
    });
  }
  
  refreshGenPass(): void {
    this.copiedToClipboard = false;
    this.isShowingPassword = false;
  }
  deleteAccount(): void {
    this.accountService.deleteAccount(this.account.username).subscribe({
      next: () => {
        console.log(`Account ${this.account.username} deleted successfully.`);
        this.accountDeleted.emit(); // Notify parent
        this.close.emit(); // Close the account modal
      },
      error: (err) => {
        console.error('Error deleting account:', err);
      },
    });
  }

  openDeleteModal(): void {
    this.isDeleteModalOpen = true;
  }

  closeDeleteModal(): void {
    this.isDeleteModalOpen = false;
  }
}
