import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IAccount } from '../../models/account.model';

@Component({
  selector: 'app-edit-account-modal',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './edit-account-modal.component.html',
})
export class EditAccountModalComponent {
  @Input() account!: IAccount; // Input property to receive the account details
  @Output() close = new EventEmitter<void>(); // Event to close the modal
  @Output() save = new EventEmitter<IAccount>(); // Event to save changes

  newPassword: string = ''; // To store the new password

  generateNewPassword() {
    if (this.account.type === 'guest') {
      // Simulate new password generation
      this.newPassword = Math.random().toString(36).slice(-8);
    }
  }

  saveChanges() {
    if (this.account.type === 'staff' && this.newPassword) {
      // Update the password for staff
      this.account.password = this.newPassword;
    } else if (this.account.type === 'guest') {
      // Update the password for guest
      this.account.password = this.newPassword;
    }
    // Emit the updated account and close the modal
    this.save.emit(this.account);
    this.close.emit();
  }
}
