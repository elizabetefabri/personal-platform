# Manual de Git Workflow

## Objetivo

Definir o padrão de branches, commits e fluxo de versionamento da `personal-platform`.

## Branches

### Principal

```txt
main

Desenvolvimento
develop
Features
feature/nome-da-feature
Correções
hotfix/nome-da-correcao
Commits

Usar Conventional Commits.

Exemplos:

feat: adiciona tela inicial
fix: corrige erro no deploy
docs: atualiza manual de criação Angular
style: ajusta responsividade
refactor: reorganiza estrutura de pastas
test: adiciona testes unitários
chore: atualiza dependências
Regras
Não commitar arquivos sensíveis.
Não commitar .env.
Não commitar .vscode/sftp.json.
Não commitar arquivos privados de governança.
Não fazer push sem revisar alterações.
Não fazer merge direto na main sem validação.
Antes do commit

Verificar:

[ ] Código funcionando
[ ] Sem credenciais
[ ] Sem arquivos privados
[ ] Build validado
[ ] Documentação atualizada
```
