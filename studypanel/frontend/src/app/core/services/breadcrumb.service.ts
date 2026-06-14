import { Injectable, signal } from '@angular/core';

export interface BreadcrumbItem {
  label: string;
  href?: string;
}

@Injectable({ providedIn: 'root' })
export class BreadcrumbService {
  readonly extra = signal<BreadcrumbItem | null>(null);
  readonly pageTitle = signal<string | null>(null);

  set(item: BreadcrumbItem | null): void {
    this.extra.set(item);
  }

  setPageTitle(title: string | null): void {
    this.pageTitle.set(title);
  }

  clear(): void {
    this.extra.set(null);
    this.pageTitle.set(null);
  }
}
