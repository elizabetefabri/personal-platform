import { ComponentFixture, TestBed } from '@angular/core/testing';
import { InteligenciaArtificialComponent } from './inteligencia-artificial.component';

describe('InteligenciaArtificialComponent', () => {
  let component: InteligenciaArtificialComponent;
  let fixture: ComponentFixture<InteligenciaArtificialComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InteligenciaArtificialComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(InteligenciaArtificialComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render inteligencia-artificial heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Inteligência Artificial');
  });
});
