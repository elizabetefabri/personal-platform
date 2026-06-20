# Comandos e Regras — Frontend StudyPanel

> Angular 21 · TypeScript 5.9 · SCSS · PrimeNG 19 · Jest 30 · Playwright 1.61

---

## Setup Inicial

```bash
# Instalar dependências
npm install

# Instalar browsers do Playwright (uma vez)
npx playwright install chromium
```

---

## Desenvolvimento

```bash
# Iniciar servidor de desenvolvimento
npm start
# ou
ng serve
# → http://localhost:4200

# Build de produção
ng build --configuration production

# Gerar componente standalone
ng generate component pages/nome-da-pagina --standalone --style scss

# Gerar serviço
ng generate service core/services/nome-do-servico

# Verificar tipos TypeScript
npx tsc --noEmit
```

---

## Testes

```bash
# Testes unitários (uma vez)
npm test

# Testes em modo watch
npm run test:watch

# Testes com cobertura
npm run test:coverage

# E2E (Playwright) — headless
npx playwright test

# E2E com interface visual
npx playwright test --ui

# E2E headed (browser visível)
npx playwright test --headed

# E2E para um spec específico
npx playwright test tests/e2e/navigation.spec.ts
```

---

## Regras do Projeto

### Angular

- **Standalone apenas** — nunca usar `NgModule`
- **Zoneless** — `provideZonelessChangeDetection()` já configurado; não instalar `zone.js`
- **Imports explícitos** — cada componente declara seus próprios `imports: []`
- **Signals** — preferir `signal()`, `computed()` a propriedades mutáveis
- **Lazy loading** — todas as rotas usam `loadComponent()` em `app.routes.ts`

### TypeScript

- **Strict mode** habilitado — sem `any` implícito
- **Interfaces** em `src/app/shared/interfaces/`
- **Constantes** em `src/app/core/constants/`
- **Serviços** em `src/app/core/services/` — `providedIn: 'root'`

### Estilos

- **BEM**: `.componentName__element--modifier`
- **Nunca** usar `px` para `font-size` — usar `rem` ou `clamp()`
- **Custom properties**: `var(--sidebar-width-expanded)`, `var(--header-height)`, etc.
- **Breakpoints** (em rem): 31.25 (500px) / 48 (768px) / 64 (1024px) / 90 (1440px)
- **Grid responsivo**: 1 → 2 → 3 → 4 → 6 colunas

### Componentes

- **Sem comentários** — exceto quando o WHY é não-óbvio
- **Sem lógica** em templates — lógica fica no `.ts`
- **Sem `position: fixed`** no back-button — é inline
- **PrimeIcons**: sempre usar classes `pi-*`, não SVGs inline

### Navegação

- **Links internos**: `[routerLink]` — nunca `[href]`
- **Links externos**: `window.open(url, '_blank', 'noopener,noreferrer')`
- **Botão Voltar**: NÃO incluir em rotas raiz (`/dashboard`, `/estudos-labs`, `/projetos`)
- **Botão Voltar**: NÃO incluir em `course-detail` (já tem "Voltar para trilha")

### SEO

- `index.html` tem todas as meta tags (lang, description, og:*)
- `public/robots.txt` e `public/sitemap.xml` existem
- Não remover ou alterar `<html lang="pt-BR">`

---

## Estrutura de Diretórios

```
src/app/
├── app.ts / app.html / app.scss
├── app.routes.ts
├── core/
│   ├── constants/study-topics.ts   ← STUDY_SECTIONS (única fonte de dados)
│   └── services/
│       ├── breadcrumb.service.ts
│       └── study-item.service.ts
├── pages/                           ← uma pasta por rota
└── shared/
    ├── components/                  ← componentes reutilizáveis
    └── interfaces/                  ← interfaces TypeScript
```

---

## Sidebar

- **3 itens** + Sair: Dashboard · Estudos e Labs · Projetos
- Alterar itens → atualizar `sidebar.spec.ts` (teste espera `toBe(3)`)

---

## Padrões de Teste

```ts
// Sempre mockar StudyItemService em specs que usam StudyDetailTemplate
const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  listAll: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};

// Sempre incluir RouterModule.forRoot([]) quando BackButton está no componente
imports: [ComponentName, RouterModule.forRoot([])],
providers: [{ provide: StudyItemService, useValue: studyItemServiceMock }],

// BackButton: mockResolvedValue obrigatório para evitar NG04002
jest.spyOn(router, 'navigateByUrl').mockResolvedValue(true);

// BreadcrumbsComponent: forçar recompute do computed signal
extraSignal.set({ label: 'temp' });
extraSignal.set(null);
```

---

## Checklist de PR

- [ ] `npx tsc --noEmit` sem erros
- [ ] `npm test` — todos os testes passando (107+)
- [ ] Sidebar ativa item correto para nova rota
- [ ] Breadcrumb segue hierarquia correta
- [ ] Botão Voltar não aparece em rotas raiz
- [ ] Links externos usam `noopener,noreferrer`
- [ ] Grid responsivo testado visualmente
- [ ] Novas imagens em `public/assets/images/` com fallback `bannerColor`
