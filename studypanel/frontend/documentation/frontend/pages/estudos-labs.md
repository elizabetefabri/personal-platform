# Estudos e Labs

## Status

Existente

## Caminho

`src/app/pages/estudos-labs/estudos-labs.component.ts`

## Rota

`/estudos-labs`

## Responsabilidade

Página de listagem das 9 seções de estudo técnico disponíveis. Serve como hub de entrada para cada categoria.

## Breadcrumbs

Dashboard > Estudos e Labs

## Componentes usados

- `StudyCardGrid` — renderiza os 9 cards

## Dados

Gerado a partir de `STUDY_SECTIONS` (filtrado pelos 9 slugs de estudo) em `src/app/core/constants/study-topics.ts`.

Slugs: `backend`, `banco-de-dados`, `cloud`, `containers-kubernetes`, `devops`, `frontend`, `inteligencia-artificial`, `observability`, `performance-engineering`.

## Estrutura visual

- Eyebrow "StudyPanel" + botão Voltar (oculto pois é rota raiz)
- Título "Estudos e Labs"
- Grid de 9 cards com bannerColor gradient, ícone e botão

## Regras de layout

- Grid: 1 → 2 → 3 → 4 → 6 colunas conforme breakpoints
- Botão Voltar NÃO aparece nesta página (rota raiz)

## Riscos de regressão

- Alterar slugs em `ESTUDOS_SLUGS` sem atualizar `ROUTE_MAP` pode gerar rotas quebradas
