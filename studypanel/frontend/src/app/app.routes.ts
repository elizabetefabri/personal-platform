import { Routes } from '@angular/router';

const detailChild = {
  path: ':topic',
  children: [
    {
      path: '',
      loadComponent: () =>
        import('./pages/study-detail-page/study-detail-page.component').then(
          (m) => m.StudyDetailPageComponent,
        ),
    },
    {
      path: ':itemId',
      loadComponent: () =>
        import('./pages/course-detail/course-detail.component').then(
          (m) => m.CourseDetailComponent,
        ),
    },
  ],
};

export const routes: Routes = [
  {
    path: '',
    redirectTo: '/dashboard',
    pathMatch: 'full',
  },
  {
    path: 'dashboard',
    loadComponent: () =>
      import('./pages/dashboard/dashboard.component').then((m) => m.DashboardComponent),
  },
  {
    path: 'backend',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/backend/backend.component').then((m) => m.BackendComponent),
      },
      detailChild,
    ],
  },
  {
    path: 'banco-de-dados',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/banco-de-dados/banco-de-dados.component').then(
            (m) => m.BancoDeDadosComponent,
          ),
      },
      detailChild,
    ],
  },
  {
    path: 'cloud',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/cloud/cloud.component').then((m) => m.CloudComponent),
      },
      detailChild,
    ],
  },
  {
    path: 'containers-kubernetes',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/containers-kubernetes/containers-kubernetes.component').then(
            (m) => m.ContainersKubernetesComponent,
          ),
      },
      detailChild,
    ],
  },
  {
    path: 'devops',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/devops/devops.component').then((m) => m.DevOpsComponent),
      },
      detailChild,
    ],
  },
  {
    path: 'frontend',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/frontend/frontend.component').then((m) => m.FrontendComponent),
      },
      detailChild,
    ],
  },
  {
    path: 'inteligencia-artificial',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/inteligencia-artificial/inteligencia-artificial.component').then(
            (m) => m.InteligenciaArtificialComponent,
          ),
      },
      detailChild,
    ],
  },
  {
    path: 'observability',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/observability/observability.component').then(
            (m) => m.ObservabilityComponent,
          ),
      },
      detailChild,
    ],
  },
  {
    path: 'performance-engineering',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/performance-engineering/performance-engineering.component').then(
            (m) => m.PerformanceEngineeringComponent,
          ),
      },
      detailChild,
    ],
  },
  {
    path: 'estudos-labs',
    loadComponent: () =>
      import('./pages/estudos-labs/estudos-labs.component').then((m) => m.EstudosLabsComponent),
  },
  {
    path: 'projetos',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/projetos/projetos.component').then((m) => m.ProjetosComponent),
      },
      {
        path: 'pessoais',
        loadComponent: () =>
          import('./pages/projetos-pessoais/projetos-pessoais.component').then(
            (m) => m.ProjetosPessoaisComponent,
          ),
      },
      {
        path: 'profissionais',
        loadComponent: () =>
          import('./pages/projetos-profissionais/projetos-profissionais.component').then(
            (m) => m.ProjetosProfissionaisComponent,
          ),
      },
    ],
  },
  {
    path: 'rollout-service',
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/rollout-service/rollout-service.component').then(
            (m) => m.RolloutServiceComponent,
          ),
      },
      detailChild,
    ],
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
