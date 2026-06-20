# Back Button Component

## Status

Existente

## Caminho

`src/app/shared/components/back-button/back-button.ts`

## Selector

`app-back-button`

## Descrição

Botão Voltar inline, posicionado ao lado do eyebrow text no header de cada página secundária.

## Lógica de visibilidade

```ts
private readonly ROOT_ROUTES = new Set(['/dashboard', '/estudos-labs', '/projetos']);
visible = computed(() => !this.ROOT_ROUTES.has(this.currentUrl()));
```

## Lógica de navegação `goBack()`

```ts
// URL: /backend/node → navega para /backend
// URL: /backend     → chama location.back()
```

## Onde aparece

| Componente host        | Container CSS                 |
| ---------------------- | ----------------------------- |
| `StudyCardGrid`        | `.studyPage__topRow`          |
| `ProjectCardGrid`      | `.projectPage__topRow`        |
| `StudyDetailTemplate`  | `.detailHeader__topRow`       |

## Padrão de layout

```scss
.studyPage__topRow {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 0.25rem;
}
```

## Pattern de teste

```ts
jest.spyOn(router, 'navigateByUrl').mockResolvedValue(true);
// .mockResolvedValue(true) é obrigatório para prevenir crash com NG04002
```

## NÃO incluir em

- `course-detail` (já tem "Voltar para trilha" próprio)
- Qualquer página com `position: fixed` (causa sobreposição)
