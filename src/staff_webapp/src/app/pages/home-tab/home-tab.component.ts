import { Component, OnInit } from '@angular/core';
import { SessionService } from '../../services/session.service';

@Component({
  selector: 'app-home-tab',
  standalone: true,
  imports: [],
  templateUrl: './home-tab.component.html',
})
export class HomeTabComponent implements OnInit {
  accountName: string | null = null;

  constructor(private sessionService: SessionService) {}

  ngOnInit() {
    this.sessionService.getSessionMe().subscribe(() => {
      this.accountName = this.sessionService.accountName;
    });
  }
}
