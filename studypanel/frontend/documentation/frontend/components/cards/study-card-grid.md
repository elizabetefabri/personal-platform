# Study Card Grid Component

## Status

Existente

## Caminho

`src/app/shared/components/study-card-grid/study-card-grid.ts`

## Selector

`app-study-card-grid`

## Descrição

Grid responsivo de cards de estudo. Cada card tem banner (imagem ou gradient), título, descrição e botão de ação.

## Inputs

```ts
@Input({ required: true }) pageTitle: string
@Input({ required: true }) pageDescription: string
@Input({ required: true }) items: StudyCardItem[]
```

## Interface StudyCardItem

```ts
interface StudyCardItem {
  id: string;
  title: string;
  description: string;
  bannerImage?: string;   // caminho relativo a public/
  bannerColor: string;    // fallback se imagem não carrega
  detailRoute: string;    // rota de destino do botão
}
```

## Estrutura do template

```html
<header class="studyPage__header">
  <div class="studyPage__topRow">
    <span class="studyPage__eyebrow">StudyPanel</span>
    <app-back-button></app-back-button>
  </div>
  <h1>{{ pageTitle }}</h1>
  <p>{{ pageDescription }}</p>
</header>
<div class="studyGrid">
  @for (item of items; track item.id) { <app-study-card [item]="item"> }
</div>
```

## Responsividade

1 → 2 → 3 → 4 → 6 colunas conforme breakpoints de 500px, 768px, 1024px, 1440px.

## Riscos de regressão

- BackButtonComponent deve estar nos `imports`
- RouterModule necessário nos testes (BackButton usa Router)
