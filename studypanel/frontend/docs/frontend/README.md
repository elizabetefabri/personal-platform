# StudyPanel Frontend Documentation

Documentação completa da arquitetura frontend do StudyPanel, incluindo componentes compartilhados e páginas.

## Estrutura do Projeto

```
src/app/
├── shared/
│   ├── components/        # Componentes reutilizáveis
│   └── services/         # Serviços compartilhados
├── pages/                 # Páginas (rotas principais)
└── app.component.ts       # Componente raiz
```

## Navegação

### Componentes Compartilhados

- [Sidebar](./components/SIDEBAR.md) - Navegação lateral colapsável
- [Header](./components/HEADER.md) - Cabeçalho com título da aplicação
- [Breadcrumbs](./components/BREADCRUMBS.md) - Navegação por breadcrumbs
- [Footer](./components/FOOTER.md) - Rodapé da aplicação
- [Study Card Grid](./components/STUDY_CARD_GRID.md) - Grade de cards de estudos

### Páginas

#### Categorias de Estudo

- [Dashboard](./pages/DASHBOARD.md) - Página inicial/dashboard
- [Backend](./pages/BACKEND.md) - Conteúdos de Backend
- [Banco de Dados](./pages/BANCO_DADOS.md) - Conteúdos de Banco de Dados
- [Cloud Computing](./pages/CLOUD.md) - Conteúdos de Cloud
- [Containers & Kubernetes](./pages/CONTAINERS_K8S.md) - Conteúdos de Containers
- [DevOps](./pages/DEVOPS.md) - Conteúdos de DevOps
- [Frontend](./pages/FRONTEND.md) - Conteúdos de Frontend
- [Inteligência Artificial](./pages/IA.md) - Conteúdos de IA/ML
- [Observability](./pages/OBSERVABILITY.md) - Conteúdos de Observability
- [Performance Engineering](./pages/PERFORMANCE.md) - Conteúdos de Performance
- [Projetos](./pages/PROJETOS.md) - Página de projetos
- [Rollout Service](./pages/ROLLOUT.md) - Conteúdos de Rollout

#### Outras Páginas

- [Settings](./pages/SETTINGS.md) - Configurações da aplicação
- [Study Detail Page](./pages/STUDY_DETAIL.md) - Página de detalhes do estudo

## Design System

Consulte [`styles/abstracts/_variables.scss`](../../src/styles/abstracts/_variables.scss) para:

- Cores e paleta
- Tipografia
- Spacing
- Breakpoints
- Shadows e transitions

## Routing

O projeto utiliza Angular Router com rotas aninhadas. Consulte `app.routes.ts` para a configuração completa.

## Estado e Dados

### Serviços

- Verificar `src/app/shared/services/` para serviços de dados
- O projeto pode utilizar signals para estado reativo

## Build & Deploy

- Framework: Angular 19+
- Bundler: Vite/Turbopack
- Styling: SCSS com variáveis CSS
- Components: Standalone components (Angular 14+)
- Icons: PrimeIcons
