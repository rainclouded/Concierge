import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeTabComponent } from './home-tab.component';

describe('HomeTabComponent', () => {
  let component: HomeTabComponent;
  let fixture: ComponentFixture<HomeTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HomeTabComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(HomeTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
