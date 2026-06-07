import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CloudComponent } from './cloud.component';

describe('CloudComponent', () => {
  let component: CloudComponent;
  let fixture: ComponentFixture<CloudComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CloudComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(CloudComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render cloud heading', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Cloud');
  });
});
