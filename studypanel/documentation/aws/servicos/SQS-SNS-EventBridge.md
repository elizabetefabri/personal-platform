# SQS, SNS e EventBridge — Mensageria e Eventos

**Incidência:** 🟠 Alta | **Domínio:** Integração / Serverless

---

## SQS — Simple Queue Service

Fila gerenciada para desacoplamento de sistemas.

### Standard vs FIFO

| Característica | Standard | FIFO |
|---------------|---------|------|
| Entrega | Ao menos 1 vez (pode duplicar) | Exatamente 1 vez |
| Ordem | Best-effort | Garantida |
| Throughput | Ilimitado | 3.000/s (sem batching) |
| Caso de uso | Processamento distribuído | Transações financeiras |

### Parâmetros Importantes

| Parâmetro | Descrição |
|-----------|-----------|
| **Visibility Timeout** | Tempo que msg fica invisível após ser lida (default 30s) |
| **Message Retention** | 4 dias por padrão (máx 14 dias) |
| **Delay Queue** | Atraso antes de mensagem aparecer na fila |
| **Long Polling** | Reduz chamadas vazias (recomendado) |
| **Dead Letter Queue (DLQ)** | Fila para mensagens que falharam N vezes |

### SQS + Lambda
- Lambda faz polling automático na fila (Event Source Mapping)
- Processa em batches (até 10.000 por batch)
- DLQ captura mensagens que falharam

---

## SNS — Simple Notification Service

Pub/Sub: um publisher, múltiplos subscribers.

### Funcionamento
```
Publisher
    │
    ▼
SNS Topic
    │
    ├─→ SQS Queue (Fan-out)
    ├─→ Lambda
    ├─→ HTTP/HTTPS endpoint
    ├─→ Email
    └─→ SMS
```

### Fan-out Pattern
- SNS → múltiplos SQS
- Processa o mesmo evento em paralelo por diferentes consumidores
- Muito comum na prova: **SNS + SQS = Fan-out**

### Filtros de Mensagem
- SNS pode filtrar mensagens por atributos
- Cada subscriber recebe apenas as mensagens relevantes para ele

---

## EventBridge

EventBridge é um bus de eventos para integrar serviços AWS e externos.

### Componentes
| Componente | Descrição |
|-----------|-----------|
| **Event Bus** | Canal onde eventos fluem |
| **Rule** | Filtra eventos por padrão |
| **Target** | Destino: Lambda, SQS, Step Functions, etc. |
| **Schema Registry** | Registra estrutura dos eventos |

### Tipos de Event Bus
- **Default Bus**: eventos de serviços AWS
- **Custom Bus**: eventos da sua aplicação
- **Partner Bus**: eventos de SaaS (Datadog, Zendesk, etc.)

### Scheduled Events (Cron)
- Substituiu CloudWatch Events
- Executa regras em intervalos: `rate(5 minutes)` ou cron expression
- Target mais comum: Lambda

---

## Comparação: SQS vs SNS vs EventBridge

| Critério | SQS | SNS | EventBridge |
|---------|-----|-----|-------------|
| Padrão | Pull (fila) | Push (pub/sub) | Push (event bus) |
| Consumidores | 1 consumidor por msg | Múltiplos | Múltiplos |
| Filtros | Não (fila FIFO tem grupos) | Por atributo | Por padrão JSON |
| Integrações | AWS | AWS + HTTP/Email/SMS | AWS + SaaS + Custom |
| Persistência | Até 14 dias | Não persiste | Não persiste |
| Melhor para | Desacoplamento de workers | Broadcast de notificações | Event-driven architecture |

---

## Padrões de Arquitetura

### Desacoplamento com SQS
```
Frontend → API Gateway → Lambda → SQS → Lambda (processamento)
```

### Fan-out com SNS + SQS
```
Evento → SNS Topic
             │
             ├─→ SQS 1 → Lambda (email)
             ├─→ SQS 2 → Lambda (analytics)
             └─→ SQS 3 → Lambda (notificação push)
```

### Event-driven com EventBridge
```
S3 Upload → EventBridge Rule → Lambda (redimensionar imagem)
                             → Step Functions (workflow)
                             → SQS (processar depois)
```

---

## Pegadinhas da Prova

1. **SQS Standard pode entregar mensagens fora de ordem** — FIFO garante ordem
2. **SQS não é push** — é pull (consumidor faz polling)
3. **SNS não persiste mensagens** — se subscriber não está disponível, perde
4. **DLQ deve ser configurada** — não é automática
5. **Visibility Timeout** deve ser maior que o tempo de processamento
6. **EventBridge é o sucessor do CloudWatch Events** — mesma funcionalidade, mais poderoso
7. **SNS + SQS Fan-out** é o padrão para processar o mesmo evento de múltiplas formas

---

## Caso Real

### iFood
Usa SNS + SQS para processar pedidos. Quando um pedido é feito, SNS notifica múltiplas filas SQS: uma para o restaurante, outra para o entregador, outra para analytics — tudo em paralelo.
