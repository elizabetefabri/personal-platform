import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RolloutServiceComponent } from './rollout-service.component';

describe('RolloutServiceComponent', () => {
  let component: RolloutServiceComponent;
  let fixture: ComponentFixture<RolloutServiceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RolloutServiceComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(RolloutServiceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render rollout service heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Rollout Service');
  });
});
