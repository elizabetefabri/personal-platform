# Solicitação de ajuste — Breadcrumbs, cadastros dinâmicos, backend e sidebar

## Contexto geral

O projeto **StudyPanel** precisa de ajustes na navegação, nos breadcrumbs, nos cadastros via modal, na integração com backend e na organização dos cards de **Estudos e Labs** e **Projetos**.

Execute **somente o que está descrito nesta solicitação**. Caso algum ponto esteja ambíguo ou dependa de uma decisão de produto, **pare e pergunte antes de implementar**. Não implemente comportamento novo por suposição.

---

## 1. Corrigir breadcrumbs em todas as páginas internas

Os **breadcrumbs** devem funcionar corretamente em todas as páginas internas do frontend, principalmente em:

- Estudos e Labs;
- páginas de categorias de estudo, como Backend, Banco de Dados, Cloud Computing etc.;
- páginas de itens/trilhas, como Node.js, Java, Python etc.;
- Projetos;
- Projetos Pessoais;
- Projetos Profissionais;
- detalhes de projetos;
- Vida Criativa;
- Painel Financeiro.

### Regra esperada

O breadcrumb deve representar o caminho real de navegação do usuário e cada etapa anterior deve ser clicável.

Exemplo esperado:

```txt
Dashboard > Estudos e Labs > Backend > Node.js
```

Regras:

1. `Dashboard` deve levar para a página inicial.
2. `Estudos e Labs` deve levar para a listagem principal de estudos.
3. `Backend` deve levar para a página da categoria Backend.
4. `Node.js` representa a página atual e pode ficar sem link ou com estado visual de item atual.
5. O mesmo padrão deve ser aplicado para Projetos e demais páginas internas.
6. Não remover breadcrumbs de páginas onde eles já existem; apenas corrigir o comportamento.
7. Não deixar breadcrumb visual sem link funcional.
8. Não criar breadcrumbs hardcoded quando for possível montar a estrutura a partir da rota e dos dados carregados.

### Exemplos adicionais

```txt
Dashboard > Projetos > Projetos Pessoais > Comanda Flow
Dashboard > Projetos > Projetos Profissionais > Rollout Service
Dashboard > Vida Criativa
Dashboard > Painel Financeiro
```

---

## 2. Página Estudos e Labs — cadastro dinâmico de tipos de curso

Ao entrar na página **Estudos e Labs**, deve existir uma área expansível ou seção de gerenciamento utilizando o componente de cadastro que já existe no projeto.

### Objetivo

Criar o fluxo de **Cadastrar Tipo de Curso**.

Esse cadastro deve controlar os cards principais da página **Estudos e Labs**, como:

- Backend;
- Banco de Dados;
- Cloud Computing;
- Containers e Kubernetes;
- DevOps;
- Frontend;
- Inteligência Artificial;
- Observability;
- Performance Engineering, caso esse item ainda exista no projeto.

### Comportamento esperado

1. A página **Estudos e Labs** deve exibir os cards dos tipos de curso.
2. Esses cards não devem mais depender de dados manuais fixos no frontend.
3. Os tipos de curso devem ser cadastrados no backend.
4. A listagem deve ser carregada do backend.
5. Ao cadastrar um novo tipo de curso, ele deve aparecer automaticamente:
   - na tabela de tipos de curso;
   - como novo card na página Estudos e Labs.
6. O cadastro deve acontecer sempre por **modal**.
7. O modal deve permitir informar, no mínimo:
   - nome do tipo de curso;
   - descrição curta;
   - slug/rota;
   - cor do card;
   - imagem da skill/card;
   - status ativo/inativo;
   - ordem de exibição.
8. A imagem deve substituir o ícone atual do card.
9. A cor escolhida pelo usuário deve ser aplicada no card.
10. Os itens adicionados manualmente no frontend devem ser migrados para o backend e removidos do hardcoded somente após a integração estar funcionando.

---

## 3. Página de detalhe do tipo de curso — cadastro de itens da trilha

Quando o usuário clicar em um tipo de curso, por exemplo **Backend**, deve abrir a página de detalhe dessa categoria.

Exemplo:

```txt
Dashboard > Estudos e Labs > Backend
```

Nessa página deve existir o botão para cadastrar um novo item da trilha.

### Exemplo de itens da trilha Backend

- Node.js;
- Java;
- Python;
- APIs REST;
- Microsserviços;
- Arquitetura Backend.

### Comportamento esperado

1. A página deve exibir os cards dos itens da trilha.
2. Deve existir uma tabela com todos os itens cadastrados para aquele tipo de curso.
3. O botão de cadastro deve abrir um modal.
4. O cadastro deve ser persistido no backend.
5. Ao cadastrar um novo item, ele deve aparecer automaticamente:
   - na tabela;
   - como novo card da categoria.
6. Cada card deve manter a mesma estrutura visual já existente.
7. O botão principal do card deve continuar sendo **Abrir Trilha**.
8. O card deve permitir acessar a página de detalhe do item.
9. A imagem da skill deve substituir o ícone atual do card.
10. A cor do card deve vir do cadastro.

### Campos mínimos para o item da trilha

- nome;
- descrição curta;
- slug/rota;
- tipo de curso relacionado;
- tag;
- cor do card;
- imagem da skill;
- status ativo/inativo;
- ordem de exibição.

---

## 4. Página de detalhe do item da trilha — cadastro de recursos

Quando o usuário entrar no item cadastrado, por exemplo **Node.js**, deve ser mantida a estrutura que já existe hoje no projeto.

Exemplo:

```txt
Dashboard > Estudos e Labs > Backend > Node.js
```

Essa página deve conter:

- botão **Cadastrar Recurso**;
- cadastro por modal;
- tabela com os recursos cadastrados;
- detalhe do recurso conforme estrutura já existente.

Não refaça a estrutura do zero se ela já existir. Apenas ajuste para funcionar com o fluxo dinâmico e com dados vindos do backend.

---

## 5. Página Projetos — cadastro dinâmico de projetos

Na página **Projetos**, deve existir o botão:

```txt
Cadastrar Projeto
```

Ao clicar, o usuário deve escolher o tipo de projeto que deseja cadastrar:

- Projeto Pessoal;
- Projeto Profissional.

O cadastro deve acontecer sempre por **modal** e os dados devem ser persistidos no backend.

---

## 6. Projetos Pessoais e Projetos Profissionais

Ao entrar em **Projetos Pessoais** ou **Projetos Profissionais**, deve existir um botão específico para o contexto da página:

```txt
Cadastrar Projeto Pessoal
Cadastrar Projeto Profissional
```

### Comportamento esperado

1. A página deve listar todos os projetos cadastrados em uma tabela.
2. A página também deve gerar automaticamente os cards desses projetos.
3. Os cards devem manter a mesma estrutura visual do projeto atual.
4. Cada projeto deve permitir:
   - abrir o repositório no GitHub;
   - abrir o deploy, quando existir;
   - abrir o detalhe do projeto dentro do próprio sistema.
5. Os dados não devem mais ficar fixos manualmente no frontend.
6. Os dados manuais existentes devem ser migrados para o backend antes de remover o hardcoded.

### Campos mínimos para projeto

- nome do projeto;
- tipo: pessoal ou profissional;
- descrição curta;
- tags;
- link do repositório GitHub;
- link do deploy;
- slug/rota;
- cor do card;
- imagem do projeto ou imagem da skill principal;
- status ativo/inativo;
- ordem de exibição.

### Definição de navegação do card

Para evitar confusão entre links internos e externos, usar esta regra:

1. O card deve ter um botão claro chamado **Ver Detalhes** para abrir o detalhe do projeto dentro do sistema.
2. Os links externos devem ficar separados em botões próprios:
   - **GitHub**;
   - **Deploy**.
3. O título do projeto pode continuar apenas como texto, sem link, para evitar conflito com os botões externos.
4. Adicionar tooltip nos botões externos informando para onde eles levam.

---

## 7. Modais de cadastro

Todos os novos cadastros solicitados devem ser feitos por modal.

Aplicar modal para:

- cadastrar tipo de curso;
- cadastrar item da trilha;
- cadastrar recurso;
- cadastrar projeto;
- cadastrar projeto pessoal;
- cadastrar projeto profissional.

### Regras dos modais

1. O modal deve reaproveitar componentes existentes sempre que possível.
2. Não duplicar lógica desnecessária.
3. Validar campos obrigatórios.
4. Exibir mensagem de erro quando a chamada ao backend falhar.
5. Atualizar tabela e cards após salvar com sucesso.
6. Permitir cancelar sem salvar.
7. Manter padrão visual do projeto.
8. Garantir acessibilidade básica, como foco no modal e botão de fechar com nome acessível.

---

## 8. Backend

Adicionar no backend todas as alterações necessárias para que as chamadas do frontend funcionem corretamente.

### Regras

1. Criar ou ajustar as entidades/modelos necessários.
2. Criar ou ajustar endpoints para:
   - tipos de curso;
   - itens da trilha;
   - recursos;
   - projetos;
   - tipos de projeto.
3. Garantir operações de:
   - criar;
   - listar;
   - buscar por ID ou slug;
   - atualizar;
   - excluir ou inativar.
4. Respeitar a stack já existente no backend.
5. Não trocar tecnologia do backend.
6. Não inventar arquitetura nova sem necessidade.
7. Validar se os dados manuais atuais precisam virar carga inicial/seed no banco.
8. Documentar endpoints criados ou alterados.
9. Garantir que o frontend consuma os endpoints corretos.
10. Validar frontend e backend rodando juntos.

---

## 9. Tooltips

Adicionar tooltips onde fizer sentido para melhorar a compreensão do usuário.

### Locais sugeridos

- botão de abrir trilha;
- botão de ver detalhes;
- botão GitHub;
- botão Deploy;
- botão editar;
- botão excluir;
- botão cadastrar;
- botão voltar;
- breadcrumbs longos;
- imagens ou ícones com significado visual.

Os tooltips devem ser curtos e objetivos.

Exemplos:

```txt
Abrir trilha de estudos
Ver detalhes do projeto
Abrir repositório no GitHub
Abrir projeto publicado
Voltar para a página anterior
Editar item
Excluir item
```

---

## 10. Sidebar — adicionar Vida Criativa e Painel Financeiro

Adicionar os componentes/páginas abaixo no sidebar:

- Vida Criativa;
- Painel Financeiro.

### Estrutura esperada do sidebar

```txt
Dashboard
Estudos e Labs
Projetos
Vida Criativa
Painel Financeiro
Sair
```

Regras:

1. Não alterar o comportamento do Dashboard.
2. Não alterar os ícones do sidebar sem solicitação específica.
3. Garantir que os links estejam funcionando.
4. Garantir que as rotas existam.
5. Garantir que os breadcrumbs dessas páginas funcionem.
6. Manter o padrão visual atual do sidebar.

---

## 11. Remover dados hardcoded com segurança

Os itens que hoje foram adicionados manualmente no frontend não devem ser simplesmente apagados.

### Processo correto

1. Identificar todos os dados hardcoded usados nos cards e tabelas.
2. Criar as estruturas correspondentes no backend.
3. Migrar esses dados para o banco ou seed.
4. Conectar o frontend aos endpoints.
5. Validar se os cards e tabelas continuam iguais visualmente.
6. Só depois remover os dados hardcoded do frontend.

---

## 12. Critérios de aceite

A implementação só estará concluída quando:

- os breadcrumbs funcionarem em todas as páginas internas;
- os breadcrumbs anteriores forem clicáveis;
- a página Estudos e Labs listar tipos de curso vindos do backend;
- o cadastro de tipo de curso funcionar via modal;
- o cadastro criar automaticamente tabela e card;
- a página de detalhe do tipo de curso listar os itens da trilha vindos do backend;
- o cadastro de item da trilha funcionar via modal;
- o detalhe do item manter a estrutura de recursos já existente;
- a página Projetos permitir cadastro por tipo;
- Projetos Pessoais e Projetos Profissionais listarem dados vindos do backend;
- os cards de projeto tiverem botões separados para Detalhes, GitHub e Deploy;
- os ícones dos cards forem substituídos por imagens;
- a cor do card vier do cadastro;
- tooltips forem aplicados nos pontos necessários;
- Vida Criativa e Painel Financeiro aparecerem no sidebar;
- o backend tiver endpoints funcionando para todos os cadastros;
- o frontend consumir corretamente o backend;
- os dados hardcoded forem removidos somente após migração segura;
- não houver erro no console do navegador;
- frontend e backend iniciarem corretamente;
- a documentação for atualizada com as alterações realizadas.

---

## 13. Commits sugeridos

Separar os commits por contexto, com mensagens em português e claras.

```bash
git add .
git commit -m "feat: corrigir breadcrumbs das páginas internas"
```

```bash
git add .
git commit -m "feat: adicionar cadastros dinâmicos em estudos e labs"
```

```bash
git add .
git commit -m "feat: adicionar cadastros dinâmicos de projetos"
```

```bash
git add .
git commit -m "feat: integrar frontend com endpoints de estudos e projetos"
```

```bash
git add .
git commit -m "feat: adicionar páginas vida criativa e painel financeiro no sidebar"
```

```bash
git add .
git commit -m "docs: atualizar documentação dos novos fluxos"
```

Ao final, fazer push para o repositório remoto.
