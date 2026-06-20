# Serviços

## Serviços existentes

| Serviço              | Caminho                                          | Responsabilidade                        |
| -------------------- | ------------------------------------------------ | --------------------------------------- |
| BreadcrumbService    | `src/app/core/services/breadcrumb.service.ts`    | Estado global de breadcrumbs e título   |
| StudyItemService     | `src/app/core/services/study-item.service.ts`    | CRUD de itens de estudo (MongoDB/API)   |

## BreadcrumbService

```ts
// API pública
service.set(extra: BreadcrumbItem | null)
service.setHidden(hidden: boolean)
service.setPageTitle(title: string)
service.clear()

// Signals de leitura (readonly)
service.extra()      // BreadcrumbItem | null
service.hidden()     // boolean
service.pageTitle()  // string
```

## StudyItemService

```ts
// API pública (retornam Observable)
service.list(section: string, topic: string): Observable<StudyItem[]>
service.listAll(): Observable<StudyItem[]>
service.create(item: Partial<StudyItem>): Observable<StudyItem>
service.update(id: string, item: Partial<StudyItem>): Observable<StudyItem>
service.delete(id: string): Observable<void>
```

## Pattern de mock em testes

```ts
const studyItemServiceMock = {
  list: jest.fn(() => of([])),
  listAll: jest.fn(() => of([])),
  create: jest.fn(() => of({})),
  update: jest.fn(() => of({})),
  delete: jest.fn(() => of(undefined)),
};
```
