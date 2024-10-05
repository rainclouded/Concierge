import { TestBed } from '@angular/core/testing';

import { AmenityService } from './amenity.service';

describe('AmenityService', () => {
  let service: AmenityService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AmenityService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
