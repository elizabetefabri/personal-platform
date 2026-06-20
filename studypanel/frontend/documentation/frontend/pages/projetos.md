# Projetos (Landing)

## Status

Existente

## Caminho

`src/app/pages/projetos/projetos.component.ts`

## Rota

`/projetos`

## Responsabilidade

Página de entrada para projetos. Exibe 2 cards de navegação: Projetos Pessoais e Projetos Profissionais.

## Breadcrumbs

Dashboard > Projetos

## Componentes usados

- `StudyCardGrid` — renderiza os 2 cards de categoria

## Cards

| Card                  | Rota de destino          | Botão              |
| --------------------- | ------------------------ | ------------------ |
| Projetos Pessoais     | `/projetos/pessoais`     | Abrir Projetos     |
| Projetos Profissionais | `/projetos/profissionais` | Abrir Projetos     |

## Regras de layout

- Botão Voltar NÃO aparece nesta página (rota raiz)
- Grid: 1 → 2 colunas

## Riscos de regressão

- `detailRoute` deve apontar para sub-rotas corretas de `/projetos`
