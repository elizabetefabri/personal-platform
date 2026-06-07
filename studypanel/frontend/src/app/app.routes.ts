import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    redirectTo: '/dashboard',
    pathMatch: 'full',
  },
  {
    path: 'backend',
    loadComponent: () =>
      import('./pages/backend/backend.component').then((m) => m.BackendComponent),
  },
  {
    path: 'banco-de-dados',
    loadComponent: () =>
      import('./pages/banco-de-dados/banco-de-dados.component').then(
        (m) => m.BancoDeDadosComponent,
      ),
  },
  {
    path: 'cloud',
    loadComponent: () =>
      import('./pages/cloud/cloud.component').then((m) => m.CloudComponent),
  },
  {
    path: 'containers-kubernetes',
    loadComponent: () =>
      import('./pages/containers-kubernetes/containers-kubernetes.component').then(
        (m) => m.ContainersKubernetesComponent,
      ),
  },
  {
    path: 'dashboard',
    loadComponent: () =>
      import('./pages/dashboard/dashboard.component').then((m) => m.DashboardComponent),
  },
  {
    path: 'devops',
    loadComponent: () =>
      import('./pages/devops/devops.component').then((m) => m.DevOpsComponent),
  },
  {
    path: 'frontend',
    loadComponent: () =>
      import('./pages/frontend/frontend.component').then((m) => m.FrontendComponent),
  },
  {
    path: 'inteligencia-artificial',
    loadComponent: () =>
      import('./pages/inteligencia-artificial/inteligencia-artificial.component').then(
        (m) => m.InteligenciaArtificialComponent,
      ),
  },
  {
    path: 'observability',
    loadComponent: () =>
      import('./pages/observability/observability.component').then(
        (m) => m.ObservabilityComponent,
      ),
  },
  {
    path: 'performance-engineering',
    loadComponent: () =>
      import('./pages/performance-engineering/performance-engineering.component').then(
        (m) => m.PerformanceEngineeringComponent,
      ),
  },
  {
    path: 'projetos',
    loadComponent: () =>
      import('./pages/projetos/projetos.component').then((m) => m.ProjetosComponent),
  },
  {
    path: 'rollout-service',
    loadComponent: () =>
      import('./pages/rollout-service/rollout-service.component').then(
        (m) => m.RolloutServiceComponent,
      ),
  },
  {
    path: 'settings',
    loadComponent: () =>
      import('./pages/settings/settings.component').then((m) => m.SettingsComponent),
  },
  {
    path: '**',
    redirectTo: '/dashboard',
  },
];
