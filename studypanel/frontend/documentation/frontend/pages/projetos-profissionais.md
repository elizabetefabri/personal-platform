# Projetos Profissionais

## Status

Existente

## Caminho

`src/app/pages/projetos-profissionais/projetos-profissionais.component.ts`

## Rota

`/projetos/profissionais`

## Responsabilidade

Exibe cards dos projetos profissionais com links para detalhes internos.

## Breadcrumbs

Dashboard > Projetos > Projetos Profissionais

## Componentes usados

- `ProjectCardGrid` — renderiza os cards de projetos

## Projetos incluídos

- **Rollout Service** — Feature flags graduais (`detailRoute: '/rollout-service'`)
- **IUDev** — Plataforma de UI/UX

## Botão Voltar

Aparece inline, navega para `/projetos`.

## Riscos de regressão

- Projetos com `detailRoute` devem ter a rota configurada em `app.routes.ts`
- Não confundir com projetos pessoais (têm botões de repo/deploy)
