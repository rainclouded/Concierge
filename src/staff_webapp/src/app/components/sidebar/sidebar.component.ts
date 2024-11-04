import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule],
  templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
  items = [
    {
      routerLink: 'home',
      label: 'Home',
    },
    {
      routerLink: 'amenities',
      label: 'Amenities',
    },
    {
      routerLink: 'incident_reports',
      label: 'Incident Reports',
    },
    {routerLink: 'tasks',
      label: 'Tasks'
    },
    {
      routerLink: 'permissions',
      label: 'Permissions',
    },
  ];
}
