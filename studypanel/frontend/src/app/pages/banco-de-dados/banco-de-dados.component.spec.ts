import { ComponentFixture, TestBed } from '@angular/core/testing';
import { BancoDeDadosComponent } from './banco-de-dados.component';

describe('BancoDeDadosComponent', () => {
  let component: BancoDeDadosComponent;
  let fixture: ComponentFixture<BancoDeDadosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BancoDeDadosComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(BancoDeDadosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render banco-de-dados heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Banco de Dados');
  });
});
