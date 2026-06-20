# Dados

## Fonte de dados das seções

`src/app/core/constants/study-topics.ts`

### `STUDY_SECTIONS`

Array com as 9 seções de estudo técnico. Estrutura:

```ts
export const STUDY_SECTIONS: StudyCardItem[] = [
  {
    id: 'backend',
    title: 'Backend',
    description: '...',
    bannerImage: 'assets/images/backend.png',
    bannerColor: '#1a365d',
    icon: 'pi-server',
    detailRoute: '/backend',
  },
  // ... mais 8 seções
];
```

### Importação

```ts
import { STUDY_SECTIONS } from '../../core/constants/study-topics';
```

### Regra

Qualquer alteração neste arquivo afeta:
- `EstudosLabsComponent` (lista as 9 seções)
- `DashboardComponent` (overview e gráficos)
- Todos os links de navegação das seções

Não remover campos `id` ou `detailRoute` sem atualizar as rotas correspondentes.
