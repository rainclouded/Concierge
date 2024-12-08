import { Component } from '@angular/core';
import { ReactiveFormsModule , FormGroup, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { SessionService } from '../../services/session.service';
import { ApiKeyService } from '../../services/api-key.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login-page',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './login-page.component.html',
})
export class LoginPageComponent {
  loginForm: FormGroup;

  constructor(
    private router: Router,
    private sessionService: SessionService,
    private apiKeyService: ApiKeyService,
    private toastr: ToastrService
  ) {
    this.loginForm = new FormGroup({
      username: new FormControl(''),
      password: new FormControl(''),
    });
  }

  handleLogin() {
    console.log(this.loginForm.value);
    this.sessionService.postSession(this.loginForm.value).subscribe({
      next: (response: any) => {
        console.log(response)
        console.log(`Session key: ${response.data.sessionKey}`)
        this.apiKeyService.setSession(response.data.sessionKey)
        this.router.navigate(['/dashboard']);
      },
      error: (response: any) => {
        console.log(response);
        console.log(response.status);
        this.toastr.error(`${response.error.message}`, 'Sign In Failed');
      }
    });
  }
}
