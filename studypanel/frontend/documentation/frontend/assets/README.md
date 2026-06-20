# Assets

## Localização

`public/assets/images/` — imagens dos cards de estudo e projetos

## Convenção de nomes

```
public/
└── assets/
    └── images/
        ├── backend.png
        ├── cloud.png
        ├── containers-kubernetes.png
        ├── devops.png
        ├── frontend.png
        ├── inteligencia-artificial.png
        ├── banco-de-dados.png
        ├── observability.png
        └── performance-engineering.png
```

## Como usar nos cards

```ts
// Em StudyCardItem:
{ bannerImage: 'assets/images/backend.png', bannerColor: '#1a365d' }
```

Se `bannerImage` não carregar, o `bannerColor` é usado como fallback via `onerror` no template.

## Arquivos estáticos

| Arquivo         | Localização       | Finalidade                       |
| --------------- | ----------------- | -------------------------------- |
| `robots.txt`    | `public/`         | Instruções para crawlers         |
| `sitemap.xml`   | `public/`         | URLs indexáveis                  |
| `favicon.ico`   | `public/`         | Ícone do browser                 |
