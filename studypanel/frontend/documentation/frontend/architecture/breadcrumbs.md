# Breadcrumbs

## Status

Existente

## Componente

`src/app/shared/components/breadcrumbs/breadcrumbs.component.ts`

## Hierarquia por rota

| Rota                       | Breadcrumbs                                          |
| -------------------------- | ---------------------------------------------------- |
| `/dashboard`               | — (sem breadcrumbs)                                  |
| `/estudos-labs`            | Dashboard > Estudos e Labs                           |
| `/backend`                 | Dashboard > Estudos e Labs > Backend                 |
| `/cloud`                   | Dashboard > Estudos e Labs > Cloud                   |
| `/projetos`                | Dashboard > Projetos                                 |
| `/projetos/pessoais`       | Dashboard > Projetos > Projetos Pessoais             |
| `/projetos/profissionais`  | Dashboard > Projetos > Projetos Profissionais        |
| `/rollout-service`         | Dashboard > Projetos > Rollout Service               |
| `/:section/:topic`         | Dashboard > Estudos e Labs > [Seção] > [Tópico]     |

## Como funciona

- `BreadcrumbsComponent` é global em `app.html`, acima do `<router-outlet>`
- `baseItems: BreadcrumbItem[]` é um array plain atualizado pelo método `build()` a cada NavigationEnd
- `items` é um `computed()` signal que combina `baseItems + extra` do `BreadcrumbService`
- `isHidden()` lê `BreadcrumbService.hidden()` para ocultar breadcrumbs em páginas específicas

## Riscos de regressão

- Adicionar nova rota sem atualizar `build()` gera breadcrumb incorreto
- Seções de estudo devem sempre ter "Estudos e Labs" como intermediário
- O computed `items` só re-avalia quando um signal de dependência muda
