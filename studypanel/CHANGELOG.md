# Changelog — StudyPanel

Formato baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/).

---

## [Unreleased]

---

## [0.5.0] — 2026-06-28

### Fixed
- Modais de `estudos-labs` não disparavam eventos ao clicar em "Cadastrar" e "Excluir" — causa: `<button pButton (onClick)>` não suporta `(onClick)` (evento exclusivo do componente `<p-button>`). Restaurado para `<p-button styleClass="btn-primary/btn-cancel" (onClick)>`.
- Dado de teste "QA Engineer" (slug `qa`, sem rota) removido do banco.

### Changed
- `WIP.md` e `CHANGELOG.md` atualizados com estado real do projeto.

---

## [0.4.0] — 2026-06-28

### Added
- Dados 100% do banco — `FALLBACK`, `IMAGE_MAP` e `fallbackCards()` removidos de `estudos-labs`, `projetos-pessoais` e `projetos-profissionais`.
- Empty state visual (`@empty` block) nos grids `StudyCardGrid` e `ProjectCardGrid`.
- Cores dos cards (bannerColor) dinâmicas vindas do banco.
- **Backend — novos testes Go:**
  - `course_section_usecase_test.go`: 9 testes CRUD
  - `project_usecase_test.go`: 9 testes CRUD
  - `course_section_handler_test.go`: 7 testes HTTP
  - `project_handler_test.go`: 8 testes HTTP
- **Frontend — novos testes Jest 30 (63 testes passando):**
  - `course-section.service.spec.ts`: 8 testes
  - `project.service.spec.ts`: 9 testes
  - `estudos-labs.component.spec.ts`: 14 testes
  - `projetos-pessoais.component.spec.ts`: 15 testes
  - `projetos-profissionais.component.spec.ts`: 15 testes

---

## [0.3.0] — 2026-06-27

### Added
- `COMMANDS.md` na raiz com guia completo de execução e consultas MongoDB.
- `documentation/database-queries.md` com queries por seção do sidebar.
- Novas coleções e API completa: `vida_criativa_items`, `financial_records`, `culinaria_recipes`.
- Correção de ícone dos cards — fallback PI exibido quando imagem falha ao carregar.
- Backend rebuilt com novas rotas registradas em `main.go`.

---

## [0.2.0] — 2026-06-27

### Added
- Padronização visual de botões: `btn-primary` (#201F25), `btn-cancel`, `btn-accent-outline`.
- Override das CSS vars do PrimeNG Aura em `styles.scss` (elimina verde padrão).
- Layout dos modais com footer 50/50 (flex: 1 por botão).
- Cards de projetos com botões fixos no rodapé via `margin-top: auto`.
- Descrição dos cards limitada a 3 linhas com `-webkit-line-clamp`.
- Ícone dos cards substituído por SVG proporcional.
- Ícone de editar nos cards com modal pré-preenchido.
- Back button visível em todas as páginas exceto `/dashboard`.

### Fixed
- Breadcrumbs reativos e cor amarela nos botões de projeto.

---

## [0.1.0] — 2026-06-13

### Added
- Estrutura inicial do projeto Angular 21 + Go 1.22 + MongoDB.
- Clean Architecture no backend: entity → repository interface → usecase → handler.
- Docker Compose: API + MongoDB + Mongo Express.
- Coleções: `study_items`, `study_notes`, `study_resources`, `study_sessions`, `quiz_questions`.
- Coleções: `course_sections`, `course_topics`, `projects`.
- API REST com CRUD completo para todos os recursos.
- Testes unitários backend (StudyItem: entity, usecase, handler).
- Componentes base: Header, Sidebar, Footer, Breadcrumbs, BackButton.
- Páginas de seção com rotas `/section/:topic/:itemId`.
- `StudyDetailTemplate` com tabela e modal de cadastro.
- `StudyCardGrid` e `ProjectCardGrid` com CRUD via modal.
- Documentação AWS SAA-C03 e agentes professores.
- `PADROES.md` e `README.md`.
