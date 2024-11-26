import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IAccount } from '../../models/account.model';
import { EditAccountModalComponent } from '../../components/accounts-modal/edit-account-modal.component';
import { DeleteConfirmationModalComponent } from '../../components/accounts-modal/delete-confirmation-modal.component';
import { AccountService } from '../../services/account.service';
import { PermissionService } from '../../services/permission.service';
import { mockUsers } from './mock-users';

@Component({
  selector: 'app-accounts-tab',
  standalone: true,
  imports: [
    CommonModule,
    EditAccountModalComponent,
    DeleteConfirmationModalComponent,
  ],
  templateUrl: './accounts-tab.component.html',
})
export class AccountsTabComponent implements OnInit {
  accounts: IAccount[] = [...mockUsers];
  paginatedAccounts: IAccount[] = [];
  selectedAccount!: IAccount;
  usernameToDelete!: string;

  constructor(
    private accountService: AccountService,
    private permissionService: PermissionService
  ) {}

  // Pagination variables
  currentPage = 1;
  itemsPerPage = 5;
  totalPages = Math.ceil(this.accounts.length / this.itemsPerPage);

  // Modal states
  isEditModalOpen = false;
  isDeleteModalOpen = false;

  ngOnInit(): void {
    this.updatePagination();
  }

  // Pagination logic
  updatePagination() {
    const startIndex = (this.currentPage - 1) * this.itemsPerPage;
    this.paginatedAccounts = this.accounts.slice(
      startIndex,
      startIndex + this.itemsPerPage
    );
  }

  goToPreviousPage() {
    if (this.currentPage > 1) {
      this.currentPage--;
      this.updatePagination();
    }
  }

  goToNextPage() {
    if (this.currentPage < this.totalPages) {
      this.currentPage++;
      this.updatePagination();
    }
  }

  // Edit modal logic
  openEditModal(account: IAccount) {
    this.selectedAccount = account;
    this.isEditModalOpen = true;
  }

  closeEditModal() {
    this.isEditModalOpen = false;
  }

  saveEditedAccount(updatedAccount: IAccount) {
    const index = this.accounts.findIndex(
      (acc) => acc.username === updatedAccount.username
    );
    if (index !== -1) {
      this.accounts[index] = updatedAccount;
      this.updatePagination();
      console.log('Account updated:', updatedAccount);
    }
    this.closeEditModal();
  }

  // Delete modal logic
  openDeleteModal(username: string) {
    this.usernameToDelete = username;
    this.isDeleteModalOpen = true;
  }

  closeDeleteModal() {
    this.isDeleteModalOpen = false;
  }

  confirmDeleteAccount(username: string) {
    this.accounts = this.accounts.filter((acc) => acc.username !== username);
    this.totalPages = Math.ceil(this.accounts.length / this.itemsPerPage);
    this.updatePagination();
    console.log('Account deleted:', username);
    this.closeDeleteModal();
  }

  getAllAccounts() {
    this.accountService.getAllAccounts().subscribe({
      next: (response) => {
        if (response.data) {
          this.accounts = response.data;
          console.log('accounts:');
          console.log(this.accounts);
        }
      },
      error: (err) => {
        console.error('Error fetching accounts:', err);
      },
    });
  }

  updateAccount() {
    const updatedAccount: IAccount = {
      username: '12368',
      type: 'guest',
      //password: 'password3122',
    };

    this.accountService.updateAccount(updatedAccount).subscribe({
      next: (response) => {
        console.log('Account updated successfully:', response.message);
      },
      error: (err) => {
        console.error('Error updating account:', err);
      },
    });
  }

  createAccount() {
    const newAccount = {
      username: '12368',
      type: 'guest',
      //password: 'password123',
    };

    this.accountService.addAccount(newAccount).subscribe({
      next: (response) => {
        console.log(response.message);
      },
      error: (err) => {
        console.error('Error creating account:', err);
      },
    });
  }

  deleteHardcodedAccount() {
    const hardcodedUsername = '12368';
    this.deleteAccount(hardcodedUsername);
  }

  deleteAccount(username: string) {
    this.accountService.deleteAccount(username).subscribe({
      next: (response) => {
        console.log(response.message);
        this.getAllAccounts();
      },
      error: (err) => {
        console.error('Error deleting account:', err);
      },
    });
  }

  getAllPermissionGroups() {
    this.permissionService.getAllPermissionGroups().subscribe({
      next: (response) => {
        console.log(response.data);
      },
    });
  }
}
