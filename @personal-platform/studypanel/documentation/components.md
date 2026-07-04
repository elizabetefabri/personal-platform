# StudyPanel - Components and Routes

This document lists all components currently implemented in the StudyPanel frontend and the routes they are associated with.

## Pages (routes)

- /dashboard — DashboardComponent (src/app/pages/dashboard/dashboard.component.tsx)
- /login — LoginComponent (src/app/pages/login/login.component.tsx)
- /perfil — PerfilComponent (src/app/pages/perfil/perfil.component.tsx)
- /cadastro — CadastroComponent (src/app/pages/cadastro/cadastro.component.tsx)
- /cloud — CloudComponent (src/app/pages/cloud/cloud.component.tsx)
- /rollout — RolloutComponent (src/app/pages/rollout/rollout.component.tsx)
- /settings — SettingsComponent (src/app/pages/settings/settings.component.tsx)

## Layout & Shared Components

- HeaderComponent — app/shared/components/header/header.component.tsx
  - Sticky header with navigation
- SidebarComponent — app/shared/components/sidebar/sidebar.component.tsx
  - Left fixed navigation
- FooterComponent — app/shared/components/footer/footer.component.tsx
  - Footer information
- BreadcrumbsComponent — app/shared/components/breadcrumbs/breadcrumbs.component.tsx
  - Breadcrumb trail component

## Notes

- All components are standalone components and use inline templates (.tsx files)
- Styling uses SCSS and the design system CSS custom properties defined in `src/styles/abstracts/_variables.scss`
- Add new components under `src/app/shared/components` or `src/app/pages` following the existing structure
