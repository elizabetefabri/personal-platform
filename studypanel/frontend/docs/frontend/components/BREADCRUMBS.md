# Breadcrumbs Component

## Informações Básicas

- **Selector:** `app-breadcrumbs`
- **Tipo:** Standalone Component
- **Localização:** `src/app/shared/components/breadcrumbs/`
- **Dependências:** Angular Router

## Descrição

Componente de navegação por breadcrumbs que exibe o caminho hierárquico da página atual. Atualiza automaticamente ao navegar entre rotas.

## Funcionalidades

- ✅ Breadcrumbs inteligentes baseadas em rotas
- ✅ Geração automática de labels
- ✅ Links navegáveis
- ✅ Suporte a seções e tópicos customizados

## Interfaces

```typescript
interface BreadcrumbItem {
  label: string;
  href?: string; // Opcional, sem href significa item final
}
```

## Propriedades

```typescript
items: BreadcrumbItem[] = []
```

## Lógica de Geração

### Regras

1. Dashboard → Sem breadcrumbs (array vazio)
2. Rota de um nível → `Dashboard > Seção`
3. Rota de dois níveis → `Dashboard > Seção > Tópico`

### Nomes Customizados

- Utiliza `PAGE_LABELS` para páginas especiais (ex: "settings" → "Configurações")
- Busca labels em `study-topics.constants` para seções e tópicos
- Converte slugs em labels legíveis (ex: "banco-de-dados" → "Banco de Dados")

## Métodos Privados

```typescript
private build(rawUrl: string): void
```

- Constrói o array de breadcrumbs baseado na URL

```typescript
private toLabel(slug: string): string
```

- Converte slug para label legível (title case)

## Exemplo de Uso

```typescript
// Em app.component.ts
<app-breadcrumbs></app-breadcrumbs>
```

## Integração com Banco de Dados

Para o futuro banco de dados, considerar tabelas:

```sql
CREATE TABLE sections (
  id INT PRIMARY KEY AUTO_INCREMENT,
  slug VARCHAR(100) UNIQUE NOT NULL,
  label VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE topics (
  id INT PRIMARY KEY AUTO_INCREMENT,
  section_id INT NOT NULL,
  slug VARCHAR(100) NOT NULL,
  label VARCHAR(100) NOT NULL,
  FOREIGN KEY (section_id) REFERENCES sections(id),
  UNIQUE KEY (section_id, slug),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Notas Técnicas

- Implementa OnInit e OnDestroy para gerenciamento de subscrições
- Desinscreve de observables para evitar memory leaks
- Atualiza breadcrumbs em tempo real ao navegar
- Suporta parâmetros de query na URL (ignora `?` e tudo após)
