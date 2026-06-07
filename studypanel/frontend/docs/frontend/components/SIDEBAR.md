# Sidebar Component

## Informações Básicas

- **Selector:** `app-sidebar`
- **Tipo:** Standalone Component
- **Localização:** `src/app/shared/components/sidebar/`
- **Dependências:** Angular Common, Router, PrimeIcons

## Descrição

Componente de navegação lateral (sidebar) colapsável que permite ao usuário navegar entre as diferentes seções da aplicação. O sidebar é responsável por mostrar/ocultar itens de menu e atualizar o item ativo baseado na rota atual.

## Funcionalidades

- ✅ Navegação colapsável (expandido/recolhido)
- ✅ Detecção automática de rota ativa
- ✅ Suporte a ícones (PrimeIcons)
- ✅ Botão de logout
- ✅ Feedback visual para itens ativos
- ✅ Transições suaves entre estados

## Inputs

```typescript
@Input() expanded: boolean = true
```

- Controla se o sidebar está expandido ou recolhido

## Outputs

```typescript
@Output() toggleExpanded = new EventEmitter<void>()
```

- Emitido quando o usuário clica no botão de toggle

## Itens de Navegação

| Key                     | Label                   | Href                     | Ícone         |
| ----------------------- | ----------------------- | ------------------------ | ------------- |
| dashboard               | Dashboard               | /dashboard               | pi-th-large   |
| backend                 | Backend                 | /backend                 | pi-server     |
| database                | Banco de Dados          | /banco-de-dados          | pi-database   |
| cloud                   | Cloud Computing         | /cloud                   | pi-cloud      |
| containers-kubernetes   | Containers e Kubernetes | /containers-kubernetes   | pi-box        |
| devops                  | DevOps                  | /devops                  | pi-cog        |
| frontend                | Frontend                | /frontend                | pi-desktop    |
| artificial-intelligence | Inteligência Artificial | /inteligencia-artificial | pi-sparkles   |
| observability           | Observability           | /observability           | pi-eye        |
| performance-engineering | Performance Engineering | /performance-engineering | pi-chart-line |
| projects                | Projetos                | /projetos                | pi-folder     |
| rollout-service         | Rollout Service         | /rollout-service         | pi-clone      |

## Métodos Públicos

```typescript
signOut(): void
```

- Atualmente apenas loga no console
- Deve ser integrado com serviço de autenticação

## Estilos

- **Arquivo:** `sidebar.component.scss`
- **Largura expandida:** 15rem
- **Largura colapsada:** 4rem
- **Altura:** 100dvh (viewport height dinâmica)
- **Fundo:** Tema escuro (`--color-sidebar-bg-dark`)
- **Transição:** 0.3s ease

### Estados CSS

- `.navItemActive` - Item de menu ativo
- `.navItemExpanded` - Estilo quando expandido
- `.navItemCollapsed` - Estilo quando recolhido

## Exemplo de Uso

```typescript
// Em app.component.ts
<app-sidebar
  [expanded]="sidebarExpanded"
  (toggleExpanded)="sidebarExpanded = !sidebarExpanded"
></app-sidebar>
```

## Integração com Banco de Dados

Para o futuro banco de dados, considerar tabela:

```sql
CREATE TABLE sidebar_items (
  id INT PRIMARY KEY AUTO_INCREMENT,
  key VARCHAR(50) UNIQUE NOT NULL,
  label VARCHAR(100) NOT NULL,
  href VARCHAR(255) NOT NULL,
  iconClass VARCHAR(100) NOT NULL,
  order INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Notas Técnicas

- Usa Router para detectar mudanças de rota
- Implementa OnInit e OnDestroy para gerenciamento de subscrições
- Desinscreve de observables no ngOnDestroy para evitar memory leaks
- Suporta acesso por teclado (aria labels)
