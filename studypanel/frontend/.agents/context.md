# Contexto do Projeto — StudyPanel Frontend

## Stack Tecnológica
- **Framework:** Angular 21 (standalone components)
- **Linguagem:** TypeScript 5.9 (strict mode)
- **Estilos:** SCSS com arquitetura SMACSS
- **State Management:** NgRx 21 (store, effects, router-store, signals)
- **Ícones:** PrimeIcons (classe CSS `pi pi-*`)
- **Build:** @angular/build (esbuild)
- **Testes:** Angular TestBed + Karma/Jasmine (spec.ts)

## Estrutura de Diretórios

```
src/app/
├── app.ts              # Componente raiz (sinal sidebarExpanded)
├── app.html            # Template raiz
├── app.scss            # Layout principal
├── app.routes.ts       # Rotas lazy-loaded
├── app.config.ts       # Providers (NgRx, Router, HTTP)
├── shared/
│   └── components/
│       ├── header/     # Cabeçalho da aplicação
│       ├── sidebar/    # Navegação lateral (expandível/recolhível)
│       ├── footer/     # Rodapé
│       └── breadcrumbs/ # Navegação por trilha
└── pages/
    ├── dashboard/
    ├── active-releases/
    ├── rollouts/
    ├── settings/
    ├── login/
    ├── perfil/
    ├── cadastro/
    ├── cloud/
    └── rollout/
```

## Design System (CSS Variables)

As variáveis de design estão em `src/styles/abstracts/_variables.scss`:

### Layout
- `--sidebar-width-expanded: 15rem`
- `--sidebar-width-collapsed: 4rem`
- `--header-height: 4rem`
- `--icon-size: 1.25rem`
- `--text-xs: 0.75rem`

### Cores principais
- `--color-sidebar-bg-dark: #201f25` (sidebar background)
- `--color-bg: #f3f1ee` (page background)
- `--color-surface: #fbfaf8` (content surface)
- `--color-primary: #e59a9a` (action primary)
- `--color-status-warning: #c58b2a` (active item highlight)
- `--color-text-on-dark: #ffffff` (text over dark backgrounds)

## Convenções de Código

1. Componentes Angular standalone, sem NgModule
2. Templates em arquivos `.html` separados (nunca inline no decorator)
3. Estilos em arquivos `.scss` separados
4. Testes em arquivos `.spec.ts` colocalizados
5. `@if` / `@for` (Angular 17+ control flow) em vez de `*ngIf`/`*ngFor`
6. Signals para estado local reativo (ex: `sidebarExpanded = signal(true)`)
7. Lazy loading em todas as rotas de página
8. `@Input()` / `@Output()` para comunicação entre componentes
9. Sem comentários desnecessários no código
