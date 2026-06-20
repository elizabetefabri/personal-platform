import { NgClass } from '@angular/common';
import { Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { NavigationEnd, Router, RouterLink } from '@angular/router';
import { filter, Subscription } from 'rxjs';

export type SidebarItemKey = 'dashboard' | 'estudos-labs' | 'projetos';

export interface SidebarItem {
  key: SidebarItemKey;
  label: string;
  href: string;
  iconClass: string;
}

const ESTUDOS_PREFIXES = [
  '/estudos-labs',
  '/backend',
  '/banco-de-dados',
  '/cloud',
  '/containers-kubernetes',
  '/devops',
  '/frontend',
  '/inteligencia-artificial',
  '/observability',
  '/performance-engineering',
];

const PROJETOS_PREFIXES = ['/projetos', '/rollout-service'];

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [NgClass, RouterLink],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent implements OnInit, OnDestroy {
  @Input() expanded = true;
  @Output() toggleExpanded = new EventEmitter<void>();

  activeItem: SidebarItemKey = 'dashboard';

  private sub = new Subscription();

  readonly items: SidebarItem[] = [
    { key: 'dashboard', label: 'Dashboard', href: '/dashboard', iconClass: 'pi-th-large' },
    { key: 'estudos-labs', label: 'Estudos e Labs', href: '/estudos-labs', iconClass: 'pi-book' },
    { key: 'projetos', label: 'Projetos', href: '/projetos', iconClass: 'pi-folder' },
  ];

  constructor(private router: Router) {}

  ngOnInit(): void {
    this.updateActive(this.router.url);

    this.sub.add(
      this.router.events
        .pipe(filter((e): e is NavigationEnd => e instanceof NavigationEnd))
        .subscribe((e) => this.updateActive(e.urlAfterRedirects)),
    );
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  private updateActive(url: string): void {
    const firstSegment = '/' + (url.split('?')[0].split('/').filter(Boolean)[0] ?? '');

    if (ESTUDOS_PREFIXES.includes(firstSegment)) {
      this.activeItem = 'estudos-labs';
      return;
    }
    if (PROJETOS_PREFIXES.includes(firstSegment)) {
      this.activeItem = 'projetos';
      return;
    }
    this.activeItem = 'dashboard';
  }

  get sidebarWidth(): string {
    return this.expanded ? 'var(--sidebar-width-expanded)' : 'var(--sidebar-width-collapsed)';
  }

  signOut(): void {
    console.log('Usuário clicou em sair.');
  }
}
