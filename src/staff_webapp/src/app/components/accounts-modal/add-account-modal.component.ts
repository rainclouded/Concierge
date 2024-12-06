import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IAccount } from '../../models/account.model';

@Component({
  selector: 'add-account-modal',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './add-account-modal.component.html',
})
export class AddAccountModalComponent {
  @Output() close = new EventEmitter<void>();
  @Output() save = new EventEmitter<IAccount>(); // Notify parent with account data

  accountType: string = '';
  username: string = '';
  password: string = '';
  isPasswordEnabled: boolean = false;

  onAccountTypeChange(event: Event): void {
    const target = event.target as HTMLSelectElement | null; // Safely cast to HTMLSelectElement
    const type = target?.value || ''; // Access the value property safely
    this.accountType = type;
    this.isPasswordEnabled = type === 'staff';

    // Clear the password if the type changes to guest
    if (!this.isPasswordEnabled) {
      this.password = '';
    }
  }

  submitAccount(): void {
    if (!this.accountType || !this.username) {
      console.error('Type and Username are required.');
      return;
    }
    if (this.accountType === 'guest' && isNaN(Number(this.username))) {
      console.error('Guest usernames must be numeric.');
      return;
    }
    if (
      this.accountType === 'staff' &&
      (!/\d/.test(this.username) || !/[a-zA-Z]/.test(this.username))
    ) {
      console.error(
        'Staff usernames must contain at least one letter and one number.'
      );
      return;
    }

    const accountData: IAccount = {
      type: this.accountType,
      username: this.username,
      password: this.isPasswordEnabled ? this.password : undefined,
    };

    this.save.emit(accountData);
    this.close.emit();
  }
}
