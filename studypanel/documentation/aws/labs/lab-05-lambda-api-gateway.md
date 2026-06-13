# Lab 05 — Lambda + API Gateway

**Dia:** 38 (Semana 6) | **Duração:** 50 minutos

---

## Objetivo

Criar uma API REST com API Gateway + Lambda que retorna dados JSON.

---

## Arquitetura

```
Cliente HTTP
     │
     ▼
API Gateway (REST API)
     │
     ▼
Lambda Function (Python)
     │
     ▼
DynamoDB (opcional)
```

---

## Passo a Passo

### 1. Criar Lambda Function

1. Lambda → Create function
2. Author from scratch
3. Name: `studylab-hello-api`
4. Runtime: Python 3.12
5. Architecture: x86_64
6. Role: Create new role (basic permissions)
7. Create function

**Código (Editor inline):**
```python
import json
from datetime import datetime

def lambda_handler(event, context):
    method = event.get('httpMethod', 'GET')
    path = event.get('path', '/')
    
    body = {
        "message": "StudyPanel API - AWS SAA-C03 Lab",
        "path": path,
        "method": method,
        "timestamp": datetime.utcnow().isoformat(),
        "studente": "Elizabete",
        "certification": "AWS SAA-C03"
    }
    
    return {
        "statusCode": 200,
        "headers": {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*"
        },
        "body": json.dumps(body)
    }
```

8. Deploy

---

### 2. Testar Lambda Diretamente

1. Test → Configure test event
2. Event JSON:
```json
{
  "httpMethod": "GET",
  "path": "/hello",
  "queryStringParameters": null
}
```
3. Test → Verificar Response com status 200

---

### 3. Criar API Gateway

1. API Gateway → Create API → REST API → Build
2. Name: `studylab-api`
3. Endpoint Type: Regional
4. Create API

**Criar Resource:**
1. Resources → Actions → Create Resource
2. Resource Path: `hello`

**Criar Method GET:**
1. Selecione `/hello` → Actions → Create Method → GET
2. Integration type: Lambda Function
3. Lambda Function: `studylab-hello-api`
4. Save → OK (dar permissão para API GW invocar Lambda)

---

### 4. Fazer Deploy da API

1. Actions → Deploy API
2. Stage: New Stage → Name: `dev`
3. Deploy

Anote a URL: `https://xxx.execute-api.sa-east-1.amazonaws.com/dev`

---

### 5. Testar a API

```bash
# Via curl
curl https://xxx.execute-api.sa-east-1.amazonaws.com/dev/hello

# Resultado esperado:
{
  "message": "StudyPanel API - AWS SAA-C03 Lab",
  "path": "/hello",
  "method": "GET",
  "timestamp": "2024-01-15T10:30:00",
  "studente": "Elizabete",
  "certification": "AWS SAA-C03"
}
```

---

### 6. Adicionar Variável de Ambiente na Lambda

1. Lambda → `studylab-hello-api` → Configuration → Environment variables
2. Add: `ENVIRONMENT = production`
3. No código, acesse: `import os; env = os.environ.get('ENVIRONMENT')`

---

### 7. Configurar DLQ (Dead Letter Queue)

1. Lambda → Configuration → Asynchronous invocation
2. Dead-letter queue: Criar SQS Standard queue `studylab-dlq`
3. Selecionar a queue criada
4. Máximo de retentativas: 2

---

## Resultado Esperado

- Lambda funcionando e retornando JSON
- API Gateway configurado com endpoint público
- API respondendo via curl/Postman

---

## Validação

1. `curl` retorna status 200 com JSON correto
2. CloudWatch Logs mostra execuções da Lambda
3. Teste com método não suportado retorna 403 ou 405

---

## Limpeza

1. Delete API Gateway
2. Delete Lambda function
3. Delete CloudWatch Log Groups

---

## Perguntas de Revisão

1. Qual o timeout máximo de uma Lambda?
2. O que acontece com o cold start quando a Lambda fica sem invocações?
3. Como reduzir cold start em produção?
4. Qual a diferença entre invocação síncrona e assíncrona no Lambda?
