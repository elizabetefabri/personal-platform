# Study Detail Template Component

## Status

Existente

## Caminho

`src/app/shared/components/study-detail-template/study-detail-template.ts`

## Selector

`app-study-detail`

## Descrição

Template para a página de detalhe de um curso/tópico específico. Exibe conteúdo estruturado com accordion de módulos e lista de tópicos.

## Inputs

```ts
@Input({ required: true }) pageTitle: string
@Input({ required: true }) section: string      // slug da seção (ex: 'backend')
@Input({ required: true }) topic: string        // slug do tópico
```

## Estrutura do template

```html
<header class="detailHeader">
  <div class="detailHeader__topRow">
    <span class="detailHeader__eyebrow">Detalhes do estudo</span>
    <app-back-button></app-back-button>
  </div>
  <h1>{{ pageTitle }}</h1>
</header>
```

## Dependências

- `StudyItemService` — carrega itens do tópico
- `BackButtonComponent` — no header inline

## Pattern de teste

```ts
const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  listAll: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};
// + RouterModule.forRoot([]) obrigatório (BackButton usa Router)
```

## Riscos de regressão

- Esquecer `StudyItemService` mock nos testes → erro de injeção
- Esquecer `RouterModule.forRoot([])` → NG04002 (BackButton crash)
