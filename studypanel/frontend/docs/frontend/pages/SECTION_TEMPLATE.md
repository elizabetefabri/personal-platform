# Section Page Template

## Padrão para Páginas de Seção

Use este template como base para documentar outras páginas de seção (Backend, Frontend, Cloud, etc).

### Informações Básicas
- **Selector:** `app-section-name`
- **Tipo:** Standalone Component
- **Route:** `/section-route` (ex: `/backend`, `/frontend`)
- **Localização:** `src/app/pages/section-name/`

### Descrição
[Breve descrição da seção e seu propósito no programa de estudos]

### Funcionalidades
- ✅ Exibe cards de tópicos relacionados
- ✅ Integração com breadcrumbs automáticos
- ✅ Navegação para detalhes de cada tópico

### Estrutura

```typescript
export class SectionComponent {
  pageTitle = '[Título da Seção]';
  pageDescription = '[Descrição breve]';
  items: StudyCardItem[] = [
    // Cards de tópicos
  ];
}
```

### Dados Exibidos
Lista de tópicos/cards disponíveis nesta seção.

### Integração com Banco de Dados

```sql
CREATE TABLE section_pages (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  route VARCHAR(255) UNIQUE NOT NULL,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Notas Técnicas
- Utiliza `StudyCardGrid` para exibição dos cards
- Breadcrumbs exibem: Dashboard > [Seção]
- Sidebar marca este item como ativo ao navegar
- Cores dos cards são customizáveis por tópico

### Exemplo de Estrutura do Código

```typescript
import { Component } from '@angular/core';
import { StudyCardGrid } from '@shared/components/study-card-grid/study-card-grid';
import { StudyCardItem } from '@shared/interfaces/study-template.interface';

@Component({
  selector: 'app-section',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './section.component.html',
  styleUrl: './section.component.scss',
})
export class SectionComponent {
  pageTitle = 'Section Title';
  pageDescription = 'Section description';
  
  items: StudyCardItem[] = [
    {
      id: '1',
      title: 'Topic 1',
      description: 'Description...',
      bannerColor: '#8f48eb',
      href: '/section/topic-1'
    },
    // ... mais topics
  ];
}
```

