# Estrutura do Projeto

## Status

Existente

## Visão geral

```
frontend/
├── src/
│   ├── app/
│   │   ├── app.ts / app.html / app.scss  # Root + layout global
│   │   ├── app.routes.ts                  # Rotas lazy-loaded
│   │   ├── core/
│   │   │   ├── constants/
│   │   │   │   └── study-topics.ts        # STUDY_SECTIONS — única fonte de dados de seções
│   │   │   └── services/
│   │   │       ├── breadcrumb.service.ts
│   │   │       └── study-item.service.ts
│   │   ├── pages/
│   │   │   ├── dashboard/
│   │   │   ├── estudos-labs/
│   │   │   ├── projetos/
│   │   │   ├── projetos-pessoais/
│   │   │   ├── projetos-profissionais/
│   │   │   ├── backend/ | cloud/ | devops/ | ...  # Seções de estudo
│   │   │   ├── rollout-service/
│   │   │   ├── course-detail/
│   │   │   └── settings/
│   │   └── shared/
│   │       ├── components/
│   │       │   ├── back-button/
│   │       │   ├── breadcrumbs/
│   │       │   ├── footer/ | header/ | sidebar/
│   │       │   ├── project-card-grid/
│   │       │   ├── study-card-grid/
│   │       │   └── study-detail-template/
│   │       └── interfaces/
│   │           └── study-template.interface.ts
│   ├── styles/                  # Estilos globais e abstracts
│   ├── setup-jest.ts
│   └── index.html               # SEO meta tags
├── public/
│   ├── robots.txt
│   ├── sitemap.xml
│   └── assets/images/           # Imagens de cards
├── tests/e2e/                   # Playwright specs
├── documentation/               # Esta documentação
├── jest.config.ts
└── playwright.config.ts
```

## Padrões

- **Standalone components** — sem NgModules, `standalone: true`
- **Importações explícitas** — cada componente declara seus próprios `imports`
- **CSS BEM** — `.componentName__element--modifier`
- **Sem comentários** — exceto quando o WHY é não-óbvio
- **Signals** — `signal()` e `computed()` preferíveis a propriedades mutáveis
