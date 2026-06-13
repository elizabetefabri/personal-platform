# StudyPanel

Plataforma de gerenciamento de estudos pessoal — trilhas de aprendizado, rastreamento de cursos e objetivos técnicos.

## Estrutura

```
studypanel/
├── frontend/          → Angular 21 + PrimeNG (interface web)
├── backend/           → Go 1.22 + MongoDB (API REST)
└── documentation/     → Documentação e materiais de estudo
    ├── agentes-professores/  → Prompts de agentes IA para cada tema
    └── aws/                  → Material completo para AWS SAA-C03
```

## Frontend

- **Tech:** Angular 21, TypeScript, SCSS, PrimeNG
- **Rotas:** `/cloud`, `/frontend`, `/devops`, `/backend`, e demais seções
- **Funcionalidade:** Cards de tópicos, tabela de cursos com CRUD, modais de cadastro

### Rodar o frontend

```bash
cd frontend
npm install
ng serve
# Acessa em http://localhost:4200
```

## Backend

- **Tech:** Go 1.22, MongoDB, Docker, Mongo Express
- **Endpoints:** CRUD para `study_items`
- **Clean Architecture:** entity → usecase → handler → repository

### Rodar o backend

```bash
cd backend
make docker-up    # API + MongoDB + Mongo Express
# API em http://localhost:8080
# Mongo Express em http://localhost:8081
```

## Documentação de Estudo

### Agentes Professores

Prompts completos para usar com IA como professor particular:

| Arquivo | Tema |
|---------|------|
| `professor-aws-solutions-architect-saa-c03.md` | AWS SAA-C03 |
| `professor-kubernetes-cka.md` | Kubernetes CKA |
| `professor-terraform-infrastructure.md` | Terraform |
| `professor-datadog-observability.md` | Datadog e Observabilidade |
| `professor-platform-engineering.md` | Platform Engineering |

### Plano AWS SAA-C03

Material organizado para a certificação AWS:

```
documentation/aws/
├── plano-de-estudos/    → Cronograma 60 dias, matriz de probabilidade
├── servicos/            → IAM, VPC, EC2, S3, RDS, Lambda, DynamoDB...
├── labs/                → Laboratórios práticos por serviço
├── simulados/           → Questões estilo prova real
└── recursos/            → Links, cursos, whitepapers
```

## Padrões do Projeto

- [PADROES.md](PADROES.md) — Padrões técnicos deste projeto
- [WIP.md](WIP.md) — Status atual e próximos passos
