# Solicitação completa — Implementação do módulo Culinária

## Contexto do projeto

Estamos evoluindo o projeto **Caderno de Anotações / Tech Note**, que hoje possui um fluxo funcional de navegação, cadastro e detalhamento na área **Estudos e Labs**.

O módulo **Estudos e Labs** já segue esta estrutura:

```txt
Estudos e Labs
  └── Área de estudo
        └── Item da trilha
              └── Recurso / Conteúdo
                    └── Detalhe do recurso
```

Exemplo atual:

```txt
Estudos e Labs
  └── Backend
        └── Node.js
              └── xxx
                    └── Anotações e Resumos
                    └── Progresso e Metas
                    └── Professor IA
                    └── Recursos e Referências
                    └── Laboratórios Práticos
                    └── Quiz de Revisão
                    └── Histórico de Sessões
```

Agora precisamos iniciar a criação completa do item **Culinária**, que atualmente aparece no menu lateral, mas ainda não possui funcionalidade implementada.

A implementação deve seguir o mesmo padrão visual, técnico e arquitetural já usado no projeto:

- Angular 21
- Standalone Components
- TypeScript
- SCSS
- PrimeNG + PrimeIcons
- Backend em Go
- MongoDB
- Clean Architecture
- Integração real com backend
- CRUD completo
- Validação no frontend e no backend
- Modais para cadastro e edição
- Confirmação antes de exclusão
- Breadcrumb funcional
- Cards padronizados
- Tabelas quando necessário
- Botões com o mesmo padrão visual do projeto

---

## Objetivo principal

Criar o módulo **Culinária** como uma área pessoal para organização de receitas, categorias culinárias, links úteis, vídeos do YouTube, ingredientes, modo de preparo, dicas, substituições, planejamento e histórico de receitas testadas.

O módulo deve funcionar como um **caderno culinário digital**, mantendo a mesma lógica de organização do restante do projeto, porém adaptado para receitas.

---

## Estrutura esperada do módulo Culinária

O fluxo deve ser organizado em três níveis principais:

```txt
Culinária
  └── Categoria culinária
        └── Receita
              └── Detalhe da receita
```

Exemplo:

```txt
Culinária
  └── Massas
        └── Lasanha de frango
              └── Ingredientes
              └── Modo de preparo
              └── Vídeo do YouTube
              └── Dicas e substituições
              └── Lista de compras
              └── Histórico de preparo
```

---

## 1. Página principal: Culinária

### Rota

```txt
/culinaria
```

### Objetivo

Exibir todas as categorias culinárias cadastradas.

Essa página deve funcionar de forma parecida com a página **Estudos e Labs**, porém com o contexto de receitas.

### Elementos obrigatórios

A página deve conter:

- Título da página: **Culinária**
- Descrição curta da página
- Breadcrumb: `Dashboard > Culinária`
- Botão **Voltar**
- Botão **Nova categoria culinária**
- Grid de cards com as categorias cadastradas
- Estado vazio quando não houver categorias cadastradas
- Integração com backend para listar categorias
- Loading enquanto carrega os dados
- Mensagem de erro caso a API falhe

### Texto sugerido para a página

Título:

```txt
Culinária
```

Descrição:

```txt
Organize suas receitas, categorias, vídeos, ingredientes e ideias culinárias em um só lugar.
```

### Estado vazio

Quando não houver nenhuma categoria cadastrada, exibir uma mensagem amigável:

```txt
Nenhuma categoria culinária cadastrada ainda.
Crie sua primeira categoria para começar a organizar suas receitas.
```

Botão no estado vazio:

```txt
Criar primeira categoria
```

---

## 2. Card de categoria culinária

Cada categoria culinária deve ser exibida em um card.

### Exemplos de categorias

- Café da manhã
- Almoço
- Jantar
- Sobremesas
- Massas
- Saladas
- Marmitas
- Receitas rápidas
- Receitas sem glúten
- Receitas econômicas
- Receitas para congelar
- Receitas favoritas
- Receitas testadas
- Receitas pendentes

### Informações do card

| Campo | Descrição |
|---|---|
| Nome | Nome da categoria |
| Descrição | Resumo curto da categoria |
| Ícone | Ícone PrimeIcons |
| Cor do card | Cor ou gradiente configurável |
| Tag | Texto curto de identificação |
| Total de receitas | Quantidade de receitas cadastradas |
| Botão principal | Abrir categoria |
| Ação editar | Editar categoria |
| Ação excluir | Excluir categoria |

### Exemplo de card

```txt
Massas
Receitas de massas, molhos, lasanhas e pratos italianos.

Tag: MASSAS
Total: 8 receitas

[ Abrir categoria ]
```

---

## 3. Modal: Nova categoria culinária

### Quando abrir

Abrir ao clicar no botão:

```txt
Nova categoria culinária
```

### Título do modal

Para cadastro:

```txt
Cadastrar categoria culinária
```

Para edição:

```txt
Editar categoria culinária
```

### Campos do formulário

| Campo | Tipo | Obrigatório | Validação |
|---|---|---:|---|
| Nome | input text | Sim | mínimo 2 caracteres |
| Slug | input text | Sim | kebab-case, único |
| Descrição | input text ou textarea | Não | máximo 180 caracteres |
| Tag | input text | Não | máximo 20 caracteres |
| Cor do card | input text | Não | aceitar hex ou linear-gradient |
| Ícone | input text | Não | PrimeIcons, exemplo: `pi pi-book` |
| URL da imagem | input text | Não | URL válida ou caminho local |
| Ordem | number | Não | número maior ou igual a 0 |
| Ativo | checkbox/toggle | Não | boolean |

### Botões do modal

Os botões devem seguir o padrão visual do projeto:

```txt
[ Cancelar ] [ Salvar categoria ]
```

Regras:

- Cada botão deve ocupar 50% da largura.
- Total dos dois botões deve ocupar 100%.
- Botão cancelar: transparente com borda `#201F25`.
- Botão salvar: fundo `#201F25`.
- Remover borda/outline verde dos inputs.
- Focus dos inputs deve usar `#201F25`.

---

## 4. Página da categoria culinária

### Rota

```txt
/culinaria/:categorySlug
```

Exemplo:

```txt
/culinaria/massas
```

### Objetivo

Listar todas as receitas cadastradas dentro de uma categoria.

### Elementos obrigatórios

A página deve conter:

- Breadcrumb: `Dashboard > Culinária > Massas`
- Botão **Voltar**
- Título da categoria
- Descrição da categoria
- Botão **Nova receita**
- Grid ou tabela de receitas cadastradas
- Estado vazio quando não houver receitas
- Loading
- Mensagem de erro em falha de API

### Nome do botão

```txt
Nova receita
```

---

## 5. Card de receita

Cada receita deve ser exibida em um card.

### Informações do card

| Campo | Descrição |
|---|---|
| Nome da receita | Exemplo: `Lasanha de frango` |
| Descrição curta | Breve resumo da receita |
| Tempo de preparo | Exemplo: `45 min` |
| Rendimento | Exemplo: `6 porções` |
| Dificuldade | Fácil, Médio ou Difícil |
| Status | Pendente, Testada, Favorita, Ajustar receita |
| Tags | Exemplo: `rápida`, `sem glúten`, `econômica` |
| Imagem | Opcional |
| Botão principal | Ver receita |
| Ação editar | Editar receita |
| Ação excluir | Excluir receita |

### Regra visual importante

Os botões do card devem ficar sempre no footer do card.

Se a descrição for muito longa, deve exibir `...` com reticências.

Usar estrutura com `display: flex`, `flex-direction: column` e `margin-top: auto` na área dos botões.

---

## 6. Modal: Nova receita

### Quando abrir

Abrir ao clicar no botão:

```txt
Nova receita
```

### Título do modal

Para cadastro:

```txt
Cadastrar receita
```

Para edição:

```txt
Editar receita
```

### Campos do formulário

| Campo | Tipo | Obrigatório | Validação |
|---|---|---:|---|
| Nome da receita | input text | Sim | mínimo 3 caracteres |
| Slug | input text | Sim | kebab-case, único dentro da categoria |
| Descrição curta | textarea | Não | máximo 240 caracteres |
| Categoria | select | Sim | categoria atual já selecionada |
| Tempo de preparo | number | Não | maior ou igual a 0 |
| Tempo de cozimento | number | Não | maior ou igual a 0 |
| Rendimento | input text | Não | exemplo: `6 porções` |
| Dificuldade | select | Não | Fácil, Médio, Difícil |
| Status | select | Sim | Pendente, Testada, Favorita, Ajustar receita |
| Tags | chips ou input | Não | separadas por vírgula |
| URL da imagem | input text | Não | URL válida ou `/assets/images/...` |
| URL do YouTube | input text | Não | aceitar links do YouTube |
| Link da receita original | input text | Não | URL válida |
| Observações | textarea | Não | anotações livres |

### Campos culinários adicionais

Adicionar campos específicos para receita:

| Campo | Tipo | Descrição |
|---|---|---|
| Ingredientes | textarea ou lista dinâmica | Lista de ingredientes |
| Modo de preparo | textarea ou lista dinâmica | Passo a passo da receita |
| Utensílios necessários | textarea | Exemplo: panela, forma, liquidificador |
| Dicas | textarea | Dicas pessoais para melhorar a receita |
| Substituições | textarea | Trocas possíveis de ingredientes |
| Como armazenar | textarea | Geladeira, freezer, validade aproximada |
| Custo estimado | input number | Valor aproximado da receita |
| Nota pessoal | rating ou number | Avaliação de 1 a 5 |
| Já testei? | checkbox/toggle | Marca se a receita já foi preparada |
| Data em que testei | date | Data opcional |
| Link de vídeo alternativo | input text | Outro vídeo ou referência |

### Botões do modal

```txt
[ Cancelar ] [ Salvar receita ]
```

Regras:

- Os dois botões devem ocupar 100% da largura.
- Cada botão deve ocupar 50%.
- Botão cancelar transparente com borda `#201F25`.
- Botão salvar com fundo `#201F25`.
- Validação visual clara nos campos inválidos.
- Não usar outline verde nos inputs.

---

## 7. Página de detalhe da receita

### Rota

```txt
/culinaria/:categorySlug/:recipeSlug
```

Exemplo:

```txt
/culinaria/massas/lasanha-de-frango
```

### Objetivo

Exibir o detalhe completo da receita cadastrada.

### Elementos obrigatórios

A página deve conter:

- Breadcrumb: `Dashboard > Culinária > Massas > Lasanha de frango`
- Botão **Voltar para Massas**
- Título da receita
- Categoria
- Status
- Tempo de preparo
- Rendimento
- Dificuldade
- Tags
- Imagem, se existir
- Blocos expansíveis com os detalhes

---

## 8. Blocos da página de detalhe da receita

A página de detalhe deve ter blocos parecidos com os blocos de estudos, mas adaptados para culinária.

### 8.1 Ingredientes

- Exibir todos os ingredientes da receita.
- Permitir marcar ingredientes como conferidos.
- Futuramente, permitir gerar lista de compras.

### 8.2 Modo de preparo

- Exibir o passo a passo da receita.
- Cada etapa deve ser clara e numerada.

### 8.3 Vídeos e links úteis

- Exibir link do YouTube.
- Exibir link da receita original.
- Permitir cadastrar links extras.
- Validar se o link é válido.

Sugestões de links:

- YouTube
- Blog de receita
- Instagram
- TikTok
- Site oficial
- Livro ou referência manual

### 8.4 Dicas, substituições e observações

- Registrar dicas pessoais.
- Guardar ajustes realizados.
- Registrar substituições de ingredientes.
- Escrever observações do que funcionou ou não.

### 8.5 Planejamento e lista de compras

- Gerar ou organizar uma lista de compras com base nos ingredientes.
- Marcar itens comprados.
- Registrar quantidade necessária.
- Futuramente, permitir exportar lista.

### 8.6 Histórico de preparo

- Registrar quando a receita foi feita.
- Registrar nota pessoal.
- Registrar comentários.
- Marcar se a receita deu certo.
- Registrar ajustes para a próxima tentativa.

### 8.7 Variações da receita

Permitir guardar versões alternativas, como:

- versão sem glúten;
- versão mais econômica;
- versão para air fryer;
- versão para congelar;
- versão vegetariana.

### 8.8 Professor IA / Assistente culinário

Criar um bloco futuro para IA sugerir:

- substituições de ingredientes;
- organização da receita;
- melhoria do modo de preparo;
- ideias de acompanhamento;
- lista de compras;
- reaproveitamento de ingredientes.

Esse bloco pode iniciar como **BETA**, igual ao padrão usado em Estudos e Labs.

---

## 9. Backend — implementação obrigatória

Criar integração completa com backend em Go e MongoDB.

Seguir Clean Architecture:

```txt
backend/
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   └── repository/
│   ├── usecase/
│   ├── handler/
│   ├── repository/mongodb/
│   └── middleware/
```

---

## 10. Entidades sugeridas

### CulinaryCategory

```ts
interface CulinaryCategory {
  id: string;
  name: string;
  slug: string;
  description?: string;
  tag?: string;
  color?: string;
  icon?: string;
  imageUrl?: string;
  order?: number;
  active: boolean;
  recipesCount?: number;
  createdAt: string;
  updatedAt: string;
}
```

### Recipe

```ts
interface Recipe {
  id: string;
  categoryId: string;
  name: string;
  slug: string;
  description?: string;
  prepTimeMinutes?: number;
  cookTimeMinutes?: number;
  servings?: string;
  difficulty?: 'easy' | 'medium' | 'hard';
  status: 'pending' | 'tested' | 'favorite' | 'needs_adjustment';
  tags?: string[];
  imageUrl?: string;
  youtubeUrl?: string;
  sourceUrl?: string;
  ingredients?: string[];
  preparationSteps?: string[];
  utensils?: string[];
  tips?: string;
  substitutions?: string;
  storageInstructions?: string;
  estimatedCost?: number;
  personalRating?: number;
  tested: boolean;
  testedAt?: string;
  notes?: string;
  createdAt: string;
  updatedAt: string;
}
```

### RecipeHistory

```ts
interface RecipeHistory {
  id: string;
  recipeId: string;
  preparedAt: string;
  rating?: number;
  notes?: string;
  adjustments?: string;
  result: 'worked' | 'needs_adjustment' | 'failed';
  createdAt: string;
}
```

---

## 11. Endpoints obrigatórios

### Categorias culinárias

```txt
GET    /api/v1/culinary/categories
GET    /api/v1/culinary/categories/:id
GET    /api/v1/culinary/categories/slug/:slug
POST   /api/v1/culinary/categories
PUT    /api/v1/culinary/categories/:id
DELETE /api/v1/culinary/categories/:id
```

### Receitas

```txt
GET    /api/v1/culinary/recipes
GET    /api/v1/culinary/recipes/:id
GET    /api/v1/culinary/recipes/slug/:categorySlug/:recipeSlug
GET    /api/v1/culinary/categories/:categoryId/recipes
POST   /api/v1/culinary/recipes
PUT    /api/v1/culinary/recipes/:id
DELETE /api/v1/culinary/recipes/:id
```

### Histórico de preparo

```txt
GET    /api/v1/culinary/recipes/:recipeId/history
POST   /api/v1/culinary/recipes/:recipeId/history
PUT    /api/v1/culinary/history/:id
DELETE /api/v1/culinary/history/:id
```

---

## 12. Validações obrigatórias no backend

### Categoria

- `name` obrigatório.
- `name` com mínimo de 2 caracteres.
- `slug` obrigatório.
- `slug` deve estar em kebab-case.
- `slug` deve ser único.
- `description` deve respeitar limite máximo.
- `order` não pode ser negativo.

### Receita

- `categoryId` obrigatório.
- `name` obrigatório.
- `name` com mínimo de 3 caracteres.
- `slug` obrigatório.
- `slug` deve estar em kebab-case.
- `slug` deve ser único dentro da categoria.
- `prepTimeMinutes` não pode ser negativo.
- `cookTimeMinutes` não pode ser negativo.
- `estimatedCost` não pode ser negativo.
- `personalRating` deve estar entre 1 e 5.
- `youtubeUrl`, se preenchida, deve ser uma URL válida do YouTube.
- `sourceUrl`, se preenchida, deve ser uma URL válida.
- `status` deve aceitar apenas valores permitidos.
- `difficulty` deve aceitar apenas valores permitidos.

---

## 13. Frontend — implementação obrigatória

Criar páginas, services, models e componentes seguindo o padrão do projeto.

### Rotas sugeridas

```ts
{
  path: 'culinaria',
  loadComponent: () => import('./pages/culinary/culinary-page.component')
    .then(c => c.CulinaryPageComponent),
},
{
  path: 'culinaria/:categorySlug',
  loadComponent: () => import('./pages/culinary/culinary-category-page.component')
    .then(c => c.CulinaryCategoryPageComponent),
},
{
  path: 'culinaria/:categorySlug/:recipeSlug',
  loadComponent: () => import('./pages/culinary/recipe-detail-page.component')
    .then(c => c.RecipeDetailPageComponent),
}
```

### Estrutura sugerida no frontend

```txt
frontend/src/app/pages/culinary/
├── culinary-page/
├── culinary-category-page/
├── recipe-detail-page/
├── components/
│   ├── culinary-category-card/
│   ├── recipe-card/
│   ├── culinary-category-modal/
│   ├── recipe-modal/
│   └── recipe-history-modal/
├── services/
│   └── culinary.service.ts
├── models/
│   └── culinary.model.ts
└── styles/
```

---

## 14. Service Angular obrigatório

Criar um service para integração real com backend.

Nome sugerido:

```txt
culinary.service.ts
```

Métodos obrigatórios:

```ts
getCategories()
getCategoryBySlug(slug: string)
createCategory(payload)
updateCategory(id: string, payload)
deleteCategory(id: string)

getRecipesByCategory(categoryId: string)
getRecipeBySlug(categorySlug: string, recipeSlug: string)
createRecipe(payload)
updateRecipe(id: string, payload)
deleteRecipe(id: string)

getRecipeHistory(recipeId: string)
createRecipeHistory(recipeId: string, payload)
updateRecipeHistory(id: string, payload)
deleteRecipeHistory(id: string)
```

---

## 15. Regras visuais obrigatórias

Manter consistência com o restante do projeto:

- Mesmo layout da sidebar.
- Mesmo padrão de header.
- Mesmo padrão de breadcrumb.
- Mesmo botão voltar.
- Mesma estrutura de cards.
- Mesmo padrão de modal.
- Mesma tipografia.
- Mesmos espaçamentos.
- Botões dos modais ocupando 50% cada.
- Inputs sem outline verde.
- Focus dos inputs em `#201F25`.
- Botões principais com `#201F25`.
- Botões de cancelar transparentes com borda `#201F25`.
- Cards com botões sempre no footer.
- Descrições longas com `...`.

---

## 16. Sugestão de categorias iniciais para seed

Criar seed opcional com categorias iniciais:

```txt
Café da manhã
Almoço
Jantar
Sobremesas
Massas
Marmitas
Receitas rápidas
Receitas sem glúten
Receitas econômicas
Receitas para congelar
Favoritas
Testadas
```

---

## 17. Sugestão de receitas iniciais para teste

Criar exemplos simples para validar o fluxo:

```txt
Categoria: Massas
Receita: Lasanha de frango

Categoria: Café da manhã
Receita: Panqueca de banana

Categoria: Receitas rápidas
Receita: Omelete simples

Categoria: Sobremesas
Receita: Bolo de cenoura

Categoria: Marmitas
Receita: Frango desfiado com arroz e legumes
```

---

## 18. Critérios de aceite

A implementação será considerada concluída quando:

- O item **Culinária** do menu lateral abrir a página `/culinaria`.
- A página listar categorias vindas do backend.
- O botão **Nova categoria culinária** abrir o modal.
- O modal cadastrar categoria no banco de dados.
- A categoria cadastrada aparecer como card.
- O card permitir editar e excluir categoria.
- Ao clicar em **Abrir categoria**, abrir a página da categoria.
- A página da categoria listar receitas vindas do backend.
- O botão **Nova receita** abrir o modal.
- O modal cadastrar receita no banco de dados.
- A receita aparecer como card.
- O card permitir editar e excluir receita.
- Ao clicar em **Ver receita**, abrir a página de detalhe.
- A página de detalhe exibir os blocos de receita.
- Os links do YouTube e da receita original serem salvos e exibidos corretamente.
- As validações funcionarem no frontend e no backend.
- O layout ficar responsivo.
- Não existir mock fixo quando houver backend disponível.
- O código seguir os padrões já definidos no projeto.

---

## 19. Importante

Não implementar apenas layout estático.

A entrega deve conter:

- Frontend funcional.
- Backend funcional.
- MongoDB persistindo os dados.
- Services Angular integrados com a API.
- Validações.
- CRUD completo.
- Estados de loading, erro e vazio.
- Modais de cadastro e edição.
- Confirmação de exclusão.
- Navegação por breadcrumbs.
- Rotas dinâmicas por slug.

---

## 20. Resultado esperado

Ao final da implementação, o módulo **Culinária** deve permitir que a usuária organize suas receitas como um caderno culinário digital.

A experiência deve ser parecida com o fluxo de **Estudos e Labs**, porém adaptada para receitas, com categorias, cards, modais, detalhes e histórico.

O módulo deve estar pronto para crescer futuramente com recursos como:

- planejamento semanal de refeições;
- lista de compras automática;
- IA para sugerir substituições;
- IA para gerar receitas com ingredientes disponíveis;
- favoritos;
- receitas testadas;
- controle de custo;
- exportação da receita;
- fotos da receita;
- avaliação pessoal.
