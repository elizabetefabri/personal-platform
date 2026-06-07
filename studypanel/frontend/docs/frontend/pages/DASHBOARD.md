# Dashboard Page

## Informações Básicas

- **Selector:** `app-dashboard`
- **Tipo:** Standalone Component
- **Route:** `/dashboard`
- **Localização:** `src/app/pages/dashboard/`

## Descrição

Página inicial da aplicação que exibe um overview de todas as categorias de estudo disponíveis.

## Responsabilidades

- Exibir estatísticas ou resumo (futuro)
- Navegação rápida para todas as seções
- Ponto de entrada da aplicação

## Estrutura

```typescript
export class DashboardComponent {
  pageTitle = 'Dashboard';
  pageDescription = 'Selecione uma categoria para começar';
  items: StudyCardItem[] = [...];
}
```

## Dados Exibidos

O Dashboard exibe cards para cada seção principal:

- Backend
- Banco de Dados
- Cloud Computing
- Containers & Kubernetes
- DevOps
- Frontend
- Inteligência Artificial
- Observability
- Performance Engineering
- Projetos
- Rollout Service

## Integração com Banco de Dados

```sql
CREATE TABLE dashboard_stats (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT,
  total_studies INT DEFAULT 0,
  completed_studies INT DEFAULT 0,
  last_visited TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Notas Técnicas

- Primeira página renderizada após login
- Sem breadcrumbs (exibir vazio)
- Base para futuras expansões (widgets, stats, etc.)
