import { CanActivateFn, Router } from '@angular/router';
import { ApiKeyService } from '../services/api-key.service';
import { inject } from '@angular/core';

export const loginGuard: CanActivateFn = (route, state) => {
  const apiKeyService = inject(ApiKeyService);
  const router = inject(Router);
  if (!apiKeyService.isKeyExpired()) {
    router.navigate(['/dashboard'])
    return false;
  }

  return true
};
