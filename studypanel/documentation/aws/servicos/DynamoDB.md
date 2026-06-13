# DynamoDB — Banco de Dados NoSQL

**Incidência:** 🔴 Muito Alta | **Domínio:** Banco de Dados

---

## Conceito

DynamoDB é um banco de dados NoSQL serverless, totalmente gerenciado pela AWS, com latência de milissegundos em qualquer escala.

---

## Estrutura de Dados

| Conceito | Descrição |
|---------|-----------|
| **Tabela** | Container de itens (como uma coleção) |
| **Item** | Registro (como um documento JSON) |
| **Atributo** | Campo do item |
| **Partition Key (PK)** | Chave de partição — obrigatória, define a distribuição |
| **Sort Key (SK)** | Chave de ordenação — opcional, cria chave composta |

---

## Modos de Capacidade

### On-Demand
- Paga por operação de leitura/escrita
- Escala automaticamente sem planejamento
- Ideal para cargas imprevisíveis ou novos projetos

### Provisioned
- Define RCU (Read Capacity Units) e WCU (Write Capacity Units)
- Mais barato para cargas previsíveis
- Pode usar **Auto Scaling** para ajustar automaticamente

### Calculando Capacity Units:
- **1 RCU** = 1 leitura forte consistente de até 4KB por segundo
- **1 WCU** = 1 escrita de até 1KB por segundo

---

## Índices

### GSI (Global Secondary Index)
- Partition key e sort key **diferentes** da tabela original
- Pode ser criado após a tabela
- Tem sua própria capacidade de leitura/escrita
- Suporta apenas eventual consistency

### LSI (Local Secondary Index)
- Mesma partition key, **sort key diferente**
- Deve ser criado junto com a tabela
- Compartilha capacidade com a tabela
- Suporta leitura fortemente consistente

---

## Tipos de Consistência

| Tipo | Descrição | Custo |
|------|-----------|-------|
| **Eventual Consistency** | Dados propagados em <1s | 0.5 RCU |
| **Strong Consistency** | Dados mais atuais garantidos | 1 RCU |

**Padrão do DynamoDB:** Eventually consistent.

---

## DynamoDB Streams

- Captura mudanças (INSERT, MODIFY, DELETE) em ordem
- Retém eventos por 24 horas
- Pode disparar Lambda para processamento em tempo real
- Use case: auditoria, replicação, notificações em tempo real

---

## DynamoDB Accelerator (DAX)

- Cache em memória nativo do DynamoDB
- Latência de microssegundos (vs milissegundos)
- Totalmente gerenciado, compatível com APIs do DynamoDB
- Ideal quando leituras são muito frequentes

---

## TTL (Time to Live)

- Expiração automática de itens baseada em timestamp
- Gratuito
- Ideal para sessões, tokens, dados temporários

---

## DynamoDB Global Tables

- Replicação multi-region ativa-ativa
- Escrita em qualquer região
- Latência global baixa
- Resolução de conflitos: "last writer wins"

---

## Padrões de Acesso

### Single Table Design
DynamoDB funciona melhor com **tudo em uma tabela**:
- Diferentes tipos de entidades na mesma tabela
- PK e SK representam hierarquias
- GSIs para diferentes padrões de acesso

Exemplo:
```
PK              SK              Entidade
USER#123        PROFILE         Perfil do usuário
USER#123        ORDER#456       Pedido do usuário
PRODUCT#789     DETAILS         Detalhes do produto
```

---

## Pegadinhas da Prova

1. **DynamoDB não suporta SQL** — use PartiQL se precisar de consultas SQL-like
2. **GSI não garante strong consistency** — sempre eventual
3. **Chave composta = PK + SK** — permite múltiplos itens com mesmo PK
4. **Hot partition**: evite PK muito concentrada (ex: data atual como PK)
5. **Scans são caros** — use Query sempre que possível
6. **DAX não é adequado para escritas intensas** — é para leitura
7. **LSI deve ser criado na criação da tabela** — não pode adicionar depois

---

## RDS vs DynamoDB

| Critério | RDS | DynamoDB |
|---------|-----|---------|
| Modelo | Relacional (SQL) | NoSQL (key-value/document) |
| Esquema | Fixo | Flexível |
| Escala | Vertical + Read Replicas | Horizontal automática |
| Performance | Milissegundos | Milissegundos (DAX: microssegundos) |
| Transações | Completas (ACID) | Transações DynamoDB (limitadas) |
| Joins | Sim | Não (design diferente) |

---

## Caso Real

### Amazon
O próprio site da Amazon usa DynamoDB para carrinho de compras, sessões e recomendações — processando milhões de operações por segundo com latência consistente.
