import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'confirm-deletion-modal',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './confirm-deletion-modal.component.html',
})
export class ConfirmDeletionModalComponent {
  @Input() accountUsername: string = ''; // Account to be deleted
  @Output() confirm = new EventEmitter<void>(); // Emits when deletion is confirmed
  @Output() cancel = new EventEmitter<void>(); // Emits when deletion is canceled
}
