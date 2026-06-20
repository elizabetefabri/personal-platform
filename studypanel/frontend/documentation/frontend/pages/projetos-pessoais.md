# Projetos Pessoais

## Status

Existente

## Caminho

`src/app/pages/projetos-pessoais/projetos-pessoais.component.ts`

## Rota

`/projetos/pessoais`

## Responsabilidade

Exibe cards dos projetos pessoais de Elizabete Fabri com links para repositório e deploy.

## Breadcrumbs

Dashboard > Projetos > Projetos Pessoais

## Componentes usados

- `ProjectCardGrid` — renderiza os cards de projetos

## Projetos incluídos

- Comanda Flow (Angular, deploy)
- StudyPanel (Angular)
- Personal Platform (Angular)
- E outros projetos pessoais

## Estrutura visual do card

- Banner com imagem ou gradient de cor + ícone
- Título, descrição, tags de tecnologia
- Botões: Repositório e/ou Deploy (links externos, `_blank`)

## Botão Voltar

Aparece inline, navega para `/projetos`.

## Regras de layout

- Grid: 1 → 2 → 3 → 4 → 6 colunas
- Imagens com fallback automático para `bannerColor`
- Links externos: `window.open(url, '_blank', 'noopener,noreferrer')`

## Riscos de regressão

- Não usar `[href]` para navegação interna — usar `[routerLink]`
- Links externos devem abrir em nova aba
