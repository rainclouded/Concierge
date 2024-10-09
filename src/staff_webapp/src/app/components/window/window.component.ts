import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-window',
  standalone: true,
  imports: [],
  templateUrl: './window.component.html',
})
export class WindowComponent {
  @Input() isOpen = false;
  @Output() closeWindow = new EventEmitter();

  onCloseWindow() {
    this.closeWindow.emit(false);
  }
}
