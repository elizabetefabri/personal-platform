# Rotas

## Status

Existente

## Arquivo

`src/app/app.routes.ts`

## Tabela de Rotas

| Rota                       | Componente                       | Tipo             |
| -------------------------- | -------------------------------- | ---------------- |
| `/`                        | redirect → `/dashboard`          | Redirect         |
| `/dashboard`               | DashboardComponent               | Página raiz      |
| `/estudos-labs`            | EstudosLabsComponent             | Página raiz      |
| `/projetos`                | ProjetosComponent                | Página raiz      |
| `/projetos/pessoais`       | ProjetosPessoaisComponent        | Sub-rota         |
| `/projetos/profissionais`  | ProjetosProfissionaisComponent   | Sub-rota         |
| `/backend`                 | BackendComponent                 | Seção de estudo  |
| `/banco-de-dados`          | BancoDeDadosComponent            | Seção de estudo  |
| `/cloud`                   | CloudComponent                   | Seção de estudo  |
| `/containers-kubernetes`   | ContainersKubernetesComponent    | Seção de estudo  |
| `/devops`                  | DevOpsComponent                  | Seção de estudo  |
| `/frontend`                | FrontendComponent                | Seção de estudo  |
| `/inteligencia-artificial` | InteligenciaArtificialComponent  | Seção de estudo  |
| `/observability`           | ObservabilityComponent           | Seção de estudo  |
| `/performance-engineering` | PerformanceEngineeringComponent  | Seção de estudo  |
| `/rollout-service`         | RolloutServiceComponent          | Projeto profis.  |
| `/:section/:topic/:id`     | CourseDetailComponent            | Detalhe do curso |
| `/settings`                | SettingsComponent                | Configurações    |

## Padrões

- Todas as rotas são lazy-loaded via `loadComponent()`
- Rotas raiz: `/dashboard`, `/estudos-labs`, `/projetos` — sem botão Voltar
- Páginas de estudo (`/backend`, `/cloud`, etc.) → botão Voltar aparece inline

## Sidebar active state

| URL prefix           | Item ativo     |
| -------------------- | -------------- |
| `/dashboard`         | `dashboard`    |
| `/estudos-labs`      | `estudos-labs` |
| `/backend`, `/cloud` | `estudos-labs` |
| `/projetos`          | `projetos`     |
| `/rollout-service`   | `projetos`     |
