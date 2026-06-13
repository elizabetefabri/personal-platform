# Lambda — Serverless Functions

**Incidência:** 🔴 Muito Alta | **Domínio:** Serverless / Compute

---

## Conceito

Lambda executa código sem gerenciar servidores. Você paga apenas pelo tempo de execução.

Analogia: é como pagar por corrida de táxi — não compra o carro, não paga quando não usa.

---

## Características Principais

| Característica | Valor |
|---------------|-------|
| Timeout máximo | **15 minutos** |
| Memória | 128MB a 10GB |
| CPU | Proporcional à memória |
| Tamanho do pacote (zip) | 50MB (250MB descomprimido) |
| Tamanho com container | Até 10GB |
| Invocações simultâneas | 1.000 por conta (ajustável) |
| Ephemeral storage (/tmp) | 512MB a 10GB |

---

## Triggers (Gatilhos)

Lambda pode ser invocado por:
- **API Gateway** → Requisições HTTP
- **S3** → Upload/deleção de objetos
- **DynamoDB Streams** → Mudanças no banco
- **SQS** → Processamento de mensagens
- **SNS** → Notificações
- **EventBridge** → Eventos agendados ou de outros serviços
- **ALB** → HTTP requests
- **CloudFront** → Lambda@Edge

---

## Modelo de Invocação

| Tipo | Comportamento | Exemplo |
|------|--------------|---------|
| **Síncrono** | Aguarda resposta | API Gateway, ALB |
| **Assíncrono** | Dispara e esquece | S3, SNS |
| **Event Source Mapping** | Polling | SQS, DynamoDB Streams, Kinesis |

---

## Cold Start vs Warm Start

**Cold Start**: Lambda precisa inicializar o container (~100ms a segundos)
**Warm Start**: Reutiliza container já existente (milissegundos)

### Como mitigar Cold Start:
- **Provisioned Concurrency**: mantém instâncias aquecidas (custo extra)
- **Lambda SnapStart** (Java): snapshot de estado inicializado
- Usar runtimes mais leves (Python, Node.js)
- Reduzir o tamanho do pacote de deploy

---

## Lambda + VPC

Lambda pode rodar dentro de uma VPC para acessar recursos privados:
- RDS, ElastiCache, serviços internos
- Adiciona ~1-3s de cold start (criação de ENI)
- Precisa de subnets com NAT Gateway para acessar a internet

---

## Layers

- Camadas compartilhadas entre funções
- Ideal para: dependências, utilitários, configurações
- Até 5 layers por função
- Reduz tamanho do deploy e facilita compartilhamento

---

## Concorrência

| Tipo | Descrição |
|------|-----------|
| **Concorrência Reservada** | Limita a função a N execuções simultâneas |
| **Concorrência Provisionada** | Mantém N instâncias aquecidas |
| **Account Concurrency Limit** | 1.000 por padrão (pode ser aumentado) |

---

## Arquitetura Serverless Típica

```
Cliente
  │
  ▼
API Gateway
  │
  ▼
Lambda
  │
  ├─→ DynamoDB (persistência)
  ├─→ SQS (processamento assíncrono)
  ├─→ S3 (arquivos)
  └─→ SNS (notificações)
```

---

## Lambda@Edge

- Executa Lambda próximo ao usuário (CloudFront Edge Locations)
- Usada para: A/B testing, redirecionamentos, autenticação na borda
- Timeout menor: 5s (viewer) ou 30s (origin)

---

## Pegadinhas da Prova

1. **Timeout máximo é 15 minutos** — use Step Functions ou ECS para processos longos
2. **Lambda não mantém estado** entre invocações (use /tmp com cuidado)
3. **Dead Letter Queue** — configure para processar erros em invocações assíncronas
4. **SQS como trigger** faz polling automático (Event Source Mapping)
5. **Lambda dentro de VPC** requer subnets e SG — sem internet sem NAT
6. **Versões e aliases** — versão imutável; alias aponta para versão
7. **Concorrência reservada = 0** desabilita a função completamente

---

## Caso Real

### Airbnb
Usa Lambda para processar imagens de propriedades. Quando um host faz upload, S3 dispara Lambda que redimensiona, converte formatos e aplica filtros automaticamente — sem servidor gerenciado.

---

## Lab Relacionado

[Lab 05 — Lambda + API Gateway](../labs/lab-05-lambda-api-gateway.md)
