import { Routes } from '@angular/router';
import { DashboardPageComponent } from './pages/dashboard-page/dashboard-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { HomeTabComponent } from './pages/home-tab/home-tab.component';
import { AmenitiesTabComponent } from './pages/amenities-tab/amenities-tab.component';
import { IncidentReportsTabComponent } from './pages/incident-reports-tab/incident-reports-tab.component';
import { authGuard } from './guards/auth.guard';
import { loginGuard } from './guards/login.guard';


export const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full',
  },
  {
    path: 'login',
    component: LoginPageComponent,
    canActivate: [loginGuard],
  },
  {
    path: 'dashboard',
    component: DashboardPageComponent,
    canActivate: [authGuard],
    children: [
      {
        path: 'home',
        component: HomeTabComponent,
        canActivate: [authGuard],
      },
      {
        path: 'amenities',
        component: AmenitiesTabComponent,
        canActivate: [authGuard],
      },
      {
        path: 'incident_reports',
        component: IncidentReportsTabComponent,
        canActivate: [authGuard],
      },
      { path: '', redirectTo: 'home', pathMatch: 'full' },
    ],
  },
  { path: '**', redirectTo: '/login' },
];
