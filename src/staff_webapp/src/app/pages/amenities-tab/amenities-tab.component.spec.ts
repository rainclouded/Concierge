import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AmenitiesTabComponent } from './amenities-tab.component';
import { WindowComponent } from '../../components/window/window.component';
import { AmenityFormComponent } from '../../components/amenity-form/amenity-form.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AmenitiesTabComponent', () => {
  let component: AmenitiesTabComponent;
  let fixture: ComponentFixture<AmenitiesTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WindowComponent, AmenityFormComponent, HttpClientTestingModule],
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
