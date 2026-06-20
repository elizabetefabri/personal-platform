# Project Card Grid Component

## Status

Existente

## Caminho

`src/app/shared/components/project-card-grid/project-card-grid.ts`

## Selector

`app-project-card-grid`

## Descrição

Grid de cards para projetos pessoais e profissionais. Diferente de `StudyCardGrid` por suportar botões de repositório e deploy.

## Inputs

```ts
@Input({ required: true }) pageTitle: string
@Input({ required: true }) pageDescription: string
@Input({ required: true }) items: ProjectCardItem[]
```

## Interface ProjectCardItem

```ts
interface ProjectCardItem {
  id: string;
  title: string;
  description: string;
  bannerImage?: string;
  bannerColor: string;
  tags?: string[];         // tecnologias usadas
  repoUrl?: string;        // link GitHub (abre em _blank)
  deployUrl?: string;      // link de deploy (abre em _blank)
  detailRoute?: string;    // rota interna (alternativa a repoUrl/deployUrl)
}
```

## Diferenças em relação a StudyCardGrid

- Suporte a `tags` (chips de tecnologia)
- Botões externos (repo/deploy) em vez de `[routerLink]`
- Links externos sempre com `noopener,noreferrer`
