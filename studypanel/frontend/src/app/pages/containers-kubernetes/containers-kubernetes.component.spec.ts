import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ContainersKubernetesComponent } from './containers-kubernetes.component';

describe('ContainersKubernetesComponent', () => {
  let component: ContainersKubernetesComponent;
  let fixture: ComponentFixture<ContainersKubernetesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ContainersKubernetesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ContainersKubernetesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render backend heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Backend');
  });
});
