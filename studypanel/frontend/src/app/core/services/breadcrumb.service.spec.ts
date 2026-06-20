import { TestBed } from '@angular/core/testing';
import { BreadcrumbService } from './breadcrumb.service';

describe('BreadcrumbService', () => {
  let service: BreadcrumbService;

  beforeEach(() => {
    TestBed.configureTestingModule({ providers: [BreadcrumbService] });
    service = TestBed.inject(BreadcrumbService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should start with extra as null', () => {
    expect(service.extra()).toBeNull();
  });

  it('should start with hidden as false', () => {
    expect(service.hidden()).toBe(false);
  });

  it('should start with pageTitle as null', () => {
    expect(service.pageTitle()).toBeNull();
  });

  it('should update extra via set()', () => {
    service.set({ label: 'Node.js' });
    expect(service.extra()).toEqual({ label: 'Node.js' });
  });

  it('should update hidden via setHidden()', () => {
    service.setHidden(true);
    expect(service.hidden()).toBe(true);
  });

  it('should update pageTitle via setPageTitle()', () => {
    service.setPageTitle('Backend');
    expect(service.pageTitle()).toBe('Backend');
  });

  it('should reset all signals on clear()', () => {
    service.set({ label: 'Node.js' });
    service.setHidden(true);
    service.setPageTitle('Backend');
    service.clear();
    expect(service.extra()).toBeNull();
    expect(service.hidden()).toBe(false);
    expect(service.pageTitle()).toBeNull();
  });
});
