## `CHANGELOG.md`

# Changelog

Todas as alterações relevantes deste repositório serão documentadas neste arquivo.

O formato segue uma organização simples por versão e data.

---

## [0.2.0] - 2026-06-06

### Adicionado

- **StudyPanel Frontend - Angular v21**
  - Conversão de componentes para formato .tsx com templates inline
  - 7 componentes de página (Dashboard, Login, Perfil, Cadastro, Cloud, Rollout, Settings)
  - 4 componentes de layout (Header, Sidebar, Footer, Breadcrumbs)
  - Sistema de rotas completo (/dashboard, /login, /perfil, /cadastro, /cloud, /rollout, /settings)
  - CSS Design System com 20+ variáveis de cor
  - Deploy via FTP com script deploy.ts
  - Suporte JSX em arquivos TypeScript
  - Responsive design com media queries (320px - 1920px)

### Alterado

- Atualização do SCSS para usar módulos @use em vez de @import (eliminação de deprecations)
- Atualização de tsconfig.json para suportar JSX
- Adição de basic-ftp como dependência dev
- Package.json com scripts de deploy (npm run deploy)

### Segurança

- .env.deployment adicionado ao .gitignore
- Arquivo .env.deployment.example como template

---

## [0.1.0] - AAAA-MM-DD

### Adicionado

- Criação da estrutura inicial da `personal-platform`.
- Criação do diretório `docs/`.
- Criação do manual de criação de projetos Angular.
- Criação do manual de deploy Angular em subdomínios Hostinger.
- Criação do manual de workflow Git.
- Definição inicial da estrutura `frontend/` e `backend/`.

### Alterado

- Padronização dos novos projetos frontend com Angular v21, TypeScript, SCSS e Standalone Components.

### Segurança

- Definição de arquivos privados no `.gitignore`.
- Bloqueio de arquivos internos de governança e contexto.
