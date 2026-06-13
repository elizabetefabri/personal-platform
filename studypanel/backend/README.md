# StudyPanel — Backend

API REST em Go para gerenciar itens de estudo do StudyPanel.

## Stack

- **Go 1.22** — net/http com pattern matching nativo
- **MongoDB 7.0** — banco de dados NoSQL
- **Mongo Express** — interface web para administrar o banco
- **Docker Compose** — orquestração local

## Pré-requisitos

- Go 1.22+
- Docker e Docker Compose

## Setup Rápido

```bash
# Clonar e entrar no diretório
cd personal-platform/studypanel/backend

# Configurar variáveis de ambiente
cp .env.example .env

# Subir toda a stack (API + MongoDB + Mongo Express)
make docker-up
```

Acesse:
- **API:** http://localhost:8080
- **Mongo Express:** http://localhost:8081 (admin / studypanel123)

## Desenvolvimento Local

```bash
# Instalar dependências
make tidy

# Rodar localmente (MongoDB via Docker)
make docker-up     # sobe só o mongo
make run           # roda a API localmente
```

## Testes

```bash
# Rodar todos os testes
make test

# Rodar com cobertura
make test-cover
# Abre coverage.html com o relatório visual
```

## API Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/health` | Health check |
| GET | `/api/v1/study-items` | Listar itens (filtros: `?section=&topic=`) |
| POST | `/api/v1/study-items` | Criar item |
| GET | `/api/v1/study-items/{id}` | Buscar por ID |
| PUT | `/api/v1/study-items/{id}` | Atualizar item |
| DELETE | `/api/v1/study-items/{id}` | Excluir item |

## Exemplo de Uso

```bash
# Health check
curl http://localhost:8080/health

# Listar todos
curl http://localhost:8080/api/v1/study-items

# Filtrar por seção/tópico
curl "http://localhost:8080/api/v1/study-items?section=cloud&topic=aws"

# Criar item
curl -X POST http://localhost:8080/api/v1/study-items \
  -H "Content-Type: application/json" \
  -d '{
    "section": "cloud",
    "topic": "aws",
    "courseName": "AWS SAA-C03",
    "status": "Em andamento",
    "date": "2024-06-13",
    "url": "https://udemy.com/course/aws"
  }'

# Atualizar
curl -X PUT http://localhost:8080/api/v1/study-items/ID \
  -H "Content-Type: application/json" \
  -d '{"courseName": "AWS SAA", "status": "Concluído", "date": "2024-07-01", "url": ""}'

# Excluir
curl -X DELETE http://localhost:8080/api/v1/study-items/ID
```

## Estrutura

```
backend/
├── cmd/server/main.go        → Entrada da aplicação
├── config/                   → Configuração via env vars
├── internal/
│   ├── domain/               → Entidades e interfaces
│   ├── usecase/              → Regras de negócio + testes
│   ├── handler/              → HTTP handlers + testes
│   ├── repository/mongodb/   → Implementação MongoDB
│   └── middleware/           → CORS
├── pkg/response/             → Helpers de resposta
├── docker/mongo-init.js      → Init do MongoDB
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── .env.example
```

Veja [PADROES.md](PADROES.md) e [CASOS_DE_USO.md](CASOS_DE_USO.md) para detalhes.
