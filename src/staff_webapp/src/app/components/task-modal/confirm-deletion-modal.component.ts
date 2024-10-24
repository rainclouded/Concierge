import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-confirm-deletion-modal',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './confirm-deletion-modal.component.html',
})
export class ConfirmDeletionModalComponent {
  @Output() confirm = new EventEmitter<void>();  // Confirm event
  @Output() cancel = new EventEmitter<void>();   // Cancel event

  confirmDelete() {
    this.confirm.emit();  // Emit confirm event when delete is confirmed
  }

  cancelDelete() {
    this.cancel.emit();  // Emit cancel event when deletion is canceled
  }
}
