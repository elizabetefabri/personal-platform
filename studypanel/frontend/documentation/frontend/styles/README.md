# Estilos

## Custom Properties (Design Tokens)

Definidas em `src/styles/abstracts/_variables.scss` ou no `:root` de `src/styles.scss`:

```css
:root {
  /* Layout */
  --sidebar-width-expanded:  220px;
  --sidebar-width-collapsed: 64px;
  --header-height:           64px;

  /* Tipografia */
  --text-xs:   0.75rem;  /* 12px */
  --text-sm:   0.875rem; /* 14px */
  --text-base: 1rem;     /* 16px */
  --text-lg:   1.125rem; /* 18px */

  /* Ícone */
  --icon-size: 1.25rem;
}
```

## Metodologia

- **BEM**: `.componentName__element--modifier`
- `.app-sidebar__item--active`, `.studyCard__title`, `.detailHeader__topRow`

## Convenções

- Nunca usar `px` para font-size — usar `rem` ou `clamp()`
- `clamp(min, preferred, max)` para tipografia fluida
- Transitions: `transition: all 0.3s ease` padrão para sidebar/collapse
- Z-indexes: sidebar `10`, header `5`, overlay `100`

## Breakpoints padrão

| Breakpoint   | Valor rem | Valor px |
| ------------ | --------- | -------- |
| `--bp-xs`    | 31.25rem  | 500px    |
| `--bp-sm`    | 48rem     | 768px    |
| `--bp-md`    | 64rem     | 1024px   |
| `--bp-lg`    | 90rem     | 1440px   |
