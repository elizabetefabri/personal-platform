# Changelog — StudyPanel Frontend

All notable changes to the StudyPanel frontend are documented here.

---

## [Unreleased] — 2026-06-20

### Added
- **BackButtonComponent** inline in `StudyCardGrid`, `ProjectCardGrid`, and `StudyDetailTemplate` headers (replaces fixed overlay)
- **New pages**: `EstudosLabsComponent` (`/estudos-labs`), `ProjetosComponent` (`/projetos`), `ProjetosPessoaisComponent` (`/projetos/pessoais`), `ProjetosProfissionaisComponent` (`/projetos/profissionais`)
- **BreadcrumbService** unit tests (`breadcrumb.service.spec.ts`)
- **BackButtonComponent** unit tests (`back-button.spec.ts`) — 10 tests with `mockResolvedValue` pattern
- **ProjectCardGrid** unit tests (`project-card-grid.spec.ts`) — 8 tests
- **StudyCardGrid** unit tests expanded — 7 tests with Router and input validation
- **StudyDetailTemplate** unit tests with `StudyItemService` mock
- **EstudosLabs, ProjetosPessoais, ProjetosProfissionais** component specs
- **Playwright E2E**: installed `@playwright/test` v1.61, `playwright.config.ts`, 7 spec files
- **Documentation**: `.AGENTS.md` comprehensive guide, `changelog.md`, `WIP.md` updated
- Scripts `test:e2e` and `test:e2e:ui` in `package.json`

### Changed
- **Sidebar**: reduced from original count to exactly 3 items (Dashboard, Estudos e Labs, Projetos)
- **Sidebar active state**: `/backend`, `/cloud`, etc. now activate `estudos-labs`; `/rollout-service` activates `projetos`
- **BackButton**: moved from `position: fixed` overlay in `app.html` to inline placement within content headers
- **BreadcrumbsComponent**: `build()` rewritten for correct hierarchy including "Estudos e Labs" intermediate step
- **`setup-jest.ts`**: updated from `jest-preset-angular/setup-jest` to direct `@angular/core/testing` setup for Angular 21 zoneless compatibility
- All study section page specs: added `RouterModule.forRoot([])` and `StudyItemService` mock

### Fixed
- Sidebar spec: items count corrected from 4 to 3, `spyOn` → `jest.spyOn`, `toBeTrue()` → `toBe(true)`
- Breadcrumbs spec: `component.items = items` (invalid) → `component.baseItems = items` with signal recompute trigger
- Dashboard spec: fixed h1 assertion to match actual template (h2 with "Seções de estudo")
- Header spec: removed `component.title` assertion (property doesn't exist)
- All page specs with `StudyDetailTemplate` now provide `StudyItemService` mock

---

## [0.3.0] — 2026-06-13

### Added
- Responsive grids: 1→2→3→4→6 columns (breakpoints 500/768/1024/1440px)
- `clamp()` for responsive fonts and gaps
- Image support in cards with fallback to `bannerColor` gradient
- SEO: `lang="pt-BR"`, title, description, keywords, robots, OG tags in `index.html`
- `public/robots.txt` and `public/sitemap.xml`
- `BreadcrumbService` with Angular signals (`extra`, `hidden`, `pageTitle`)
- `BackButtonComponent` (initial implementation)
- `StudyCardGrid`, `ProjectCardGrid` shared components
- `StudyDetailTemplate` with PrimeNG Table, Dialog, CRUD

---

## [0.2.0] — 2026-06-06

### Added
- PrimeNG and PrimeIcons integration
- CSS layout variables (`--sidebar-width-expanded`, `--sidebar-width-collapsed`, etc.)
- All page components extracted to Angular structure (`.ts` + `.html` + `.scss` + `.spec.ts`)
- Sidebar with expand/collapse, toggle, PrimeIcons
- BreadcrumbsComponent with `pi-chevron-right` separators
- Footer, Header components
- Angular Router with lazy-loaded routes

---

## [0.1.0] — Initial

- Project scaffolded with Angular CLI
- Basic routing and standalone component structure
