import { Component } from '@angular/core';
import { ReactiveFormsModule , FormGroup, FormControl } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-page',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './login-page.component.html',
  styleUrl: './login-page.component.css'
})
export class LoginPageComponent {
  loginForm: FormGroup;

  constructor(private router: Router) {
    this.loginForm = new FormGroup({
      name: new FormControl(''),
      password: new FormControl('')
    })
  }

  handleLogin() {
    console.log(this.loginForm.value);
    this.router.navigate(['/dashboard']);
  }
}
