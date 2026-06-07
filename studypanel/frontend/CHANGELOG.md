# Changelog — StudyPanel Frontend

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.
Formato baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/).

---

## [Unreleased]

### Adicionado
- Instalação da biblioteca `primeicons` para ícones vetoriais via CSS
- Variáveis CSS de layout no design system: `--sidebar-width-expanded`, `--sidebar-width-collapsed`, `--header-height`, `--icon-size`, `--text-xs`
- Componente `SidebarComponent` reimplementado com suporte a expand/collapse, PrimeIcons e `@Input`/`@Output` Angular
- Componente `BreadcrumbsComponent` novo, com separador de ícone (`pi-chevron-right`), suporte a `@Input() items` e acessibilidade (`aria-current`)
- Páginas `ActiveReleasesComponent` e `RolloutsComponent` com rotas `/active-releases` e `/rollouts`
- Rotas `/active-releases` e `/rollouts` adicionadas ao `app.routes.ts`
- Arquivos `.spec.ts` para todos os componentes (header, sidebar, footer, breadcrumbs, todas as páginas)
- Arquivo `WIP.md` com status de desenvolvimento
- Diretório `.agents/` com configuração de agentes de desenvolvimento

### Alterado
- Todos os componentes migrados de `.tsx` (inline template/styles) para estrutura Angular padrão: `.ts` + `.html` + `.scss` + `.spec.ts`
- Componente `AppComponent` (`app.tsx` → `app.ts` + `app.html`): adicionado sinal `sidebarExpanded` e método `toggleSidebar()`
- `app.scss` atualizado para usar variáveis CSS (`--sidebar-width-expanded/collapsed`) e classes `.sidebar-expanded`/`.sidebar-collapsed` com transição suave
- `angular.json`: PrimeIcons CSS adicionado ao array `styles`
- `header.component.html`: corrigido uso de `className` (JSX) para `class` (HTML padrão Angular), title atualizado para 'StudyPanel'
- Rotas `/active-releases` e `/rollouts` priorizadas sobre as antigas `/cloud` e `/rollout` na sidebar

### Removido
- Todos os arquivos `.tsx` substituídos por `.ts` com arquivos externos de template/estilos

---

## [0.1.0] — 2026-01-01

### Adicionado
- Estrutura inicial do projeto Angular 21
- Configuração NgRx (store, effects, router-store, devtools)
- Design system com variáveis CSS em `_variables.scss`
- Componentes base: Header, Sidebar, Footer, Breadcrumbs
- Páginas: Dashboard, Login, Perfil, Cadastro, Cloud, Rollout, Settings
- Lazy loading para todas as rotas
- SCSS com arquitetura SMACSS (abstracts, base, components, layouts, pages, themes)
- Reset CSS e tipografia base (Open Sans / Lato)
- Suporte a deploy via FTP (`basic-ftp`)
