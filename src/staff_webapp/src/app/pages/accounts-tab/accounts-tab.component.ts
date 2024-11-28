import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IAccount } from '../../models/account.model';
import { AccountService } from '../../services/account.service';
import { AccountModalComponent } from '../../components/accounts-modal/account-modal.component';
import { AddAccountModalComponent } from '../../components/accounts-modal/add-account-modal.component';

@Component({
  selector: 'accounts-tab',
  standalone: true,
  imports: [CommonModule, FormsModule, AccountModalComponent, AddAccountModalComponent],
  templateUrl: './accounts-tab.component.html',
})
export class AccountsTabComponent implements OnInit {
  allAccounts: IAccount[] = []; // Master list of all accounts
  filteredAccounts: IAccount[] = []; // Current page of accounts
  searchTerm: string = '';
  currentPage: number = 1;
  pageSize: number = 10; // Number of accounts per page
  isAddAccountModalOpen: boolean = false;
  selectedAccount!: IAccount | null;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {
    this.fetchAccounts();
  }

  fetchAccounts(): void {
    this.accountService.getAllAccounts().subscribe({
      next: (response) => {
        this.allAccounts = response;
        this.updatePage(); // Show the first page
      },
      error: (err) => {
        console.error('Error fetching accounts:', err);
      },
    });
  }

  addNewAccount(account: IAccount): void {
    this.accountService.addAccount(account).subscribe({
      next: (response) => {
        console.log(response.message);
        this.fetchAccounts(); // Refresh the accounts list
      },
      error: (err) => {
        console.error('Error adding account:', err);
      },
    });
  }

  updatePage(): void {
    const startIndex = (this.currentPage - 1) * this.pageSize;
    const endIndex = this.currentPage * this.pageSize;
    this.filteredAccounts = this.allAccounts.slice(startIndex, endIndex);
  }

  searchAccounts(): void {
    const searchTerm = this.searchTerm.trim().toLowerCase();
    const filtered = this.allAccounts.filter((account) =>
      account.username.toLowerCase().includes(searchTerm)
    );
    this.currentPage = 1; // Reset to the first page
    this.filteredAccounts = filtered.slice(0, this.pageSize);
  }

  openAddAccountModal(): void {
    this.isAddAccountModalOpen = true;
  }

  closeAddAccountModal(): void {
    this.isAddAccountModalOpen = false;
  }

  openAccountModal(account: IAccount): void {
    this.selectedAccount = account;
  }

  closeAccountModal(): void {
    this.selectedAccount = null;
  }

  goToPreviousPage(): void {
    if (this.currentPage > 1) {
      this.currentPage--;
      this.updatePage();
    }
  }

  goToNextPage(): void {
    const maxPages = Math.ceil(this.allAccounts.length / this.pageSize);
    if (this.currentPage < maxPages) {
      this.currentPage++;
      this.updatePage();
    }
  }
}
