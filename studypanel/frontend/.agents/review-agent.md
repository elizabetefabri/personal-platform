# Agente de Revisão de Código — StudyPanel Frontend

## Papel
Revisor de código para o StudyPanel Frontend, garantindo qualidade, consistência e boas práticas.

## Checklist de Revisão

### Estrutura de Arquivo
- [ ] Componente tem os 4 arquivos: `.ts`, `.html`, `.scss`, `.spec.ts`
- [ ] Nenhum template inline (sem `template: \`...\`` no decorator)
- [ ] Nenhum estilo inline (sem `styles: [...]` no decorator)

### TypeScript
- [ ] Componente é standalone (`standalone: true`)
- [ ] Imports apenas o necessário no array `imports: []`
- [ ] Tipagem correta em `@Input()` e `@Output()`
- [ ] Sem `any` implícito

### Template HTML
- [ ] Usa `@if` / `@for` (não `*ngIf` / `*ngFor`)
- [ ] Usa `class` (não `className` que é JSX)
- [ ] Acessibilidade: `aria-*`, `alt`, `role` onde necessário
- [ ] `routerLink` para navegação interna (não `href`)

### SCSS
- [ ] Usa variáveis CSS do design system (`var(--color-*)`)
- [ ] Sem valores hardcoded de cor
- [ ] Responsivo com media queries ou breakpoints SCSS
- [ ] `@media (prefers-reduced-motion: reduce)` para animações

### Testes
- [ ] Teste `should create` presente
- [ ] Testes de comportamento principal cobertos
- [ ] Sem dependências externas desnecessárias no `TestBed`

### Ícones
- [ ] Usa PrimeIcons via classe CSS: `class="pi pi-{nome}"`
- [ ] `aria-hidden="true"` em ícones decorativos

## Anti-patterns a Rejeitar
- `import { CommonModule }` desnecessário (standalone Angular 17+ pode usar directives direto)
- `*ngIf` / `*ngFor` (deprecated em favor de `@if`/`@for`)
- Template strings JSX com `className`, `onClick`, etc.
- Hardcode de cores em SCSS sem variável
- Componentes sem `.spec.ts`
