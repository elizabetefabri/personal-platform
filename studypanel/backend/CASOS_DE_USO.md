# Casos de Uso — StudyPanel Backend

Documentação dos casos de uso implementados.

---

## Diagrama de Arquitetura

```
┌─────────────────────────────────────────────────────────────┐
│                        Frontend Angular                      │
│                     (localhost:4200)                         │
└───────────────────────────┬─────────────────────────────────┘
                            │ HTTP/JSON
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Middleware CORS                           │
├─────────────────────────────────────────────────────────────┤
│                       Handler HTTP                          │
│                                                             │
│  GET  /api/v1/study-items          → ListStudyItemsUseCase  │
│  POST /api/v1/study-items          → CreateStudyItemUseCase │
│  GET  /api/v1/study-items/{id}     → GetStudyItemUseCase    │
│  PUT  /api/v1/study-items/{id}     → UpdateStudyItemUseCase │
│  DELETE /api/v1/study-items/{id}   → DeleteStudyItemUseCase │
├─────────────────────────────────────────────────────────────┤
│                       Use Cases                             │
│             (regras de negócio e validações)                │
├─────────────────────────────────────────────────────────────┤
│              Repository Interface (contrato)                │
├─────────────────────────────────────────────────────────────┤
│               MongoDB Repository (implementação)            │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                       MongoDB                               │
│              Database: studypanel                           │
│              Collection: study_items                        │
└─────────────────────────────────────────────────────────────┘
```

---

## UC-01 — Criar Item de Estudo

**Caso de Uso:** `CreateStudyItemUseCase`
**Endpoint:** `POST /api/v1/study-items`

### Fluxo Principal

1. Frontend abre modal de cadastro
2. Usuário preenche formulário (nome do curso, status, data, link, imagem)
3. Frontend envia `POST /api/v1/study-items`
4. Handler valida e parseia o body
5. UseCase valida regras de negócio
6. Repository persiste no MongoDB
7. Handler retorna `201 Created` com o item criado
8. Frontend adiciona o item à tabela e exibe link para detalhe

### Dados de Entrada

```json
{
  "section": "cloud",
  "topic": "aws",
  "courseName": "AWS SAA-C03",
  "status": "Em andamento",
  "date": "2024-06-13",
  "url": "https://udemy.com/course/aws",
  "imageUrl": "https://example.com/aws.png"
}
```

### Dados de Saída

```json
{
  "success": true,
  "data": {
    "id": "6676abc123def456",
    "section": "cloud",
    "topic": "aws",
    "courseName": "AWS SAA-C03",
    "status": "Em andamento",
    "date": "2024-06-13",
    "url": "https://udemy.com/course/aws",
    "imageUrl": "https://example.com/aws.png",
    "detailRoute": "/cloud/aws",
    "createdAt": "2024-06-13T10:00:00Z",
    "updatedAt": "2024-06-13T10:00:00Z"
  }
}
```

### Validações

| Campo | Obrigatório | Regra |
|-------|------------|-------|
| section | Sim | Não vazio |
| topic | Sim | Não vazio |
| courseName | Sim | Não vazio |
| status | Sim | Um de: "Não iniciado", "Em andamento", "Concluído", "Pausado" |
| date | Não | Formato YYYY-MM-DD |
| url | Não | — |
| imageUrl | Não | — |

### Regra Automática

`detailRoute` é gerado automaticamente como `/{section}/{topic}`.

---

## UC-02 — Listar Itens de Estudo

**Caso de Uso:** `ListStudyItemsUseCase`
**Endpoint:** `GET /api/v1/study-items`

### Fluxo Principal

1. Página de detalhe carrega
2. Frontend envia `GET /api/v1/study-items?section=cloud&topic=aws`
3. UseCase aplica filtros
4. Repository busca no MongoDB com índices
5. Retorna lista ordenada por data de criação (mais recente primeiro)

### Query Params

| Param | Obrigatório | Descrição |
|-------|------------|-----------|
| section | Não | Filtrar por seção (ex: "cloud") |
| topic | Não | Filtrar por tópico (ex: "aws") |

### Dados de Saída

```json
{
  "success": true,
  "data": [
    {
      "id": "...",
      "section": "cloud",
      "topic": "aws",
      "courseName": "AWS SAA-C03",
      "status": "Em andamento",
      "date": "2024-06-13",
      "url": "...",
      "detailRoute": "/cloud/aws",
      "createdAt": "...",
      "updatedAt": "..."
    }
  ]
}
```

---

## UC-03 — Buscar Item por ID

**Caso de Uso:** `GetStudyItemUseCase`
**Endpoint:** `GET /api/v1/study-items/{id}`

### Fluxo Principal

1. Frontend clica no ícone de olho (view) na tabela
2. Envia `GET /api/v1/study-items/{id}`
3. UseCase valida que o ID não é vazio
4. Repository busca pelo ObjectID no MongoDB
5. Retorna o item ou 404 se não encontrado

---

## UC-04 — Atualizar Item de Estudo

**Caso de Uso:** `UpdateStudyItemUseCase`
**Endpoint:** `PUT /api/v1/study-items/{id}`

### Fluxo Principal

1. Usuário clica em editar na tabela
2. Modal de edição é preenchido com dados atuais
3. Usuário altera campos e salva
4. Frontend envia `PUT /api/v1/study-items/{id}`
5. UseCase verifica que item existe
6. UseCase atualiza campos editáveis
7. Repository persiste no MongoDB com `updated_at` atualizado
8. Frontend atualiza a linha na tabela

### Campos Editáveis

- courseName
- status
- date
- url
- imageUrl

### Campos Imutáveis (gerados automaticamente)

- id
- section
- topic
- detailRoute
- createdAt

---

## UC-05 — Excluir Item de Estudo

**Caso de Uso:** `DeleteStudyItemUseCase`
**Endpoint:** `DELETE /api/v1/study-items/{id}`

### Fluxo Principal

1. Usuário clica em excluir na tabela
2. Confirmação é exibida (frontend)
3. Frontend envia `DELETE /api/v1/study-items/{id}`
4. UseCase verifica que item existe
5. Repository remove do MongoDB
6. Frontend remove a linha da tabela

---

## Modelo de Dados — MongoDB

### Collection: `study_items`

```json
{
  "_id": ObjectId("..."),
  "section": "cloud",
  "topic": "aws",
  "course_name": "AWS SAA-C03",
  "status": "Em andamento",
  "date": "2024-06-13",
  "url": "https://udemy.com/course/aws",
  "image_url": "https://...",
  "detail_route": "/cloud/aws",
  "created_at": ISODate("..."),
  "updated_at": ISODate("...")
}
```

### Índices

```javascript
db.study_items.createIndex({ section: 1 });
db.study_items.createIndex({ topic: 1 });
db.study_items.createIndex({ section: 1, topic: 1 });
db.study_items.createIndex({ status: 1 });
db.study_items.createIndex({ created_at: -1 });
```

---

## Status Permitidos

| Valor | Descrição |
|-------|-----------|
| `Não iniciado` | Curso/recurso ainda não foi começado |
| `Em andamento` | Estudo em progresso |
| `Concluído` | Estudo finalizado |
| `Pausado` | Estudo pausado temporariamente |

---

## Integração com Frontend

### Como o frontend deve usar a API

1. **Ao carregar a página de detalhe** (`/cloud/aws`):
   ```
   GET /api/v1/study-items?section=cloud&topic=aws
   ```

2. **Ao abrir o modal de cadastro e salvar:**
   ```
   POST /api/v1/study-items
   ```
   O `detailRoute` retornado é o link para a página de detalhe.

3. **Ao clicar em editar:**
   ```
   PUT /api/v1/study-items/{id}
   ```

4. **Ao clicar em excluir:**
   ```
   DELETE /api/v1/study-items/{id}
   ```

### URL Base

Desenvolvimento: `http://localhost:8080`
Produção: configurar via `environment.prod.ts` do Angular
