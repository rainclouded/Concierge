import { Component, OnInit } from '@angular/core';
import { Event, NavigationEnd, Router, RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { SessionService } from '../../services/session.service';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, CommonModule],
  templateUrl: './sidebar.component.html',
})
export class SidebarComponent implements OnInit {
  currentRoute: string = '';
  accountName: string | null = null;

  constructor(public router: Router, private sessionService: SessionService) {
    // Subscribe to router events to update current route
    this.router.events.subscribe((event: Event) => {
      if (event instanceof NavigationEnd) {
        this.currentRoute = event.urlAfterRedirects.split('/').pop() || '';
      }
    });
  }

  ngOnInit() {
    this.sessionService.getSessionMe().subscribe(() => {
      this.accountName = this.sessionService.accountName;
    });
  }

  items = [
    {
      routerLink: 'home',
      label: 'Home',
      icon: 'fa-solid fa-house',
    },
    {
      routerLink: 'accounts',
      label: 'Accounts',
      icon: 'fa-solid fa-users',
    },
    {
      routerLink: 'amenities',
      label: 'Amenities',
      icon: 'fa-solid fa-hotel',
    },
    {
      routerLink: 'incident_reports',
      label: 'Incident Reports',
      icon: 'fa-solid fa-file-signature',
    },
    {
      routerLink: 'tasks',
      label: 'Tasks',
      icon: 'fa-solid fa-tasks',
    },
    {
      routerLink: 'permissions',
      label: 'Permissions',
      icon: 'fa-solid fa-user-shield',
    },
  ];

  isRouteActive(route: string): boolean {
    return this.currentRoute === route;
  }
}
