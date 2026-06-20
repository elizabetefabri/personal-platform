# Comandos e Regras — Backend StudyPanel

> Go 1.22 · MongoDB · godotenv · mongo-driver v1.15

---

## Setup Inicial

```bash
# Na pasta backend/
go mod download

# Copiar variáveis de ambiente
cp .env.example .env
# Editar .env com string de conexão MongoDB

# Subir com Docker Compose (recomendado)
docker compose up -d

# Ou rodar direto
go run ./cmd/main.go
# → http://localhost:8080
```

---

## Desenvolvimento

```bash
# Build
go build ./...

# Rodar localmente
go run ./cmd/main.go

# Rodar com live reload (air)
air

# Verificar código
go vet ./...
go fmt ./...

# Lint (golangci-lint)
golangci-lint run

# Testes
go test ./...
go test -v ./...
go test -cover ./...

# Testes em pacote específico
go test ./internal/...
```

---

## Docker

```bash
# Build da imagem
docker build -t studypanel-backend .

# Subir stack completa (app + MongoDB)
docker compose up -d

# Logs
docker compose logs -f

# Parar
docker compose down

# Parar e remover volumes (cuidado: apaga dados)
docker compose down -v
```

---

## Makefile

```bash
# Atalhos definidos no Makefile
make build       # go build
make run         # go run
make test        # go test ./...
make lint        # golangci-lint run
make docker-up   # docker compose up -d
make docker-down # docker compose down
```

---

## Estrutura de Diretórios

```
backend/
├── cmd/
│   └── main.go              # Entrypoint
├── internal/
│   ├── domain/              # Entidades (structs) e interfaces de repositório
│   ├── handlers/            # HTTP handlers
│   ├── middleware/          # Middlewares (CORS, auth, logging)
│   ├── repository/          # Implementações MongoDB
│   └── service/             # Regras de negócio
├── pkg/
│   └── database/            # Conexão MongoDB
├── config/                  # Carregamento de .env
├── docker/
│   └── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
└── go.sum
```

---

## Variáveis de Ambiente

```env
# .env
PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DB=studypanel
```

---

## API REST

### Base URL

`http://localhost:8080/api/v1`

### Endpoints (StudyItems)

| Método  | Rota                                | Descrição                        |
| ------- | ----------------------------------- | -------------------------------- |
| `GET`   | `/study-items`                      | Listar todos                     |
| `GET`   | `/study-items?section=X&topic=Y`    | Filtrar por seção/tópico         |
| `GET`   | `/study-items/:id`                  | Buscar por ID                    |
| `POST`  | `/study-items`                      | Criar item                       |
| `PUT`   | `/study-items/:id`                  | Atualizar item                   |
| `DELETE`| `/study-items/:id`                  | Remover item                     |

### Corpo de StudyItem

```json
{
  "section": "backend",
  "topic": "node",
  "title": "Event Loop",
  "status": "studying",
  "notes": "...",
  "createdAt": "2026-06-20T00:00:00Z",
  "updatedAt": "2026-06-20T00:00:00Z"
}
```

### Status válidos

```
studying | completed | planned
```

---

## Regras do Projeto

### Go

- **Go 1.22** — não usar features de versões mais novas sem atualizar `go.mod`
- **Erros explícitos** — sem `panic` em código de produção; propagar errors
- **Context propagation** — sempre receber `context.Context` como primeiro argumento em funções que acessam banco
- **Estrutura hexagonal** — handlers não acessam banco diretamente; sempre via service → repository
- **MongoDB** — usar `mongo-driver` v1.15; não usar ORMs

### CORS

- O frontend roda em `http://localhost:4200` — CORS deve permitir esta origem em desenvolvimento
- Em produção, restringir à URL de deploy

### Segurança

- **Não logar** dados sensíveis (tokens, senhas, connection strings)
- **Validar** todos os campos de entrada antes de persistir
- **MongoDB injection**: usar `bson.D` com chaves e valores separados, nunca interpolação de string em queries

### Migrations

- MongoDB é schemaless, mas mudanças de estrutura de documento devem ser documentadas aqui

---

## Checklist de PR (Backend)

- [ ] `go build ./...` sem erros
- [ ] `go vet ./...` sem warnings
- [ ] `go test ./...` — todos os testes passando
- [ ] CORS configurado para ambiente correto
- [ ] Variáveis de ambiente documentadas em `.env.example`
- [ ] Nenhum segredo commitado (verificar `go.sum` e `.env`)
- [ ] Docker Compose funciona localmente
