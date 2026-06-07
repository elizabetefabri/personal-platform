import { NgClass } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { RouterLink } from '@angular/router';

export type SidebarItemKey =
  | 'backend'
  | 'cloud'
  | 'containers-kubernetes'
  | 'dashboard'
  | 'database'
  | 'devops'
  | 'frontend'
  | 'artificial-intelligence'
  | 'observability'
  | 'performance-engineering'
  | 'projects'
  | 'rollout-service';

export interface SidebarItem {
  key: SidebarItemKey;
  label: string;
  href: string;
  iconClass: string;
}

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [NgClass, RouterLink],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent {
  @Input() expanded = true;
  @Input() activeItem: SidebarItemKey = 'dashboard';
  @Output() toggleExpanded = new EventEmitter<void>();

  readonly items: SidebarItem[] = [
    {
      key: 'dashboard',
      label: 'Dashboard',
      href: '/dashboard',
      iconClass: 'pi-th-large',
    },
    {
      key: 'backend',
      label: 'Backend',
      href: '/backend',
      iconClass: 'pi-server',
    },
    {
      key: 'database',
      label: 'Banco de Dados',
      href: '/banco-de-dados',
      iconClass: 'pi-database',
    },
    {
      key: 'cloud',
      label: 'Estudo Cloud Computing',
      href: '/cloud',
      iconClass: 'pi-cloud',
    },
    {
      key: 'containers-kubernetes',
      label: 'Containers e Kubernetes',
      href: '/containers-kubernetes',
      iconClass: 'pi-box',
    },
    {
      key: 'devops',
      label: 'DevOps',
      href: '/devops',
      iconClass: 'pi-cog',
    },
    {
      key: 'frontend',
      label: 'Frontend',
      href: '/frontend',
      iconClass: 'pi-desktop',
    },
    {
      key: 'artificial-intelligence',
      label: 'Inteligência Artificial',
      href: '/inteligencia-artificial',
      iconClass: 'pi-sparkles',
    },
    {
      key: 'observability',
      label: 'Estudo Observability',
      href: '/observability',
      iconClass: 'pi-eye',
    },
    {
      key: 'performance-engineering',
      label: 'Performance Engineering',
      href: '/performance-engineering',
      iconClass: 'pi-chart-line',
    },
    {
      key: 'projects',
      label: 'Projetos',
      href: '/projetos',
      iconClass: 'pi-folder',
    },
    {
      key: 'rollout-service',
      label: 'Rollout Service',
      href: '/rollout-service',
      iconClass: 'pi-clone',
    },
  ];

  get sidebarWidth(): string {
    return this.expanded ? 'var(--sidebar-width-expanded)' : 'var(--sidebar-width-collapsed)';
  }

  signOut(): void {
    console.log('Usuário clicou em sair.');
  }
}
