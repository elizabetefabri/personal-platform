import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DevOpsComponent } from './devops.component';

describe('DevOpsComponent', () => {
  let component: DevOpsComponent;
  let fixture: ComponentFixture<DevOpsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DevOpsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DevOpsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render frontend heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Frontend');
  });
});
