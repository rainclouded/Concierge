// import { Component, OnInit } from '@angular/core';
// import { SessionService } from '../../services/session.service';

// @Component({
//   selector: 'app-home-tab',
//   standalone: true,
//   imports: [],
//   templateUrl: './home-tab.component.html',
// })
// export class HomeTabComponent implements OnInit {
//   accountName: string | null = null;

//   constructor(private sessionService: SessionService) {}

//   ngOnInit() {
//     this.sessionService.getSessionMe().subscribe(() => {
//       this.accountName = this.sessionService.accountName;
//     });
//   }
// }

import { Component, OnInit } from '@angular/core';
import { SessionService } from '../../services/session.service';
import { ISessionData } from '../../models/session-data.model';

@Component({
  selector: 'app-home-tab',
  standalone: true,
  imports: [],
  templateUrl: './home-tab.component.html',
})
export class HomeTabComponent implements OnInit {
  sessionData?: ISessionData;

  constructor(private sessionService: SessionService) {}

  ngOnInit(): void {
    this.getSessionDetails();
  }

  getSessionDetails(): void {
    this.sessionService.getSessionMe()
    .subscribe({
      next: (response) => {
        this.sessionData = response.data;
        console.log(this.sessionData);
      },
      error: (err) => {
        console.error('Error fetching session details:', err);
      }
    });
  }
}
