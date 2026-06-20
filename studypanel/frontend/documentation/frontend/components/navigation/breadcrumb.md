# Breadcrumbs Component

## Status

Existente

## Caminho

`src/app/shared/components/breadcrumbs/breadcrumbs.component.ts`

## Selector

`app-breadcrumbs`

## Descrição

Exibe o caminho hierárquico da página atual. Atualizado automaticamente a cada NavigationEnd.

## Interfaces

```ts
interface BreadcrumbItem {
  label: string;
  href?: string; // sem href = item final (sem link)
}
```

## Propriedades

```ts
baseItems: BreadcrumbItem[] = []                     // array plain (não signal)
items: Signal<BreadcrumbItem[]>                      // computed: baseItems + extra
isHidden: Signal<boolean>                            // lê BreadcrumbService.hidden()
```

## BreadcrumbService API

```ts
service.set(extra: BreadcrumbItem | null)            // item extra final
service.setHidden(hidden: boolean)                   // ocultar breadcrumbs
service.setPageTitle(title: string)                  // título da página (Header)
service.clear()                                      // resetar
```

## Hierarquia por rota

| Rota                      | Breadcrumbs                                    |
| ------------------------- | ---------------------------------------------- |
| `/dashboard`              | — (hidden)                                     |
| `/estudos-labs`           | Dashboard > Estudos e Labs                     |
| `/backend`                | Dashboard > Estudos e Labs > Backend           |
| `/projetos`               | Dashboard > Projetos                           |
| `/projetos/pessoais`      | Dashboard > Projetos > Projetos Pessoais       |
| `/projetos/profissionais` | Dashboard > Projetos > Projetos Profissionais  |

## Pattern de teste

```ts
// computed signal só re-avalia quando signal de dependência muda
// baseItems é array plain, então é preciso toggle em extra
extraSignal.set({ label: 'temp' });
extraSignal.set(null);
expect(component.items()).toEqual(base);
```
