import { Component, Input, Output, EventEmitter } from '@angular/core';
import { RouterLink } from '@angular/router';
import { NgClass } from '@angular/common';

export type SidebarItemKey = 'dashboard' | 'active-releases' | 'rollouts' | 'settings';

export interface SidebarItem {
  key: SidebarItemKey;
  label: string;
  href: string;
  iconClass: string;
}

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterLink, NgClass],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent {
  @Input() expanded = true;
  @Input() activeItem: SidebarItemKey = 'dashboard';
  @Output() toggleExpanded = new EventEmitter<void>();

  readonly items: SidebarItem[] = [
    { key: 'dashboard', label: 'Dashboard', href: '/dashboard', iconClass: 'pi-th-large' },
    { key: 'active-releases', label: 'Active Releases', href: '/active-releases', iconClass: 'pi-send' },
    { key: 'rollouts', label: 'Rollouts', href: '/rollouts', iconClass: 'pi-clone' },
    { key: 'settings', label: 'Configurações', href: '/settings', iconClass: 'pi-cog' },
  ];

  get sidebarWidth(): string {
    return this.expanded ? 'var(--sidebar-width-expanded)' : 'var(--sidebar-width-collapsed)';
  }
}
