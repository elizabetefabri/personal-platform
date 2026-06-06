# ARQUIVO 4 — `manual/manual-deploy-subdominio-angular.md`

````md
# Manual Oficial de Deploy Angular v21 em Subdomínios Hostinger

## Objetivo

Este manual define o processo oficial para criar, configurar e publicar aplicações Angular v21 em subdomínios hospedados na Hostinger.

Este documento substitui qualquer documentação anterior baseada em Next.js.

---

# 1. Infraestrutura

## DNS

O domínio utiliza:

```txt
ns1.dns-parking.com
ns2.dns-parking.com
```
````

## Hospedagem

Fornecedor:

```txt
Hostinger
```

## Deploy

Método:

```txt
FTP/SFTP via extensão SFTP do VS Code
```

---

# 2. Pré-requisitos

Antes de iniciar, verificar:

```txt
[ ] Node.js instalado
[ ] npm instalado
[ ] Angular CLI instalado
[ ] VS Code instalado
[ ] Extensão SFTP instalada
[ ] Acesso ao painel Hostinger
[ ] Credenciais FTP disponíveis
[ ] Subdomínio criado ou a criar
```

---

# 3. Instalar Angular CLI

```bash
npm install -g @angular/cli
```

Validar instalação:

```bash
ng version
```

---

# 4. Criar Projeto Angular

Entrar no diretório do projeto:

```bash
cd personal-platform/studypanel
```

Criar frontend:

```bash
ng new frontend --routing --style=scss --standalone
```

Entrar no frontend:

```bash
cd frontend
```

---

# 5. Limpeza Inicial

Remover do projeto:

- conteúdo padrão do Angular;
- imagens de exemplo;
- ícones padrão;
- favicon padrão;
- estilos desnecessários.

Adicionar temporariamente no `app.component.html`:

```html
<h1>Hello World</h1>
```

---

# 6. Build de Produção

Executar:

```bash
ng build --configuration production
```

Pasta gerada:

```txt
dist/frontend/browser/
```

Esta é a pasta que deverá ser enviada para a Hostinger.

---

# 7. Criar Subdomínio na Hostinger

## Passo 1

Acessar o painel da Hostinger.

## Passo 2

Entrar em:

```txt
Hospedagem → Gerenciar
```

## Passo 3

Acessar:

```txt
Domínios → Subdomínios
```

## Passo 4

Criar subdomínio.

Exemplo:

```txt
studypanel.seudominio.com.br
```

## Passo 5

Definir diretório raiz.

Exemplo:

```txt
public_html/studypanel
```

## Passo 6

Salvar.

## Passo 7

Aguardar propagação.

---

# 8. Configuração Angular SPA

Aplicações Angular precisam de configuração para que as rotas funcionem corretamente ao atualizar a página.

Criar arquivo:

```txt
src/.htaccess
```

Conteúdo:

```apache
RewriteEngine On

RewriteCond %{REQUEST_FILENAME} !-f
RewriteCond %{REQUEST_FILENAME} !-d

RewriteRule ^ index.html [L]
```

Antes do deploy, garantir que esse arquivo esteja dentro de:

```txt
dist/frontend/browser/
```

---

# 9. Configuração SFTP no VS Code

Instalar extensão:

```txt
SFTP
ID: natizyskunk.sftp
```

Criar arquivo:

```txt
.vscode/sftp.json
```

Modelo:

```json
{
  "name": "StudyPanel",
  "host": "ftp.seudominio.com.br",
  "protocol": "ftp",
  "port": 21,
  "username": "usuario_ftp",
  "password": "senha_ftp",
  "remotePath": "/public_html/studypanel/",
  "uploadOnSave": false,
  "useTempFile": false,
  "openSsh": false,
  "watcher": {
    "files": "dist/frontend/browser/**/*",
    "autoUpload": true,
    "autoDelete": false
  },
  "ignore": [
    ".git",
    ".vscode",
    "node_modules",
    "src",
    "*.md",
    ".env",
    ".env.local",
    ".env.production"
  ]
}
```

---

# 10. Segurança

Nunca commitar:

```txt
.vscode/sftp.json
.env
.env.local
.env.production
```

Adicionar ao `.gitignore`:

```gitignore
.vscode/sftp.json
.env
.env.local
.env.production
node_modules/
dist/
```

---

# 11. Processo Oficial de Deploy

## Passo 1

Gerar build:

```bash
ng build --configuration production
```

## Passo 2

Validar pasta gerada:

```txt
dist/frontend/browser/
```

## Passo 3

Abrir Command Palette no VS Code:

```txt
Ctrl + Shift + P
```

## Passo 4

Executar:

```txt
SFTP: Upload Folder
```

## Passo 5

Selecionar:

```txt
dist/frontend/browser
```

## Passo 6

Enviar para:

```txt
/public_html/studypanel/
```

---

# 12. Validação Pós-Deploy

Acessar:

```txt
https://studypanel.seudominio.com.br
```

Validar:

```txt
[ ] Aplicação carrega
[ ] Estilos carregam
[ ] Imagens carregam
[ ] Rotas funcionam
[ ] Refresh da página não quebra
[ ] Console do navegador sem erros
[ ] Arquivo .htaccess está no servidor
```

---

# 13. Problemas Comuns

## Erro 404 ao atualizar página

Causa:

Angular é SPA e precisa redirecionar rotas para `index.html`.

Solução:

Verificar `.htaccess`.

---

## Estilos não carregam

Causas possíveis:

- build não foi gerado corretamente;
- upload incompleto;
- caminho errado no servidor.

Solução:

Rodar novamente:

```bash
ng build --configuration production
```

E reenviar a pasta:

```txt
dist/frontend/browser/
```

---

## SFTP não conecta

Verificar:

- host;
- usuário;
- senha;
- porta;
- remotePath;
- acesso FTP na Hostinger.

---

## Site não abriu após criar subdomínio

Possíveis causas:

- propagação ainda não concluída;
- diretório raiz incorreto;
- arquivos enviados para pasta errada.

---

# 14. Checklist Final

```txt
[ ] Angular CLI instalado
[ ] Projeto Angular criado
[ ] Projeto limpo
[ ] Build gerado
[ ] Subdomínio criado
[ ] Diretório raiz configurado
[ ] SFTP configurado
[ ] .htaccess criado
[ ] Arquivos enviados
[ ] Site validado
[ ] WIP.md atualizado
```

---

# 15. Regra Final

Sempre que um novo projeto Angular for criado e publicado em subdomínio, este manual deve ser seguido e o `WIP.md` deve ser atualizado com a data, o projeto, o subdomínio e as decisões tomadas.
