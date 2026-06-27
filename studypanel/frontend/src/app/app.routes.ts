import { Routes } from '@angular/router';

const sectionChild = {
  path: '',
  loadComponent: () =>
    import('./pages/study-section-page/study-section-page.component').then(
      (m) => m.StudySectionPageComponent,
    ),
};

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
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  {
    path: 'dashboard',
    loadComponent: () =>
      import('./pages/dashboard/dashboard.component').then((m) => m.DashboardComponent),
  },
  { path: 'backend', children: [sectionChild, detailChild] },
  { path: 'banco-de-dados', children: [sectionChild, detailChild] },
  { path: 'cloud', children: [sectionChild, detailChild] },
  { path: 'containers-kubernetes', children: [sectionChild, detailChild] },
  { path: 'devops', children: [sectionChild, detailChild] },
  { path: 'frontend', children: [sectionChild, detailChild] },
  { path: 'inteligencia-artificial', children: [sectionChild, detailChild] },
  { path: 'observability', children: [sectionChild, detailChild] },
  { path: 'performance-engineering', children: [sectionChild, detailChild] },
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
    children: [sectionChild, detailChild],
  },
  {
    path: 'vida-criativa',
    loadComponent: () =>
      import('./pages/vida-criativa/vida-criativa.component').then((m) => m.VidaCriativaComponent),
  },
  {
    path: 'painel-financeiro',
    loadComponent: () =>
      import('./pages/painel-financeiro/painel-financeiro.component').then(
        (m) => m.PainelFinanceiroComponent,
      ),
  },
    {
    path: 'culinaria',
    loadComponent: () =>
      import('./pages/culinaria/culinaria').then(
        (m) => m.Culinaria,
      ),
  },
  {
    path: 'settings',
    loadComponent: () =>
      import('./pages/settings/settings.component').then((m) => m.SettingsComponent),
  },
  { path: '**', redirectTo: '/dashboard' },
];
