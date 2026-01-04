# 📘 Manual de Deploy - Next.js para FTP (Hostinger)

## 🎯 Objetivo

Este manual documenta como configurar e fazer deploy de projetos Next.js via FTP para subdomínios na Hostinger.

## 📋 Pré-requisitos

1. **Extensão VS Code SFTP instalada:**

   - Nome: `SFTP`
   - ID: `natizyskunk.sftp`
   - Instalar via: `Ctrl+Shift+X` → buscar "SFTP"

2. **Projeto Next.js criado**
3. **Acesso FTP do servidor**

## 🗂️ Arquivos Necessários

### 1. `next.config.ts` (raiz do projeto)

**Configuração para exportação estática:**

```typescript
/** @type {import('next').NextConfig} */

const nextConfig = {
  output: "export", // Gera site estático na pasta "out/"
  images: {
    unoptimized: true, // Necessário para export estático
  },
  trailingSlash: true, // Adiciona "/" no final das URLs
};

module.exports = nextConfig;
```

**Observações:**

- `output: "export"` → Gera arquivos estáticos (HTML, CSS, JS)
- `images.unoptimized: true` → Desabilita otimização de imagens (necessário para FTP)
- `trailingSlash: true` → Melhora compatibilidade com servidores estáticos
- ⚠️ `redirects()` e `headers()` NÃO funcionam com `output: "export"`

### 2. `.vscode/sftp.json` (configuração FTP)

**Crie a pasta `.vscode/` na raiz do projeto e adicione:**

```json
{
  "name": "Nome do Projeto",
  "host": "ftp.seudominio.com.br",
  "protocol": "ftp",
  "port": 21,
  "username": "seu_usuario_ftp",
  "password": "sua_senha_ftp",
  "remotePath": "/subdominio/",
  "uploadOnSave": true,
  "useTempFile": false,
  "openSsh": false,
  "watcher": {
    "files": "out/**/*",
    "autoUpload": true,
    "autoDelete": false
  },
  "ignore": [
    ".vscode",
    ".git",
    ".DS_Store",
    "node_modules",
    ".next",
    "src",
    "public",
    "*.md",
    "*.json",
    "*.ts",
    "*.js",
    "*.mjs",
    ".env*"
  ]
}
```

**Parâmetros importantes:**

- `name`: Nome identificador do projeto
- `host`: Endereço FTP (geralmente `ftp.seudominio.com.br`)
- `username`: Usuário FTP fornecido pela Hostinger
- `password`: Senha FTP
- `remotePath`: Caminho no servidor (ex: `/subdominio/`)
- `uploadOnSave`: `true` para upload automático ao salvar
- `watcher.files`: `"out/**/*"` para monitorar apenas a pasta de build
- `ignore`: Arquivos que não devem ser enviados

### 3. `deploy.js` (script de deploy automático)

**Crie na raiz do projeto:**

```javascript
const ftp = require("basic-ftp");
const path = require("path");

async function deploy() {
  const client = new ftp.Client();
  client.ftp.verbose = true;

  try {
    console.log("🔌 Conectando ao servidor FTP...");
    await client.access({
      host: "ftp.seudominio.com.br",
      user: "seu_usuario_ftp",
      password: "sua_senha_ftp",
      port: 21,
      secure: false,
    });

    console.log("✅ Conectado!");
    console.log("📤 Enviando arquivos da pasta out/ para /subdominio/...");

    await client.ensureDir("/subdominio");
    await client.clearWorkingDir();
    await client.uploadFromDir(path.join(__dirname, "out"));

    console.log("✅ Deploy concluído com sucesso!");
    console.log("🌐 Site disponível em: https://subdominio.seudominio.com.br/");
  } catch (err) {
    console.error("❌ Erro no deploy:", err);
  }

  client.close();
}

deploy();
```

**Configurar:**

- Altere `host`, `user`, `password`
- Altere `remotePath` (`/subdominio/`)
- Altere a URL final no console.log

### 4. `package.json` (adicionar script de deploy)

**Adicione no objeto `scripts`:**

```json
{
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "eslint",
    "deploy": "npm run build && node deploy.js"
  }
}
```

## 🚀 Instalação e Configuração

### Passo 1: Instalar dependência FTP

```bash
npm install --save-dev basic-ftp
```

### Passo 2: Configurar informações FTP

1. **Obter credenciais FTP na Hostinger:**

   - Painel Hostinger → Hospedagem → Gerenciar
   - Seção "FTP" → Criar/Ver credenciais
   - Anotar: Host, Usuário, Senha, Porta (geralmente 21)

2. **Criar subdomínio:**

   - Painel Hostinger → Domínios → Gerenciar
   - Criar subdomínio (ex: `comandaflow.seudominio.com.br`)
   - Verificar o caminho no servidor (ex: `/comandaflow/`)

3. **Atualizar arquivos com suas credenciais:**
   - `.vscode/sftp.json`
   - `deploy.js`

## 📦 Processo de Deploy

### Método 1: Deploy Automático (Recomendado)

**Comando único que faz build + upload:**

```bash
npm run deploy
```

Isso irá:

1. Executar `npm run build` (gera pasta `out/`)
2. Executar `node deploy.js` (envia via FTP)

### Método 2: Deploy Manual via VS Code

**Passo a passo:**

1. **Gerar build:**

   ```bash
   npm run build
   ```

2. **Upload via extensão SFTP:**
   - Abrir Command Palette: `Ctrl+Shift+P`
   - Digitar: `SFTP: Upload Folder`
   - Selecionar pasta `out/`
   - Aguardar conclusão

Ou:

- Clicar com botão direito na pasta `out/`
- Selecionar `SFTP: Upload Folder`

### Método 3: Upload Automático ao Salvar

Se `"uploadOnSave": true` estiver configurado:

- Qualquer arquivo salvo na pasta `out/` será automaticamente enviado
- Útil para pequenas correções

## 🔧 Comandos Úteis

```bash
# Desenvolvimento local
npm run dev

# Gerar build de produção
npm run build

# Deploy completo (build + upload)
npm run deploy

# Apenas lint
npm run lint
```

## 📁 Estrutura de Arquivos

```
seu-projeto/
├── .vscode/
│   └── sftp.json              ← Configuração FTP
├── out/                       ← Build gerado (não commitar)
├── src/                       ← Código fonte
├── public/                    ← Arquivos estáticos
├── deploy.js                  ← Script de deploy
├── next.config.ts             ← Configuração Next.js
├── package.json               ← Dependências e scripts
└── README.md
```

## ⚠️ Problemas Comuns

### 1. Erro "Config Not Found"

**Causa:** Arquivo `.vscode/sftp.json` não encontrado ou mal formatado

**Solução:**

- Verificar se o arquivo existe em `.vscode/sftp.json`
- Validar JSON (sem vírgulas extras, aspas corretas)
- Reabrir VS Code

### 2. Erro de conexão FTP

**Causa:** Credenciais incorretas ou firewall

**Solução:**

- Verificar host, usuário, senha no painel Hostinger
- Testar conexão com cliente FTP externo (FileZilla)
- Verificar se porta 21 está aberta

### 3. Build falha

**Causa:** Erros no código ou configuração incorreta

**Solução:**

```bash
# Ver erros detalhados
npm run build

# Verificar sintaxe
npm run lint
```

### 4. Imagens não carregam

**Causa:** Otimização de imagens não compatível com export estático

**Solução:** Já está configurado com `unoptimized: true` no `next.config.ts`

### 5. Rotas não funcionam (404)

**Causa:** Servidor precisa de configuração para SPA

**Solução:** Adicionar arquivo `.htaccess` na raiz do servidor:

```apache
# .htaccess para Next.js Export
RewriteEngine On
RewriteBase /

# Se for um arquivo ou diretório existente, servir diretamente
RewriteCond %{REQUEST_FILENAME} !-f
RewriteCond %{REQUEST_FILENAME} !-d

# Caso contrário, redirecionar para index.html
RewriteRule ^(.*)$ index.html [L]
```

## 🔒 Segurança

### ⚠️ IMPORTANTE: Proteger credenciais

**Nunca commitar senhas para Git!**

1. **Adicionar ao `.gitignore`:**

```gitignore
.vscode/sftp.json
deploy.js
.env
.env.local
```

2. **Alternativa: Usar variáveis de ambiente**

Criar `.env.local`:

```env
FTP_HOST=ftp.seudominio.com.br
FTP_USER=seu_usuario
FTP_PASS=sua_senha
FTP_PATH=/subdominio/
```

Atualizar `deploy.js`:

```javascript
require("dotenv").config();

await client.access({
  host: process.env.FTP_HOST,
  user: process.env.FTP_USER,
  password: process.env.FTP_PASS,
  port: 21,
  secure: false,
});
```

Instalar:

```bash
npm install --save-dev dotenv
```

## 📊 Checklist de Deploy

- [ ] Criar projeto Next.js
- [ ] Configurar `next.config.ts` com `output: "export"`
- [ ] Criar `.vscode/sftp.json` com credenciais FTP
- [ ] Criar `deploy.js` com script de upload
- [ ] Adicionar script `"deploy"` no `package.json`
- [ ] Instalar `basic-ftp`: `npm install --save-dev basic-ftp`
- [ ] Testar build local: `npm run build`
- [ ] Fazer primeiro deploy: `npm run deploy`
- [ ] Verificar site no navegador
- [ ] Adicionar `.gitignore` para proteger credenciais

## 🎓 Exemplo Completo

### Criar novo projeto do zero:

```bash
# 1. Criar projeto Next.js
npx create-next-app@latest meu-projeto
cd meu-projeto

# 2. Instalar dependência FTP
npm install --save-dev basic-ftp

# 3. Criar arquivos de configuração
# - Criar .vscode/sftp.json (copiar template acima)
# - Criar deploy.js (copiar template acima)
# - Editar next.config.ts (adicionar output: "export")
# - Editar package.json (adicionar script deploy)

# 4. Configurar credenciais FTP nos arquivos

# 5. Fazer deploy
npm run deploy

# 6. Acessar site
# https://subdominio.seudominio.com.br/
```

**Documentação:**

- Next.js: https://nextjs.org/docs
- Static Export: https://nextjs.org/docs/app/building-your-application/deploying/static-exports

## 📝 Notas Finais

- ✅ Este setup é ideal para sites estáticos (blogs, portfolios, landing pages)
- ❌ Não suporta Server-Side Rendering (SSR) ou API Routes
- ✅ Hospedagem FTP é mais barata que servidores Node.js
- ✅ Performance excelente (arquivos estáticos são muito rápidos)
- ✅ Compatível com qualquer hospedagem que suporte arquivos HTML

**Criado em:** Janeiro 2026  
**Última atualização:** Janeiro 2026
