# RDS e Aurora — Banco de Dados Relacional

**Incidência:** 🔴 Muito Alta | **Domínio:** Banco de Dados

---

## RDS — Relational Database Service

### Engines Suportadas
- MySQL
- PostgreSQL
- MariaDB
- Oracle
- Microsoft SQL Server
- (Aurora tem engine própria)

### Multi-AZ
- Cria réplica síncrona em outra AZ
- Failover automático (~1-2 min) sem intervenção manual
- **Uso: Alta Disponibilidade** — não é para leitura

### Read Replicas
- Cópias assíncronas para consultas de leitura
- Até 15 read replicas por instância
- Podem ficar em regiões diferentes (Cross-Region Read Replica)
- **Uso: Performance / escalar leitura** — não é para HA

### Diferença crítica:
| | Multi-AZ | Read Replica |
|--|---------|-------------|
| Objetivo | Alta Disponibilidade | Escalar leitura |
| Sincronização | Síncrona | Assíncrona |
| Failover | Automático | Manual (pode ser promovida) |
| Endpoint | Um (muda no failover) | Endpoint próprio por réplica |

---

## Aurora

Aurora é a engine de banco da AWS — compatível com MySQL e PostgreSQL, mas reconstruída para a nuvem.

### Por que Aurora é diferente?
- **Storage auto-escalável**: começa em 10GB, cresce até 128TB automaticamente
- **6 cópias dos dados** distribuídas em 3 AZs
- **Até 15 read replicas** com failover automático em ~30 segundos
- **Performance 3x PostgreSQL e 5x MySQL** (benchmark Aurora)
- **Alta disponibilidade nativa** sem configuração extra

### Aurora Cluster
```
┌─────────────────────────────────────┐
│           Aurora Cluster            │
│                                     │
│  Writer Endpoint → Instância Master │
│                         │           │
│  Reader Endpoint → Réplicas (1-15)  │
│                                     │
│  Volume compartilhado (AZ1/AZ2/AZ3) │
└─────────────────────────────────────┘
```

### Aurora Serverless v2
- Escala automaticamente entre 0.5 e 128 ACUs (Aurora Capacity Units)
- Ideal para cargas imprevisíveis, dev/test, multi-tenant
- Para completamente quando não há uso (cost saving)

### Aurora Global Database
- Replicação em até 5 regiões
- Latência de replicação <1 segundo
- Failover cross-region em <1 minuto
- Ideal para aplicações globais com disaster recovery

---

## Backups e Snapshots

| Tipo | Retenção | Automático |
|------|---------|-----------|
| Automated Backup | 1-35 dias | Sim (janela de manutenção) |
| Manual Snapshot | Indefinido | Não (usuário cria) |
| Aurora Backtrack | Até 72h | Não (voltar no tempo sem restore) |

### Point-in-Time Recovery
- Restaura para qualquer segundo dentro do período de retenção
- Cria uma nova instância RDS (não modifica a atual)

---

## Segurança

- **Encryption at rest**: KMS — ativado na criação
- **Encryption in transit**: SSL/TLS
- **VPC**: RDS deve ficar em subnet privada
- **IAM Authentication**: login via token IAM ao invés de senha
- **Secrets Manager**: rotação automática de credenciais

---

## Quando Usar RDS vs Aurora vs DynamoDB

| Situação | Recomendação |
|---------|-------------|
| MySQL/PostgreSQL existente | RDS |
| Alta performance + HA nativa | Aurora |
| Cargas imprevisíveis / dev | Aurora Serverless |
| Presença global + DR | Aurora Global |
| NoSQL key-value / escala massiva | DynamoDB |

---

## Pegadinhas da Prova

1. **Read Replicas não fazem failover automático** — Multi-AZ faz
2. **Aurora cria 2 réplicas por padrão** para HA — RDS Multi-AZ cria 1
3. **Snapshots funcionam mesmo com banco encriptado** (mesma chave KMS)
4. **Promover Read Replica** cria um banco independente (não replica mais)
5. **Aurora Serverless** não suporta read replicas públicas
6. **Multi-AZ para RDS Oracle** requer licença "bring your own license" para 2 nós

---

## Caso Real

### Amazon Prime Video
Usa Aurora Global Database para garantir que usuários em qualquer país experimentem o mesmo banco de dados com latência mínima, com failover automático em caso de falha regional.
