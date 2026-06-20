# Dashboard

## Status

Existente

## Caminho

`src/app/pages/dashboard/dashboard.component.ts`

## Rota

`/dashboard`

## Responsabilidade

Página inicial da aplicação. Exibe overview de todas as seções de estudo com estatísticas, distribuição de status e progresso recente.

## Breadcrumbs

Nenhum (página raiz).

## Componentes usados

- `ChartModule` (PrimeNG) — gráficos de status e evolução
- `ButtonModule`, `TagModule`, `DialogModule` — UI
- `StudyItemService` — carrega todos os itens para estatísticas

## Estrutura visual

- Cabeçalho com nome de boas-vindas e resumo geral
- Seção "Seções de Estudo" — lista das categorias com atalhos
- Card de gráfico — distribuição por status (donut chart)
- Card de gráfico — evolução mensal (line chart)

## Regras de layout

- Sem botão Voltar (rota raiz)
- Layout em grid responsivo

## Riscos de regressão

- Se `StudyItemService.listAll()` falhar silenciosamente, os gráficos ficam vazios
- Mockar o serviço em testes unitários (`listAll: jest.fn(() => of([]))`)
