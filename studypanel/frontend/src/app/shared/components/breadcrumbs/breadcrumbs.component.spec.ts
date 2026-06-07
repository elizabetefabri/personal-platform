import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterModule } from '@angular/router';
import { BreadcrumbsComponent, BreadcrumbItem } from './breadcrumbs.component';

describe('BreadcrumbsComponent', () => {
  let component: BreadcrumbsComponent;
  let fixture: ComponentFixture<BreadcrumbsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BreadcrumbsComponent, RouterModule.forRoot([])],
    }).compileComponents();

    fixture = TestBed.createComponent(BreadcrumbsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should not render nav when items is empty', () => {
    component.items = [];
    fixture.detectChanges();
    const nav = fixture.nativeElement.querySelector('nav');
    expect(nav).toBeNull();
  });

  it('should render nav when items are provided', () => {
    const items: BreadcrumbItem[] = [{ label: 'Home', href: '/' }, { label: 'Dashboard' }];
    component.items = items;
    fixture.detectChanges();
    const nav = fixture.nativeElement.querySelector('nav');
    expect(nav).toBeTruthy();
  });

  it('should render separator icon between items', () => {
    component.items = [{ label: 'Home', href: '/' }, { label: 'Dashboard' }];
    fixture.detectChanges();
    const separators = fixture.nativeElement.querySelectorAll('.separatorIcon');
    expect(separators.length).toBe(1);
  });

  it('should mark last item with aria-current="page"', () => {
    component.items = [{ label: 'Home', href: '/' }, { label: 'Dashboard' }];
    fixture.detectChanges();
    const current = fixture.nativeElement.querySelector('[aria-current="page"]');
    expect(current?.textContent?.trim()).toBe('Dashboard');
  });
});
