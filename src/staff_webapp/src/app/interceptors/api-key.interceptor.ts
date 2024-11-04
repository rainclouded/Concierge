import { HttpInterceptorFn } from '@angular/common/http';
import { ApiKeyService } from '../services/api-key.service';
import { inject } from '@angular/core';


//See https://angular.dev/guide/http/interceptors#dependency-injection-in-interceptors
export const apiKeyInterceptor: HttpInterceptorFn = (req, next) => {
  const apiKey = inject(ApiKeyService).getSession();
  const newReq = !!apiKey ? req.clone({headers: req.headers.append('X-Api-Key', apiKey),}) :  req;
  return next(newReq);
};
