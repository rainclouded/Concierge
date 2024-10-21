import { Injectable } from '@angular/core';

// See https://chintanonweb.medium.com/mastering-angular-session-security-a-deep-dive-into-sessionmanagementservice-ad24e71b86ed
@Injectable({
  providedIn: 'root'
})
export class ApiKeyService {
  private sessionKey = 'active_session';

  constructor() { }

  setSession(sessionData: any): void {
    localStorage.setItem(this.sessionKey, JSON.stringify(sessionData));
  }

  getSession(): any | null {
    const session = localStorage.getItem(this.sessionKey)
    return session ? JSON.parse(session) : null;
  }

  endSession(): void {
    localStorage.removeItem(this.sessionKey)
  }

  isSessionActive(): boolean {
    return this.getSession() != null;
  }

  isSessionKeyValid(): boolean {
    //TODO : validate JWT token
    return true;
  }

  getRemainingSessionTime(): bigint {
    //TODO : parse JWT token
    return 1n;
  }

}
