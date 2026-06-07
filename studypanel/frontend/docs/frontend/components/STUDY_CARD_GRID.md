# Study Card Grid Component

## Informações Básicas

- **Selector:** `app-study-card-grid`
- **Tipo:** Standalone Component
- **Localização:** `src/app/shared/components/study-card-grid/`
- **Dependências:** Angular Common, Router, PrimeNG Button

## Descrição

Componente de grade responsivo que exibe cards de estudos/tópicos. Cada card contém uma imagem/cor de banner, título, descrição e botão de ação.

## Funcionalidades

- ✅ Grid responsiva
- ✅ Cores customizáveis para cada card
- ✅ Extração automática de cores do banner
- ✅ Links navegáveis
- ✅ Integração com PrimeNG Button

## Inputs

```typescript
@Input({ required: true }) pageTitle: string = ''
```

- Título principal da página

```typescript
@Input({ required: true }) pageDescription: string = ''
```

- Descrição/subtítulo da página

```typescript
@Input({ required: true }) items: StudyCardItem[] = []
```

- Array de cards a serem exibidos

## Interface de Dados

```typescript
interface StudyCardItem {
  id: string;
  title: string;
  description: string;
  bannerColor?: string; // Cor HEX do banner (ex: #8f48eb)
  href: string; // Link de navegação
  // ... outras propriedades
}
```

## Métodos Públicos

```typescript
getAccentColor(item: StudyCardItem): string
```

- Extrai cor HEX válida do `bannerColor`
- Retorna cor padrão (#4f46e5) se inválida ou não fornecida
- Suporta formato HEX completo: `#RRGGBB`

## Arquivo de Template

- `study-card-grid.html` - Layout da grade

## Arquivo de Estilos

- `study-card-grid.scss`
- Responsivo com breakpoints
- Usa CSS Grid para layout

## Exemplo de Uso

```typescript
import { StudyCardGrid } from '@shared/components/study-card-grid/study-card-grid';

@Component({
  imports: [StudyCardGrid],
})
export class MyPage {
  items: StudyCardItem[] = [
    {
      id: '1',
      title: 'Machine Learning',
      description: 'Algoritmos de ML...',
      bannerColor: '#8f48eb',
      href: '/artificial-intelligence/machine-learning',
    },
  ];

  pageTitle = 'Inteligência Artificial';
  pageDescription = 'Explore topicos de IA';
}
```

HTML:

```html
<app-study-card-grid
  [pageTitle]="pageTitle"
  [pageDescription]="pageDescription"
  [items]="items"
></app-study-card-grid>
```

## Integração com Banco de Dados

Para o futuro banco de dados, considerar tabela:

```sql
CREATE TABLE study_cards (
  id INT PRIMARY KEY AUTO_INCREMENT,
  section_id INT NOT NULL,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  banner_color VARCHAR(7) DEFAULT '#4f46e5',
  href VARCHAR(255),
  order INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (section_id) REFERENCES sections(id)
);
```

## Notas Técnicas

- Usa regex para validar e extrair cores HEX: `/^#[0-9a-fA-F]{6}$/`
- Input `items` é obrigatório e deve ser um array
- Cards são tipicamente renderizados em um loop `@for` no template
- Suporta navegação interna via RouterLink
