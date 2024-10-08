import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AmenitiesTabComponent } from './amenities-tab.component';

describe('AmenitiesTabComponent', () => {
  let component: AmenitiesTabComponent;
  let fixture: ComponentFixture<AmenitiesTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AmenitiesTabComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AmenitiesTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
