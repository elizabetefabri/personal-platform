import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActiveReleasesComponent } from './active-releases.component';

describe('ActiveReleasesComponent', () => {
  let component: ActiveReleasesComponent;
  let fixture: ComponentFixture<ActiveReleasesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ActiveReleasesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ActiveReleasesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render active releases heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Active Releases');
  });
});
