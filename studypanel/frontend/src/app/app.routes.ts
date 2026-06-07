import { Routes } from '@angular/router';

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
    path: 'active-releases',
    loadComponent: () =>
      import('./pages/active-releases/active-releases.component').then(
        (m) => m.ActiveReleasesComponent,
      ),
  },
  {
    path: 'rollouts',
    loadComponent: () =>
      import('./pages/rollouts/rollouts.component').then((m) => m.RolloutsComponent),
  },
  {
    path: 'settings',
    loadComponent: () =>
      import('./pages/settings/settings.component').then((m) => m.SettingsComponent),
  },
  {
    path: 'login',
    loadComponent: () =>
      import('./pages/login/login.component').then((m) => m.LoginComponent),
  },
  {
    path: 'perfil',
    loadComponent: () =>
      import('./pages/perfil/perfil.component').then((m) => m.PerfilComponent),
  },
  {
    path: 'cadastro',
    loadComponent: () =>
      import('./pages/cadastro/cadastro.component').then((m) => m.CadastroComponent),
  },
  {
    path: 'cloud',
    loadComponent: () =>
      import('./pages/cloud/cloud.component').then((m) => m.CloudComponent),
  },
  {
    path: 'rollout',
    loadComponent: () =>
      import('./pages/rollout/rollout.component').then((m) => m.RolloutComponent),
  },
  {
    path: '**',
    redirectTo: '/dashboard',
  },
];
