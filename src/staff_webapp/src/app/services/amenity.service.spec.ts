import { TestBed } from '@angular/core/testing';
import { AmenityService } from './amenity.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AmenityService', () => {
  let service: AmenityService;

  beforeEach(() => {
    TestBed.configureTestingModule(
      {
        imports: [HttpClientTestingModule],
      }
    );
    service = TestBed.inject(AmenityService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
