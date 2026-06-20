# Testes

## Stack de Testes

| Ferramenta  | Versão | Uso                              |
| ----------- | ------ | -------------------------------- |
| Jest        | 30     | Testes unitários + cobertura     |
| Playwright  | 1.61   | Testes E2E                       |

## Testes unitários

### Configuração

- `jest.config.ts` — preset `jest-preset-angular`
- `src/setup-jest.ts` — TestBed manual (Angular 21 zoneless, sem `zone.js`)

```ts
// setup-jest.ts — não usar setup-jest do preset (incompatível com zoneless)
getTestBed().initTestEnvironment(
  BrowserTestingModule,
  platformBrowserTesting([{ provide: COMPILER_OPTIONS, useValue: {}, multi: true }]),
  { errorOnUnknownElements: false, errorOnUnknownProperties: false },
);
```

### Rodar

```bash
# Na pasta frontend/
npm test                # once
npm run test:watch      # watch mode
npm run test:coverage   # com coverage
```

### Status: 107 testes passando (26 suites)

### Padrões de teste

1. **BackButton**: `jest.spyOn(router, 'navigateByUrl').mockResolvedValue(true)` obrigatório
2. **BreadcrumbsComponent**: toggle `extra` signal para forçar recompute
3. **StudyDetailTemplate**: sempre mockar `StudyItemService` + `RouterModule.forRoot([])`
4. **Seções de estudo** (backend, cloud etc.): mesmo mock acima

## Testes E2E

### Configuração

`playwright.config.ts` — `testDir: './tests/e2e'`, `baseURL: 'http://localhost:4200'`

### Rodar

```bash
# Instalar browsers (uma vez):
npx playwright install chromium

# Executar:
npx playwright test               # headless
npx playwright test --ui          # interface visual
npx playwright test --headed      # headed
```

### Specs em `tests/e2e/`

| Arquivo                  | O que testa                                   |
| ------------------------ | --------------------------------------------- |
| `navigation.spec.ts`     | Navegação entre rotas principais              |
| `sidebar.spec.ts`        | Itens e active state da sidebar               |
| `breadcrumbs.spec.ts`    | Hierarquia de breadcrumbs por rota            |
| `projects.spec.ts`       | Fluxo de projetos pessoais/profissionais      |
| `studies-labs.spec.ts`   | Cards e grid de estudos e labs                |
| `responsive-grid.spec.ts`| Responsividade em diferentes viewports        |
| `seo.spec.ts`            | Meta tags, robots.txt, sitemap.xml            |

## Cobertura esperada

- Componentes: `>= 80%`
- Serviços: `>= 90%`
- Constantes: `>= 95%`
