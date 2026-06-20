# SEO

## Status

Existente (meta tags estáticas em index.html)

## Meta tags configuradas

| Tag                        | Valor                                    |
| -------------------------- | ---------------------------------------- |
| `<html lang>`              | `pt-BR`                                  |
| `<title>`                  | StudyPanel \| Elizabete Fabri            |
| `meta[name=description]`   | Descrição do painel de estudos           |
| `meta[name=keywords]`      | StudyPanel, Angular, TypeScript, ...     |
| `meta[name=author]`        | Elizabete Fabri                          |
| `meta[name=robots]`        | index, follow, max-image-preview:large   |
| `meta[name=googlebot]`     | index, follow, max-image-preview:large   |
| `meta[property=og:title]`  | StudyPanel \| Elizabete Fabri            |
| `meta[property=og:type]`   | website                                  |
| `meta[property=og:url]`    | URL de produção                          |

## Arquivos estáticos

| Arquivo              | Localização             |
| -------------------- | ----------------------- |
| `robots.txt`         | `public/robots.txt`     |
| `sitemap.xml`        | `public/sitemap.xml`    |

## Observações

- O StudyPanel é uma SPA — não há SSR, então meta tags são estáticas
- Para conteúdo dinâmico por rota, usar `@angular/platform-browser` `Title` e `Meta` services
- `sitemap.xml` lista todas as rotas públicas para indexação
