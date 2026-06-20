import { ComponentFixture, TestBed } from '@angular/core/testing';
import { signal } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BreadcrumbsComponent, BreadcrumbItem } from './breadcrumbs.component';
import { BreadcrumbService } from '../../../core/services/breadcrumb.service';

describe('BreadcrumbsComponent', () => {
  let component: BreadcrumbsComponent;
  let fixture: ComponentFixture<BreadcrumbsComponent>;
  let hiddenSignal: ReturnType<typeof signal<boolean>>;
  let extraSignal: ReturnType<typeof signal<BreadcrumbItem | null>>;

  beforeEach(async () => {
    hiddenSignal = signal(false);
    extraSignal = signal<BreadcrumbItem | null>(null);

    const breadcrumbServiceMock = {
      extra: extraSignal,
      hidden: hiddenSignal,
      pageTitle: signal<string | null>(null),
      set: jest.fn(),
      setHidden: jest.fn(),
      setPageTitle: jest.fn(),
      clear: jest.fn(),
    };

    await TestBed.configureTestingModule({
      imports: [BreadcrumbsComponent, RouterModule.forRoot([])],
      providers: [{ provide: BreadcrumbService, useValue: breadcrumbServiceMock }],
    }).compileComponents();

    fixture = TestBed.createComponent(BreadcrumbsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should not render nav when baseItems is empty', () => {
    const nav = fixture.nativeElement.querySelector('nav');
    expect(nav).toBeNull();
  });

  it('should be visible by default (isHidden returns false)', () => {
    expect(component.isHidden()).toBe(false);
  });

  it('should reflect hidden state from BreadcrumbService', () => {
    hiddenSignal.set(true);
    expect(component.isHidden()).toBe(true);
  });

  it('should include extra item when service has one', () => {
    const base: BreadcrumbItem = { label: 'Dashboard', href: '/dashboard' };
    const extra: BreadcrumbItem = { label: 'Extra Page' };
    component.baseItems = [base];
    extraSignal.set(extra);
    expect(component.items()).toContainEqual(extra);
  });

  it('should return only baseItems when service extra is null', () => {
    const base: BreadcrumbItem = { label: 'Dashboard', href: '/dashboard' };
    component.baseItems = [base];
    // Trigger recompute by toggling extra so computed re-reads baseItems
    extraSignal.set({ label: 'temp' });
    extraSignal.set(null);
    expect(component.items()).toEqual([base]);
  });
});
