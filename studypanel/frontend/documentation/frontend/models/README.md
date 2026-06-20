# Models e Interfaces

## Interfaces principais

### StudyCardItem

```ts
// src/app/shared/interfaces/study-template.interface.ts
interface StudyCardItem {
  id: string;
  title: string;
  description: string;
  bannerImage?: string;
  bannerColor: string;
  icon?: string;           // classe PrimeIcons
  detailRoute: string;
}
```

### BreadcrumbItem

```ts
interface BreadcrumbItem {
  label: string;
  href?: string; // ausente = item final, sem link
}
```

### StudyItem (entidade do banco)

```ts
interface StudyItem {
  _id: string;
  section: string;
  topic: string;
  title: string;
  status: 'studying' | 'completed' | 'planned';
  notes?: string;
  createdAt: string;
  updatedAt: string;
}
```

### NavItem (sidebar)

```ts
interface NavItem {
  key: string;
  label: string;
  href: string;
  icon: string;
}
```

## Constantes

`src/app/core/constants/study-topics.ts` — `STUDY_SECTIONS: StudyCardItem[]`

Fonte única de dados das 9 seções de estudo. Alterações aqui afetam `EstudosLabsComponent`, `DashboardComponent` e todas as seções de estudo.
