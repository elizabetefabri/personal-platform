import { ComponentFixture, TestBed } from '@angular/core/testing';
import { PerformanceEngineeringComponent } from './performance-engineering.component';

describe('PerformanceEngineeringComponent', () => {
  let component: PerformanceEngineeringComponent;
  let fixture: ComponentFixture<PerformanceEngineeringComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PerformanceEngineeringComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PerformanceEngineeringComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render performance engineering heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Performance Engineering');
  });
});
