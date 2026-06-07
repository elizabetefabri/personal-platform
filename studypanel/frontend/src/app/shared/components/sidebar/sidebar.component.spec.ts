import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { SidebarComponent } from './sidebar.component';

describe('SidebarComponent', () => {
  let component: SidebarComponent;
  let fixture: ComponentFixture<SidebarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SidebarComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(SidebarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have 4 navigation items', () => {
    expect(component.items.length).toBe(4);
  });

  it('should default to expanded', () => {
    expect(component.expanded).toBeTrue();
  });

  it('should default activeItem to dashboard', () => {
    expect(component.activeItem).toBe('dashboard');
  });

  it('should return expanded sidebar width when expanded', () => {
    component.expanded = true;
    expect(component.sidebarWidth).toBe('var(--sidebar-width-expanded)');
  });

  it('should return collapsed sidebar width when collapsed', () => {
    component.expanded = false;
    expect(component.sidebarWidth).toBe('var(--sidebar-width-collapsed)');
  });

  it('should emit toggleExpanded when button is clicked', () => {
    const spy = spyOn(component.toggleExpanded, 'emit');
    const button = fixture.nativeElement.querySelector('.toggleButton') as HTMLButtonElement;
    button.click();
    expect(spy).toHaveBeenCalled();
  });
});
