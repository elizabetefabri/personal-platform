# Cronograma 60 Dias — AWS SAA-C03

**Estudo:** 1 hora por dia | **Meta:** Aprovação com >850 pontos

---

## Semana 1 — Fundamentos AWS e IAM

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 1 | Fundamentos AWS | Regiões, AZs, Edge Locations, modelos (IaaS/PaaS/SaaS) | 🔴 Alta |
| 2 | IAM — Parte 1 | Usuários, Grupos, Políticas, Roles | 🔴 Muito Alta |
| 3 | IAM — Parte 2 | Permissões, MFA, Access Keys, STS | 🔴 Muito Alta |
| 4 | Billing e Pricing | Free Tier, Calculadora de custos, Cost Explorer | 🟠 Alta |
| 5 | Revisão + Quiz | IAM completo — 10 questões estilo prova | 🔴 Alta |
| 6 | Lab Prático | Lab 01 — IAM: criar usuários, grupos, políticas | 🔴 Alta |
| 7 | Revisão Geral | Revisão semana 1 + mapa mental IAM | 🔴 Alta |

---

## Semana 2 — VPC e Redes

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 8 | VPC — Parte 1 | CIDR, Subnets públicas/privadas, Internet Gateway | 🔴 Muito Alta |
| 9 | VPC — Parte 2 | Route Tables, NAT Gateway, Security Groups, NACLs | 🔴 Muito Alta |
| 10 | VPC — Parte 3 | VPC Peering, VPN, Direct Connect, Transit Gateway | 🟠 Alta |
| 11 | VPC — Parte 4 | Endpoints, PrivateLink, Flow Logs | 🟠 Alta |
| 12 | Revisão + Quiz | VPC completo — 10 questões estilo prova | 🔴 Alta |
| 13 | Lab Prático | Lab 02 — VPC: criar subnets, gateway, route tables | 🔴 Alta |
| 14 | Revisão + Revisão Espaçada | VPC + IAM — questões mistas | 🔴 Alta |

---

## Semana 3 — EC2, Auto Scaling e ELB

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 15 | EC2 — Parte 1 | Tipos de instâncias, AMI, keypairs, EBS | 🔴 Muito Alta |
| 16 | EC2 — Parte 2 | On-Demand, Reserved, Spot, Dedicated | 🔴 Muito Alta |
| 17 | EC2 — Parte 3 | User Data, Placement Groups, Hibernate | 🟠 Alta |
| 18 | Auto Scaling | ASG, Políticas de escalonamento, Launch Templates | 🔴 Muito Alta |
| 19 | ELB | ALB, NLB, CLB, Target Groups, Health Checks | 🔴 Muito Alta |
| 20 | Lab Prático | Lab 03 — EC2: lançar, conectar via SSH, EBS | 🔴 Alta |
| 21 | Revisão + Quiz | EC2 + ASG + ELB — 10 questões estilo prova | 🔴 Alta |

---

## Semana 4 — Storage: S3, EBS, EFS, CloudFront

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 22 | S3 — Parte 1 | Buckets, objetos, classes de storage, versionamento | 🔴 Muito Alta |
| 23 | S3 — Parte 2 | Bucket Policies, ACLs, Presigned URLs, CORS | 🔴 Muito Alta |
| 24 | S3 — Parte 3 | Lifecycle, Replication, S3 Select, Event Notifications | 🟠 Alta |
| 25 | EBS, EFS, FSx | Volumes, tipos de disco, EFS Multi-AZ, FSx for Windows | 🟡 Média |
| 26 | CloudFront | CDN, Origins, Behaviors, Cache Policies, OAC | 🟠 Alta |
| 27 | Storage Gateway | File, Volume, Tape Gateway | 🟡 Média |
| 28 | Revisão + Quiz | Storage completo — 10 questões estilo prova | 🔴 Alta |

---

## Semana 5 — Banco de Dados + Primeiro Simulado

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 29 | RDS — Parte 1 | Engines suportadas, Multi-AZ, Read Replicas | 🔴 Muito Alta |
| 30 | RDS — Parte 2 | Backups, snapshots, encryption, maintenance | 🔴 Muito Alta |
| 31 | Aurora | Clusters, Global Database, Serverless v2 | 🔴 Muito Alta |
| 32 | DynamoDB | Tabelas, partitions, GSI, LSI, DynamoDB Streams | 🔴 Muito Alta |
| 33 | ElastiCache | Redis vs Memcached, caching patterns | 🟡 Média |
| 34 | Lab Prático | Lab 06 — RDS Multi-AZ + Read Replica | 🔴 Alta |
| 35 | **Simulado 1** | 20 questões — Semanas 1 a 5 | 🔴 Muito Alta |

---

## Semana 6 — Serverless + Mensageria

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 36 | Lambda — Parte 1 | Functions, triggers, layers, concorrência | 🔴 Muito Alta |
| 37 | Lambda — Parte 2 | Cold start, timeout, memory, VPC Lambda | 🟠 Alta |
| 38 | API Gateway | REST API, HTTP API, WebSocket, stages, auth | 🟠 Alta |
| 39 | SQS | Standard vs FIFO, DLQ, visibility timeout, batching | 🟠 Alta |
| 40 | SNS | Topics, fan-out, filtros, SMS, Email | 🟠 Alta |
| 41 | EventBridge | Event bus, Rules, Patterns, Scheduled events | 🟠 Alta |
| 42 | **Simulado 2** | 30 questões — Semanas 1 a 6 | 🔴 Muito Alta |

---

## Semana 7 — Containers + IaC + Well-Architected

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 43 | ECS | Task Definitions, Services, Clusters, Launch Types | 🟠 Alta |
| 44 | EKS + Fargate | Managed Kubernetes, Fargate profiles | 🟠 Alta |
| 45 | Route 53 | Routing Policies, Health Checks, Alias Records | 🟠 Alta |
| 46 | CloudFormation | Templates, Stacks, StackSets, Drift Detection | 🟡 Média |
| 47 | Well-Architected | 6 pilares, Trade-offs, melhores práticas | 🔴 Muito Alta |
| 48 | Disaster Recovery | RTO/RPO, Backup, Pilot Light, Warm Standby, Multi-site | 🔴 Muito Alta |
| 49 | **Simulado 3** | 40 questões — todos os domínios | 🔴 Muito Alta |

---

## Semana 8 — Revisão Final + Simulado Completo

| Dia | Tema | Serviço/Conteúdo | Prioridade |
|-----|------|------------------|-----------|
| 50 | Segurança Avançada | KMS, Secrets Manager, ACM, WAF, Shield, GuardDuty | 🔴 Muito Alta |
| 51 | Alta Disponibilidade | Patterns, multi-AZ, multi-region, failover | 🔴 Muito Alta |
| 52 | Migração AWS | DMS, SMS, Migration Hub, Snowball | 🟡 Média |
| 53 | Cost Optimization | Trusted Advisor, Cost Explorer, Savings Plans | 🟠 Alta |
| 54 | Observabilidade | CloudWatch, CloudTrail, X-Ray, Config | 🟠 Alta |
| 55 | Revisão Pontos Fracos | Revisar erros dos simulados anteriores | 🔴 Muito Alta |
| 56 | Revisão Pontos Fracos | Reforço nos domínios com <80% de acerto | 🔴 Muito Alta |
| 57 | **Simulado Final 1** | 65 questões — estilo prova real | 🔴 Muito Alta |
| 58 | Revisão do Simulado | Análise detalhada das questões erradas | 🔴 Muito Alta |
| 59 | **Simulado Final 2** | 65 questões — novo simulado completo | 🔴 Muito Alta |
| 60 | Revisão Final | Flash cards, pontos críticos, confiança | 🔴 Muito Alta |

---

## Distribuição de Tempo por Domínio

| Domínio | Peso na Prova | Semanas |
|---------|--------------|---------|
| Design de Arquiteturas Resilientes | 26% | 2, 3, 7 |
| Design de Arquiteturas de Alto Desempenho | 24% | 3, 4, 5, 6 |
| Design de Arquiteturas Seguras | 30% | 1, 4, 8 |
| Design de Arquiteturas com Otimização de Custos | 20% | 8 |

---

## Regra de Ajuste do Plano

Se errar mais de **20% em algum tema**, adicionar revisão adicional no dia seguinte antes de avançar.

Domínios com menos de **80% de acerto** nos quizzes → revisar antes do simulado final.
