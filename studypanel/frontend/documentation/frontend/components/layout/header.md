# Header Component

## Status

Existente

## Caminho

`src/app/shared/components/header/header.ts`

## Selector

`app-header`

## Descrição

Barra superior global. Exibe título da página (via `BreadcrumbService.pageTitle()`) e inclui `app-breadcrumbs`.

## Dependências

- `BreadcrumbService` — lê `pageTitle()` signal
- `BreadcrumbsComponent` — incluso no template

## Propriedades relevantes

```ts
pageTitle = this.breadcrumbService.pageTitle; // Signal<string>
```

## Riscos de regressão

- `HeaderComponent` NÃO tem propriedade `title` como @Input — título vem do serviço
