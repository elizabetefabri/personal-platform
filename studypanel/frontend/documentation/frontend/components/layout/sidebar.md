# Sidebar Component

## Status

Existente

## Caminho

`src/app/shared/components/sidebar/sidebar.ts`

## Selector

`app-sidebar`

## Descrição

Navegação lateral colapsável. Detecta rota ativa automaticamente e exibe 3 itens de menu + logout.

## Itens de Menu

| Key           | Label          | Rota                | Ícone         |
| ------------- | -------------- | ------------------- | ------------- |
| `dashboard`   | Dashboard      | `/dashboard`        | `pi-th-large` |
| `estudos-labs`| Estudos e Labs | `/estudos-labs`     | `pi-book`     |
| `projetos`    | Projetos       | `/projetos`         | `pi-folder`   |

Botão extra: **Sair** (`pi-sign-out`) — apenas visual no StudyPanel (sem autenticação real).

## Inputs / Outputs

```ts
@Input() expanded: boolean = true
@Output() toggleExpanded = new EventEmitter<void>()
```

## Active state

A rota ativa é detectada automaticamente: se `router.url.startsWith(item.href)`, o item fica ativo.

- `/backend`, `/cloud`, etc. ativam `estudos-labs`
- `/projetos/pessoais` ativa `projetos`

## Risco de regressão

- Alterar o número de itens sem atualizar `sidebar.spec.ts` (espera `toBe(3)`)
- Adicionar nova seção sem mapear para o item correto da sidebar
