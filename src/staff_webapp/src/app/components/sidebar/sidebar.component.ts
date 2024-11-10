import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, CommonModule],
  templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
  items = [
    {
      routerLink: 'home',
      label: 'Home',
      icon: 'fa-solid fa-house',
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
}
