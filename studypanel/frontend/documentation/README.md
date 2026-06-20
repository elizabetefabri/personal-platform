# Documentação do Frontend — StudyPanel

## Objetivo

Centralizar regras, estruturas e decisões técnicas do frontend Angular do StudyPanel.

## Índice

- [Arquitetura](./frontend/architecture/estrutura-do-projeto.md)
- [Rotas](./frontend/architecture/rotas.md)
- [Breadcrumbs](./frontend/architecture/breadcrumbs.md)
- [Botão Voltar](./frontend/architecture/botao-voltar.md)
- [Responsividade](./frontend/architecture/responsividade.md)
- [SEO](./frontend/architecture/seo.md)
- [Páginas](./frontend/pages/README.md)
- [Componentes](./frontend/components/README.md)
- [Serviços](./frontend/services/README.md)
- [Models / Interfaces](./frontend/models/README.md)
- [Dados](./frontend/data/README.md)
- [Estilos](./frontend/styles/README.md)
- [Assets](./frontend/assets/README.md)
- [Testes](./frontend/tests/README.md)
- [Manual de Comandos Frontend](./frontend/manual/Comandos-Regras-Frontend.md)
- [Manual de Comandos Backend](./frontend/manual/Comandos-Regras-Backend.md)

## Como manter esta documentação

1. Toda nova página deve ter um arquivo em `pages/`.
2. Todo novo componente deve ter um arquivo em `components/`.
3. Serviços e models devem ser documentados quando criados ou alterados.
4. Não documentar comportamento que não existe no código. Marcar como `Status: planejado` quando necessário.

## Checklist de prevenção de regressão

- [ ] Sidebar ativa o item correto para cada rota.
- [ ] Breadcrumbs seguem a hierarquia correta.
- [ ] Botão Voltar não aparece em páginas raiz.
- [ ] Grids são responsivos (1 → 2 → 3 → 4 → 6 colunas).
- [ ] Imagens dos cards têm fallback para `bannerColor`.
- [ ] Links externos abrem em `_blank` com `noopener,noreferrer`.
- [ ] Testes unitários passam (107 testes).
- [ ] Frontend compila sem erros.
