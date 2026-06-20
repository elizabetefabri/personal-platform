import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PainelFinanceiroComponent } from './painel-financeiro.component';

describe('PainelFinanceiroComponent', () => {
  let component: PainelFinanceiroComponent;
  let fixture: ComponentFixture<PainelFinanceiroComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PainelFinanceiroComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PainelFinanceiroComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
