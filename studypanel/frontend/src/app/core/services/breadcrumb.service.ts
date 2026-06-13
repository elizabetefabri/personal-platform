import { Injectable, signal } from '@angular/core';

export interface BreadcrumbItem {
  label: string;
  href?: string;
}

@Injectable({ providedIn: 'root' })
export class BreadcrumbService {
  readonly extra = signal<BreadcrumbItem | null>(null);

  set(item: BreadcrumbItem | null): void {
    this.extra.set(item);
  }

  clear(): void {
    this.extra.set(null);
  }
}
