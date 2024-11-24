import { Component, OnInit } from '@angular/core';
import { WindowComponent } from '../../components/window/window.component';
//import { AccountFormComponent } from "../../components/account-form/account-form.component";
import { IAccount } from '../../models/account.model';
import { AccountService } from '../../services/account.service';

@Component({
  selector: 'app-accounts-tab',
  standalone: true,
  imports: [WindowComponent], //AccountFormComponent
  templateUrl: './accounts-tab.component.html',
})
export class AccountsTabComponent implements OnInit {
  isOpenWindow = false;
  accounts: IAccount[] = [];
  account!: IAccount;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {
    this.getAllAccounts();
  }

  // Fetch all accounts from the service
  getAllAccounts() {
    this.accountService.getAllAccounts().subscribe({
      next: (response) => {
        if (response.data) {
          this.accounts = response.data;
        }
      },
      error: (err) => {
        console.error('Error fetching accounts:', err);
      },
    });
  }

  createAccount() {
    const newAccount = {
      username: '12345',
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

  // Load the selected account into the form for editing
  loadAccount(account: IAccount) {
    this.account = account;
    this.openWindow();
  }

  // New method to delete a hardcoded account
  deleteHardcodedAccount() {
    const hardcodedUsername = 'staff3';
    this.deleteAccount(hardcodedUsername);
  }

  // Delete an account by username
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

  updateAccount() {
    const updatedAccount: IAccount = {
      username: '12345',
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

  // Open the modal
  openWindow() {
    this.isOpenWindow = true;
  }

  // Close the modal and refresh the list
  closeWindow() {
    this.isOpenWindow = false;
    this.getAllAccounts();
  }
}
