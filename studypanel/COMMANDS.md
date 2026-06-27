# Comandos do Projeto — Studypanel

## Visão Geral

| Serviço         | URL                            | Descrição                          |
|-----------------|--------------------------------|------------------------------------|
| Frontend        | http://localhost:4200          | Angular — interface web            |
| API (Backend)   | http://localhost:8090          | Go REST API                        |
| Mongo Express   | http://localhost:8081          | UI visual para o MongoDB           |
| MongoDB         | mongodb://localhost:27017      | Banco de dados (direto)            |

Credenciais Mongo Express: `admin` / `studypanel123`

---

## Backend (Go + MongoDB via Docker)

```bash
cd backend

# Subir API + MongoDB + Mongo Express (rebuild ao alterar código Go)
make docker-up

# Desligar tudo
make docker-down

# Ver logs da API em tempo real
make docker-logs

# Rodar localmente sem Docker (MongoDB precisa estar rodando separado)
make run

# Build do binário
make build

# Verificar se está funcionando
curl http://localhost:8090/health
```

### Quando reconstruir o backend

Sempre que alterar código Go (handlers, usecases, entities, etc.), é preciso rebuild:
```bash
cd backend
make docker-down && make docker-up
```

### Status dos containers
```bash
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

### Logs individuais
```bash
docker logs studypanel-api -f
docker logs studypanel-mongo -f
docker logs studypanel-mongo-express -f
```

---

## Frontend (Angular)

```bash
cd frontend

# Instalar dependências (apenas na primeira vez ou após atualizar package.json)
npm install

# Iniciar servidor de desenvolvimento
npx ng serve
# ou
npm start
# Acessa em http://localhost:4200

# Build de produção
npx ng build --configuration production
```

---

## Banco de Dados — Comandos Rápidos

```bash
# Conectar via mongosh no container
docker exec -it studypanel-mongo mongosh studypanel

# Conectar com mongosh local (se instalado)
mongosh mongodb://localhost:27017/studypanel
```

---

## Endpoints da API

### Health Check
```
GET  http://localhost:8090/health
```

### Estudos e Labs — Seções (course_sections)
```
GET    http://localhost:8090/api/v1/course-sections
POST   http://localhost:8090/api/v1/course-sections
GET    http://localhost:8090/api/v1/course-sections/:id
PUT    http://localhost:8090/api/v1/course-sections/:id
DELETE http://localhost:8090/api/v1/course-sections/:id
```

### Estudos e Labs — Tópicos (course_topics)
```
GET    http://localhost:8090/api/v1/course-topics?sectionId=<id>
POST   http://localhost:8090/api/v1/course-topics
PUT    http://localhost:8090/api/v1/course-topics/:id
DELETE http://localhost:8090/api/v1/course-topics/:id
```

### Estudos e Labs — Itens de Estudo (study_items)
```
GET    http://localhost:8090/api/v1/study-items?section=<slug>
POST   http://localhost:8090/api/v1/study-items
GET    http://localhost:8090/api/v1/study-items/:id
PUT    http://localhost:8090/api/v1/study-items/:id
DELETE http://localhost:8090/api/v1/study-items/:id
```

### Estudos e Labs — Notas (study_notes)
```
GET    http://localhost:8090/api/v1/study-notes?studyItemId=<id>
POST   http://localhost:8090/api/v1/study-notes
PUT    http://localhost:8090/api/v1/study-notes/:id
DELETE http://localhost:8090/api/v1/study-notes/:id
```

### Estudos e Labs — Recursos (study_resources)
```
GET    http://localhost:8090/api/v1/study-resources?studyItemId=<id>
POST   http://localhost:8090/api/v1/study-resources
PUT    http://localhost:8090/api/v1/study-resources/:id
DELETE http://localhost:8090/api/v1/study-resources/:id
```

### Estudos e Labs — Sessões (study_sessions)
```
GET    http://localhost:8090/api/v1/study-sessions?studyItemId=<id>
POST   http://localhost:8090/api/v1/study-sessions
PUT    http://localhost:8090/api/v1/study-sessions/:id
DELETE http://localhost:8090/api/v1/study-sessions/:id
```

### Estudos e Labs — Quiz (quiz_questions)
```
GET    http://localhost:8090/api/v1/quiz-questions?topic=<slug>
POST   http://localhost:8090/api/v1/quiz-questions
PUT    http://localhost:8090/api/v1/quiz-questions/:id
DELETE http://localhost:8090/api/v1/quiz-questions/:id
```

### Projetos (projects)
```
GET    http://localhost:8090/api/v1/projects?type=pessoal
GET    http://localhost:8090/api/v1/projects?type=profissional
POST   http://localhost:8090/api/v1/projects
GET    http://localhost:8090/api/v1/projects/:id
PUT    http://localhost:8090/api/v1/projects/:id
DELETE http://localhost:8090/api/v1/projects/:id
```

### Vida Criativa (vida_criativa_items)
```
GET    http://localhost:8090/api/v1/vida-criativa?category=Design
POST   http://localhost:8090/api/v1/vida-criativa
GET    http://localhost:8090/api/v1/vida-criativa/:id
PUT    http://localhost:8090/api/v1/vida-criativa/:id
DELETE http://localhost:8090/api/v1/vida-criativa/:id
```

### Painel Financeiro (financial_records)
```
GET    http://localhost:8090/api/v1/financial-records?type=receita
GET    http://localhost:8090/api/v1/financial-records?type=despesa
POST   http://localhost:8090/api/v1/financial-records
GET    http://localhost:8090/api/v1/financial-records/:id
PUT    http://localhost:8090/api/v1/financial-records/:id
DELETE http://localhost:8090/api/v1/financial-records/:id
```

### Culinária (culinaria_recipes)
```
GET    http://localhost:8090/api/v1/culinaria/recipes?category=Sobremesas
POST   http://localhost:8090/api/v1/culinaria/recipes
GET    http://localhost:8090/api/v1/culinaria/recipes/:id
PUT    http://localhost:8090/api/v1/culinaria/recipes/:id
DELETE http://localhost:8090/api/v1/culinaria/recipes/:id
```

---

## Queries MongoDB (mongosh)

Conectar: `docker exec -it studypanel-mongo mongosh studypanel`

### Dashboard — visão geral
```js
db.study_items.countDocuments()
db.projects.countDocuments()
db.course_sections.countDocuments()
db.study_notes.countDocuments()
db.study_resources.countDocuments()
db.study_sessions.countDocuments()
db.quiz_questions.countDocuments()
db.vida_criativa_items.countDocuments()
db.financial_records.countDocuments()
db.culinaria_recipes.countDocuments()
```

### Estudos e Labs
```js
// Todas as seções ativas
db.course_sections.find({ active: true }).sort({ order: 1 }).pretty()

// Tópicos por seção
db.course_topics.find({ section_slug: "backend" }).pretty()

// Itens por seção
db.study_items.find({ section: "backend" }).pretty()

// Contagem de itens por seção
db.study_items.aggregate([
  { $group: { _id: "$section", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])

// Notas, recursos e sessões de um item (substituir ID)
db.study_notes.find({ study_item_id: ObjectId("ITEM_ID") }).pretty()
db.study_resources.find({ study_item_id: ObjectId("ITEM_ID") }).pretty()
db.study_sessions.find({ study_item_id: ObjectId("ITEM_ID") }).pretty()

// Perguntas de quiz por tópico
db.quiz_questions.find({ topic: "aws" }).pretty()
```

### Projetos
```js
db.projects.find({ type: "pessoal" }).sort({ order: 1 }).pretty()
db.projects.find({ type: "profissional" }).sort({ order: 1 }).pretty()
db.projects.aggregate([{ $group: { _id: "$type", total: { $sum: 1 } } }])
```

### Vida Criativa
```js
db.vida_criativa_items.find().sort({ order: 1 }).pretty()
db.vida_criativa_items.find({ status: "Em andamento" }).pretty()
db.vida_criativa_items.aggregate([
  { $group: { _id: "$category", total: { $sum: 1 } } }
])
```

### Painel Financeiro
```js
db.financial_records.find({ type: "receita" }).sort({ date: -1 }).pretty()
db.financial_records.find({ type: "despesa" }).sort({ date: -1 }).pretty()

// Saldo por tipo
db.financial_records.aggregate([
  { $group: { _id: "$type", total_valor: { $sum: "$amount" }, qtd: { $sum: 1 } } }
])

// Registros de um mês (ex: junho 2026)
db.financial_records.find({
  date: { $gte: ISODate("2026-06-01"), $lt: ISODate("2026-07-01") }
}).sort({ date: -1 })
```

### Culinária
```js
db.culinaria_recipes.find({ active: true }).sort({ created_at: -1 }).pretty()
db.culinaria_recipes.find({ difficulty: "Fácil" }).pretty()
db.culinaria_recipes.find({ prep_time: { $lte: 30 } }).sort({ prep_time: 1 })
db.culinaria_recipes.aggregate([
  { $group: { _id: "$category", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])
```

---

## Snapshot atual do banco (2026-06-27)

| Coleção               | Documentos | Sidebar               |
|-----------------------|-----------|------------------------|
| study_items           | 8         | Estudos e Labs         |
| study_notes           | 2         | Estudos e Labs         |
| study_resources       | 3         | Estudos e Labs         |
| study_sessions        | 1         | Estudos e Labs         |
| quiz_questions        | 5         | Estudos e Labs         |
| course_sections       | 1         | Estudos e Labs         |
| projects              | 0         | Projetos               |
| vida_criativa_items   | 0 (nova)  | Vida Criativa          |
| financial_records     | 0 (nova)  | Painel Financeiro      |
| culinaria_recipes     | 0 (nova)  | Culinária              |
