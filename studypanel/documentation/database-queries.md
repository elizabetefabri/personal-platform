# Queries MongoDB — Studypanel

Banco de dados: `studypanel`  
Container Docker: `studypanel-mongo` (porta 27017)

Para conectar via `mongosh`:
```bash
docker exec -it studypanel-mongo mongosh studypanel
```

---

## Resumo das coleções por seção do Sidebar

| Sidebar           | Coleções principais                                                                 |
|-------------------|-------------------------------------------------------------------------------------|
| Dashboard         | (todas — exibe agregações)                                                          |
| Estudos e Labs    | `course_sections`, `course_topics`, `study_items`, `study_notes`, `study_resources`, `study_sessions`, `quiz_questions` |
| Projetos          | `projects`                                                                          |
| Vida Criativa     | `vida_criativa_items`                                                               |
| Painel Financeiro | `financial_records`                                                                 |
| Culinária         | `culinaria_recipes`                                                                 |

---

## Dashboard

O Dashboard agrega dados de todas as coleções para exibir estatísticas gerais.

```js
// Total de itens de estudo
db.study_items.countDocuments()

// Total de projetos
db.projects.countDocuments()

// Total de seções de curso
db.course_sections.countDocuments()

// Total de notas de estudo
db.study_notes.countDocuments()

// Total de recursos
db.study_resources.countDocuments()

// Total de sessões de estudo
db.study_sessions.countDocuments()

// Total de perguntas de quiz
db.quiz_questions.countDocuments()
```

---

## Estudos e Labs

### Seções (course_sections)
```js
// Listar todas as seções
db.course_sections.find().pretty()

// Quantas seções existem
db.course_sections.countDocuments()

// Buscar seção por slug
db.course_sections.findOne({ slug: "backend" })

// Seções ativas ordenadas
db.course_sections.find({ active: true }).sort({ order: 1 })
```

### Tópicos por seção (course_topics)
```js
// Listar todos os tópicos
db.course_topics.find().pretty()

// Quantos tópicos existem
db.course_topics.countDocuments()

// Tópicos de uma seção específica (substituir o slug)
db.course_topics.find({ section_slug: "backend" }).pretty()

// Contagem de tópicos por seção
db.course_topics.aggregate([
  { $group: { _id: "$section_slug", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])
```

### Itens de estudo (study_items)
```js
// Listar todos os itens
db.study_items.find().pretty()

// Quantos itens existem
db.study_items.countDocuments()

// Itens por seção
db.study_items.find({ section: "backend" }).pretty()

// Contagem de itens por seção
db.study_items.aggregate([
  { $group: { _id: "$section", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])
```

### Notas de estudo (study_notes)
```js
// Todas as notas
db.study_notes.find().pretty()

// Quantas notas existem
db.study_notes.countDocuments()

// Notas de um item específico (substituir o ID)
db.study_notes.find({ study_item_id: ObjectId("6a2dc606d11c04d0ab5e6f41") }).pretty()

// Contagem de notas por item
db.study_notes.aggregate([
  { $group: { _id: "$study_item_id", total: { $sum: 1 } } }
])
```

### Recursos de estudo (study_resources)
```js
// Todos os recursos
db.study_resources.find().pretty()

// Quantos recursos existem
db.study_resources.countDocuments()

// Recursos de um item específico
db.study_resources.find({ study_item_id: ObjectId("6a2dc606d11c04d0ab5e6f41") }).pretty()

// Recursos por tipo (link, video, artigo, etc.)
db.study_resources.aggregate([
  { $group: { _id: "$type", total: { $sum: 1 } } }
])
```

### Sessões de estudo (study_sessions)
```js
// Todas as sessões
db.study_sessions.find().pretty()

// Quantas sessões existem
db.study_sessions.countDocuments()

// Sessões de um item específico
db.study_sessions.find({ study_item_id: ObjectId("6a2dc606d11c04d0ab5e6f41") }).pretty()

// Tempo total estudado (em minutos) por item
db.study_sessions.aggregate([
  { $group: { _id: "$study_item_id", total_minutes: { $sum: "$duration_minutes" } } },
  { $sort: { total_minutes: -1 } }
])
```

### Perguntas de Quiz (quiz_questions)
```js
// Todas as perguntas
db.quiz_questions.find().pretty()

// Quantas perguntas existem
db.quiz_questions.countDocuments()

// Perguntas por tópico
db.quiz_questions.find({ topic: "aws" }).pretty()

// Contagem de perguntas por tópico
db.quiz_questions.aggregate([
  { $group: { _id: "$topic", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])
```

---

## Projetos (projects)

```js
// Listar todos os projetos
db.projects.find().pretty()

// Quantos projetos existem
db.projects.countDocuments()

// Apenas projetos pessoais
db.projects.find({ type: "pessoal" }).pretty()

// Apenas projetos profissionais
db.projects.find({ type: "profissional" }).pretty()

// Contagem por tipo
db.projects.aggregate([
  { $group: { _id: "$type", total: { $sum: 1 } } }
])

// Projetos ativos ordenados
db.projects.find({ active: true }).sort({ order: 1 }).pretty()
```

---

## Vida Criativa (vida_criativa_items)

```js
// Listar todos os itens
db.vida_criativa_items.find().pretty()

// Quantos itens existem
db.vida_criativa_items.countDocuments()

// Filtrar por categoria (Design, Música, Escrita, Fotografia, Arte, Vídeo, Outro)
db.vida_criativa_items.find({ category: "Design" }).pretty()

// Filtrar por status (Ideia, Em andamento, Concluído, Pausado)
db.vida_criativa_items.find({ status: "Em andamento" }).pretty()

// Contagem por categoria
db.vida_criativa_items.aggregate([
  { $group: { _id: "$category", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])

// Contagem por status
db.vida_criativa_items.aggregate([
  { $group: { _id: "$status", total: { $sum: 1 } } }
])
```

---

## Painel Financeiro (financial_records)

```js
// Listar todos os registros
db.financial_records.find().pretty()

// Quantos registros existem
db.financial_records.countDocuments()

// Apenas receitas
db.financial_records.find({ type: "receita" }).pretty()

// Apenas despesas
db.financial_records.find({ type: "despesa" }).pretty()

// Contagem e soma por tipo
db.financial_records.aggregate([
  {
    $group: {
      _id: "$type",
      total_registros: { $sum: 1 },
      total_valor: { $sum: "$amount" }
    }
  }
])

// Registros por categoria
db.financial_records.aggregate([
  { $group: { _id: "$category", total: { $sum: 1 }, valor: { $sum: "$amount" } } },
  { $sort: { valor: -1 } }
])

// Registros de um mês específico (ex: junho 2026)
db.financial_records.find({
  date: {
    $gte: ISODate("2026-06-01T00:00:00Z"),
    $lt:  ISODate("2026-07-01T00:00:00Z")
  }
}).sort({ date: -1 })

// Saldo geral (receitas - despesas)
db.financial_records.aggregate([
  {
    $group: {
      _id: "$type",
      total: { $sum: "$amount" }
    }
  }
])
```

---

## Culinária (culinaria_recipes)

```js
// Listar todas as receitas
db.culinaria_recipes.find().pretty()

// Quantas receitas existem
db.culinaria_recipes.countDocuments()

// Filtrar por categoria
db.culinaria_recipes.find({ category: "Sobremesas" }).pretty()

// Filtrar por dificuldade (Fácil, Médio, Difícil)
db.culinaria_recipes.find({ difficulty: "Fácil" }).pretty()

// Receitas por tempo de preparo (até 30 min)
db.culinaria_recipes.find({ prep_time: { $lte: 30 } }).sort({ prep_time: 1 })

// Contagem por dificuldade
db.culinaria_recipes.aggregate([
  { $group: { _id: "$difficulty", total: { $sum: 1 } } }
])

// Contagem por categoria
db.culinaria_recipes.aggregate([
  { $group: { _id: "$category", total: { $sum: 1 } } },
  { $sort: { total: -1 } }
])

// Receitas ativas
db.culinaria_recipes.find({ active: true }).sort({ created_at: -1 })
```

---

## Snapshot atual (2026-06-27)

| Coleção              | Documentos |
|----------------------|-----------|
| study_items          | 8         |
| study_notes          | 2         |
| study_resources      | 3         |
| study_sessions       | 1         |
| quiz_questions       | 5         |
| projects             | 0         |
| course_sections      | 0         |
| vida_criativa_items  | 0 (nova)  |
| financial_records    | 0 (nova)  |
| culinaria_recipes    | 0 (nova)  |
