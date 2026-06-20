# Botão Voltar

## Status

Existente

## Componente

`src/app/shared/components/back-button/back-button.ts`

## Regras

- **Aparece inline** dentro do header de cada página, ao lado do eyebrow text
- **NÃO aparece** em páginas raiz: `/dashboard`, `/estudos-labs`, `/projetos`
- **NÃO aparece** em `course-detail` (já tem "Voltar para trilha" próprio)
- Visibilidade controlada pelo computed `visible()` baseado na URL atual

## Onde está incluído

| Componente              | Posição                        |
| ----------------------- | ------------------------------ |
| `StudyCardGrid`         | `.studyPage__topRow` (ao lado do eyebrow "StudyPanel") |
| `ProjectCardGrid`       | `.projectPage__topRow` (ao lado do eyebrow "StudyPanel") |
| `StudyDetailTemplate`   | `.detailHeader__topRow` (ao lado do eyebrow "Detalhes do estudo") |

## Lógica de navegação `goBack()`

```
URL com > 1 segmento → navega para rota pai (/backend/node → /backend)
URL com 1 segmento   → chama location.back()
Query strings são ignoradas no cálculo
```

## Riscos de regressão

- Não adicionar a `course-detail` (já tem botão próprio)
- Não tornar `position: fixed` — deve ser inline para não sobrepor conteúdo
- O computed `visible()` usa SET de rotas raiz, atualizar se novas raízes forem criadas
