import { Injectable } from '@angular/core';
import { jwtDecode, JwtPayload } from 'jwt-decode';

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

  isKeyExpired(): boolean {
    const sessionKey = this.getSession();
    if (!sessionKey || sessionKey == "") {
      return true;
    }
    try {
      const decoded = jwtDecode(sessionKey);
      console.log(decoded)
      console.log(
        `Key expired: ${(decoded.exp ?? 0) < Math.floor(Date.now() / 1000)}`
      );
      return (decoded.exp ?? 0) < Math.floor(Date.now() / 1000);
    } catch (error) {
      console.log("Failed to decode session key")
      return true;
    }
  }

  getRemainingSessionTime(): bigint {
    //TODO : parse JWT token
    return 1n;
  }

}
