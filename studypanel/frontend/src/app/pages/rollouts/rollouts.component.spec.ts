import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RolloutsComponent } from './rollouts.component';

describe('RolloutsComponent', () => {
  let component: RolloutsComponent;
  let fixture: ComponentFixture<RolloutsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RolloutsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(RolloutsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render rollouts heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Rollouts');
  });
});
