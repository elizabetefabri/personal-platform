# Pages Documentation

## Visão Geral

As páginas do StudyPanel são organizadas em duas categorias principais:

1. **Páginas de Seção** - Exibem cards de tópicos em uma categoria
2. **Páginas de Detalhe** - Exibem conteúdo detalhado de um tópico específico

## Páginas de Seção

Todas as páginas de seção seguem o mesmo padrão:

- Exibem um título e descrição da seção
- Usam `StudyCardGrid` para exibir cards de tópicos
- Implementam breadcrumbs automáticos
- Conectam-se ao sidebar para navegação

### Páginas Disponíveis

| Página                  | Route                      | Arquivo                                |
| ----------------------- | -------------------------- | -------------------------------------- |
| Dashboard               | `/dashboard`               | `dashboard.component.ts`               |
| Backend                 | `/backend`                 | `backend.component.ts`                 |
| Banco de Dados          | `/banco-de-dados`          | `banco-de-dados.component.ts`          |
| Cloud Computing         | `/cloud`                   | `cloud.component.ts`                   |
| Containers & Kubernetes | `/containers-kubernetes`   | `containers-kubernetes.component.ts`   |
| DevOps                  | `/devops`                  | `devops.component.ts`                  |
| Frontend                | `/frontend`                | `frontend.component.ts`                |
| Inteligência Artificial | `/inteligencia-artificial` | `inteligencia-artificial.component.ts` |
| Observability           | `/observability`           | `observability.component.ts`           |
| Performance Engineering | `/performance-engineering` | `performance-engineering.component.ts` |
| Projetos                | `/projetos`                | `projetos.component.ts`                |
| Rollout Service         | `/rollout-service`         | `rollout-service.component.ts`         |

## Páginas Especiais

### Settings

- **Route:** `/settings`
- **Arquivo:** `settings.component.ts`
- Página de configurações da aplicação

### Study Detail Page

- **Route:** (dinâmica, baseada em navegação)
- **Arquivo:** `study-detail-page.component.ts`
- Página que exibe detalhes de um tópico específico

## Padrão de Implementação

### Template Típico

```typescript
@Component({
  selector: 'app-section-page',
  standalone: true,
  imports: [StudyCardGrid],
  templateUrl: './section.component.html',
  styleUrl: './section.component.scss',
})
export class SectionComponent {
  pageTitle = 'Título da Seção';
  pageDescription = 'Descrição da seção';
  items: StudyCardItem[] = [
    // ... items
  ];
}
```

### Template HTML

```html
<app-study-card-grid
  [pageTitle]="pageTitle"
  [pageDescription]="pageDescription"
  [items]="items"
></app-study-card-grid>
```

## Estrutura de Dados

As páginas de seção utilizam dados de estudo definidos em:

- `src/app/core/constants/study-topics.ts` - Definições de tópicos

Cada item contém:

- `id`: Identificador único
- `title`: Título do tópico
- `description`: Breve descrição
- `bannerColor`: Cor do banner (HEX)
- `href`: Link de navegação

## Integração com Banco de Dados

Para o futuro, as páginas podem ser dinamicamente carregadas de:

```sql
CREATE TABLE page_sections (
  id INT PRIMARY KEY AUTO_INCREMENT,
  slug VARCHAR(100) UNIQUE NOT NULL,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  route VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE page_content (
  id INT PRIMARY KEY AUTO_INCREMENT,
  section_id INT NOT NULL,
  title VARCHAR(100),
  description TEXT,
  banner_color VARCHAR(7),
  content TEXT,
  FOREIGN KEY (section_id) REFERENCES page_sections(id)
);
```

## Notas sobre Roteamento

- Todas as páginas utilizam Angular Router
- Breadcrumbs são gerados automaticamente pelo `BreadcrumbsComponent`
- Sidebar atualiza item ativo automaticamente baseado na rota
- URLs usam slugs em português (ex: `/inteligencia-artificial`)
