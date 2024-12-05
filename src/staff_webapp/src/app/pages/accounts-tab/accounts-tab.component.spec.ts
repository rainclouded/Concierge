import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountsTabComponent } from './accounts-tab.component';

describe('AccountsTabComponent', () => {
  let component: AccountsTabComponent;
  let fixture: ComponentFixture<AccountsTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AccountsTabComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AccountsTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
