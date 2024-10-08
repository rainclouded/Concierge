import { Routes } from '@angular/router';
import { DashboardPageComponent } from './pages/dashboard-page/dashboard-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { HomeTabComponent } from './pages/home-tab/home-tab.component';
import { AmenitiesTabComponent } from './pages/amenities-tab/amenities-tab.component';

export const routes: Routes = [
	{
		path: '',
		redirectTo: 'login',
		pathMatch: 'full'
	},
	{
		path: 'login',
		component: LoginPageComponent,
	},
	{
		path: 'dashboard',
		component: DashboardPageComponent,
		children: [
			{
				path: 'home',
				component: HomeTabComponent
			},
			{
				path: 'amenities',
				component: AmenitiesTabComponent
			},
			{ path: '',
				redirectTo: 'home',
				pathMatch: 'full' 
			}
		]
	}
];
