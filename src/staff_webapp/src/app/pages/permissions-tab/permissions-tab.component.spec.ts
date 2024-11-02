import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WindowComponent } from '../../components/window/window.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { PermissionsTabComponent } from './permissions-tab.component';

describe('permissionsTabComponent', () => {
  let component: PermissionsTabComponent;
  let fixture: ComponentFixture<PermissionsTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WindowComponent, HttpClientTestingModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PermissionsTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
