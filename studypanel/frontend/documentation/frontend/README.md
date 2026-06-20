# Frontend StudyPanel — Documentação Técnica

## Stack

| Tecnologia    | Versão | Finalidade                                 |
| ------------- | ------ | ------------------------------------------ |
| Angular       | 21     | Framework (standalone, zoneless)           |
| TypeScript    | ~5.9   | Linguagem                                  |
| PrimeNG       | 19+    | Componentes UI                             |
| PrimeIcons    | 7+     | Ícones (`pi-*`)                            |
| SCSS          | —      | Estilos (custom properties, clamp())       |
| RxJS          | 7+     | Streams e eventos do Router                |
| Jest          | 30     | Testes unitários                           |
| Playwright    | 1.61   | Testes E2E                                 |

## Estrutura de Diretórios

```
src/app/
├── app.ts / app.html / app.scss    # Root component + layout global
├── app.routes.ts                    # Rotas lazy-loaded
├── core/
│   ├── constants/study-topics.ts   # Fonte única de dados das seções
│   └── services/
│       ├── breadcrumb.service.ts
│       └── study-item.service.ts
├── pages/                           # Páginas (uma por rota)
└── shared/
    ├── components/                  # Componentes reutilizáveis
    └── interfaces/                  # Interfaces TypeScript
```

## Seções desta documentação

- [Arquitetura](./architecture/estrutura-do-projeto.md)
- [Páginas](./pages/README.md)
- [Componentes](./components/README.md)
- [Serviços](./services/README.md)
- [Testes](./tests/README.md)
