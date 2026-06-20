# Responsividade

## Status

Existente

## Grid de Cards

Padrão aplicado em `StudyCardGrid` e `ProjectCardGrid`:

```scss
.studyGrid { grid-template-columns: 1fr; }            /* < 500px — 1 coluna */

@media (min-width: 31.25rem) { /* 500px */
  .studyGrid { grid-template-columns: repeat(2, 1fr); }
}
@media (min-width: 48rem) { /* 768px */
  .studyGrid { grid-template-columns: repeat(3, 1fr); }
}
@media (min-width: 64rem) { /* 1024px */
  .studyGrid { grid-template-columns: repeat(4, 1fr); }
}
@media (min-width: 90rem) { /* 1440px */
  .studyGrid { grid-template-columns: repeat(6, 1fr); }
}
```

## Tipografia responsiva

Fontes usam `clamp()` para transições suaves:

```scss
.studyPage__title   { font-size: clamp(2rem, 4vw, 2.5rem); }
.studyCard__title   { font-size: clamp(1rem, 1.2vw, 1.15rem); }
.studyCard__description { font-size: clamp(0.875rem, 1vw, 0.9375rem); }
```

## Sidebar

- Largura expandida: `var(--sidebar-width-expanded)` — 220px
- Largura recolhida: `var(--sidebar-width-collapsed)` — 64px
- O layout principal (`app-content`) usa `padding-left` dinâmico via classe `.sidebar-expanded`

## Regras gerais

- Usar `rem` e `%` em vez de `px` fixo
- `clamp(min, preferred, max)` para valores fluidos
- `minmax()` dentro de grids
- Transições suaves: sem "pulo" brusco de layout
