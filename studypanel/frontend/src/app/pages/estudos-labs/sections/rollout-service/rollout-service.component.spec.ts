import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { RolloutServiceComponent } from './rollout-service.component';

describe('RolloutServiceComponent', () => {
  let component: RolloutServiceComponent;
  let fixture: ComponentFixture<RolloutServiceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RolloutServiceComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(RolloutServiceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render rollout service heading', () => {
    const h1 = fixture.nativeElement.querySelector('h1');
    expect(h1?.textContent).toContain('Rollout Service');
  });
});
