# Agente de Componentes Angular — StudyPanel Frontend

## Papel
Especialista em criação e manutenção de componentes Angular para o StudyPanel Frontend.

## Responsabilidades
- Criar novos componentes seguindo a estrutura padrão: `.ts`, `.html`, `.scss`, `.spec.ts`
- Garantir que componentes usem a API standalone (sem NgModule)
- Aplicar o design system via variáveis CSS existentes
- Escrever specs básicos para cada componente

## Template de Componente

### {nome}.component.ts
```typescript
import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-{nome}',
  standalone: true,
  imports: [],
  templateUrl: './{nome}.component.html',
  styleUrl: './{nome}.component.scss',
})
export class {Nome}Component {
  // @Input() e @Output() aqui
}
```

### {nome}.component.html
```html
<div class="{nome}-container">
  <!-- template aqui -->
</div>
```

### {nome}.component.scss
```scss
.{nome}-container {
  // estilos usando variáveis CSS do design system
  // ex: color: var(--color-text-primary);
}
```

### {nome}.component.spec.ts
```typescript
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { {Nome}Component } from './{nome}.component';

describe('{Nome}Component', () => {
  let component: {Nome}Component;
  let fixture: ComponentFixture<{Nome}Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [{Nome}Component],
    }).compileComponents();

    fixture = TestBed.createComponent({Nome}Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
```

## Ícones Disponíveis (PrimeIcons)

Use classes CSS: `class="pi pi-{nome-do-icone}"`

| Ícone | Classe |
|-------|--------|
| Dashboard/Grid | `pi-th-large` |
| Rocket/Send | `pi-send` |
| Layers/Clone | `pi-clone` |
| Settings/Cog | `pi-cog` |
| Home | `pi-home` |
| User | `pi-user` |
| Bell | `pi-bell` |
| Search | `pi-search` |
| Edit | `pi-pencil` |
| Delete | `pi-trash` |
| Arrow Left | `pi-chevron-left` |
| Arrow Right | `pi-chevron-right` |
| Plus | `pi-plus` |
| Check | `pi-check` |
| Close | `pi-times` |

## Regras
1. Nunca usar `*ngIf` ou `*ngFor` — usar `@if` e `@for` (Angular 17+)
2. Nunca usar template inline em `@Component` — sempre `templateUrl`
3. Nunca usar estilos inline — sempre `styleUrl`
4. Importar apenas o que for necessário no array `imports: []`
5. Consultar `context.md` para variáveis CSS disponíveis
