import { Component } from '@angular/core';
import { RouterOutlet, RouterLink, RouterLinkActive, Router } from '@angular/router';
import { SidebarComponent } from "../../components/sidebar/sidebar.component";
import { ApiKeyService } from '../../services/api-key.service';

@Component({
  selector: 'app-dashboard-page',
  standalone: true,
  imports: [RouterOutlet, RouterLink, RouterLinkActive, SidebarComponent],
  templateUrl: './dashboard-page.component.html',
})
export class DashboardPageComponent {
  constructor(private router: Router, private apiKeyService: ApiKeyService) {}
  handleLogout() {
    this.apiKeyService.endSession();
    this.router.navigate(['staff/login']);
  }
}
