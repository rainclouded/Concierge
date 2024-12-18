import { CanActivateFn, Router } from '@angular/router';
import { ApiKeyService } from '../services/api-key.service';
import { inject } from '@angular/core';

export const authGuard: CanActivateFn = (route, state) => {
  const apiKeyService = inject(ApiKeyService);
  const router = inject(Router)
  if (apiKeyService.isKeyExpired()) {
    router.navigate(['/login']);
    return false;
  }
  return true;
};
