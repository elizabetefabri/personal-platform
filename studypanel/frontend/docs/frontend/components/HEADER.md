# Header Component

## Informações Básicas

- **Selector:** `app-header`
- **Tipo:** Standalone Component
- **Localização:** `src/app/shared/components/header/`
- **Dependências:** Angular Router

## Descrição

Componente de cabeçalho que exibe o título da aplicação (StudyPanel) e informações de navegação no topo da página.

## Funcionalidades

- ✅ Exibição do título da aplicação
- ✅ Suporte a navegação
- ✅ Integração com breadcrumbs

## Propriedades

```typescript
title = 'StudyPanel';
```

## Arquivo de Template

- `header.component.html` - Contém o layout do header

## Arquivo de Estilos

- `header.component.scss`
- **Altura:** 4rem (conforme --header-height)
- **Fundo:** `--color-header-bg`

## Exemplo de Uso

```typescript
// Em app.component.ts
<app-header></app-header>
```

## Integração com Banco de Dados

Para o futuro banco de dados, considerar:

```sql
CREATE TABLE app_metadata (
  id INT PRIMARY KEY AUTO_INCREMENT,
  app_title VARCHAR(100),
  app_version VARCHAR(50),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Notas Técnicas

- Componente simples sem lógica complexa
- Pode ser expandido para incluir notificações, perfil de usuário, etc.
- Altura fixa para manter consistência com o layout geral
