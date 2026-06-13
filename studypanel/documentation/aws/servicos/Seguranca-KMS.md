# Segurança e KMS — Criptografia e Proteção

**Incidência:** 🔴 Muito Alta | **Domínio:** Segurança

---

## KMS — Key Management Service

### Conceito
KMS gerencia chaves criptográficas para proteger dados na AWS.

### Tipos de Chaves

| Tipo | Gerenciamento | Rotação | Custo |
|------|--------------|---------|-------|
| **AWS Managed Key** | AWS | Automática (1 ano) | Gratuita |
| **Customer Managed Key (CMK)** | Você | Opcional (manual ou auto) | $1/mês |
| **Customer Provided Key (SSE-C)** | Você fornece | Você controla | — |

### Criptografia no S3

| Opção | Chave gerenciada por |
|-------|---------------------|
| **SSE-S3** | AWS (chave da AWS) |
| **SSE-KMS** | KMS (sua CMK) — audit trail no CloudTrail |
| **SSE-C** | Você fornece a chave em cada request |
| **Client-Side Encryption** | Você criptografa antes de enviar |

### Envelope Encryption
KMS usa criptografia de envelope:
1. Gera **Data Encryption Key (DEK)** para criptografar os dados
2. DEK é criptografada com a CMK (armazenada no KMS)
3. DEK criptografada é armazenada junto com os dados
4. Para descriptografar: KMS descriptografa a DEK → você usa a DEK nos dados

---

## Secrets Manager

- Armazena e rotaciona credenciais automaticamente
- Suporta RDS, Redshift, DocumentDB nativamente
- Rotação via Lambda (serverless)
- Integração com SDK AWS para recuperar secrets em runtime
- Custo: ~$0.40/secret/mês

### vs SSM Parameter Store
| Critério | Secrets Manager | SSM Parameter Store |
|---------|----------------|---------------------|
| Rotação automática | Sim (nativo) | Manual (via Lambda) |
| Custo | ~$0.40/secret | Gratuito (Standard) |
| Valor máximo | 64KB | 8KB (Standard) |
| Referência em ECS/EC2 | Sim | Sim |
| Para uso em prova | Credenciais com rotação | Configurações gerais |

---

## ACM — AWS Certificate Manager

- Provisionamento de certificados SSL/TLS
- Gratuito para certificados em serviços AWS (ALB, CloudFront, API Gateway)
- Renovação automática
- Não pode exportar chave privada (use ACM Private CA para isso)

---

## WAF — Web Application Firewall

- Protege contra ataques web (SQL Injection, XSS, etc.)
- Anexado a: ALB, API Gateway, CloudFront, AppSync
- Usa **Web ACLs** com regras customizáveis
- Regras gerenciadas pela AWS ou AWS Marketplace

---

## Shield — Proteção DDoS

| Tier | Proteção | Custo |
|------|---------|-------|
| **Shield Standard** | DDoS básico (camadas 3 e 4) | Gratuito |
| **Shield Advanced** | DDoS avançado + suporte 24/7 + proteção de custo | $3.000/mês |

Shield Standard é ativado automaticamente para todos os clientes.

---

## GuardDuty — Detecção de Ameaças

- Analisa: CloudTrail, VPC Flow Logs, DNS logs
- Machine learning para identificar comportamentos anômalos
- Exemplos de threats: acesso de IP suspeito, crypto mining, exfiltração de dados
- Ação: gera Finding → EventBridge → Lambda para remediar

---

## Inspector

- Avalia vulnerabilidades em instâncias EC2 e containers
- CVEs, exposição de rede, desvios de segurança
- Integrado com ECR para scan de imagens

---

## CloudTrail

- Auditoria de todas as chamadas de API na conta AWS
- Quem fez o quê, quando e de onde
- Logs armazenados no S3
- **Log File Integrity**: hash para garantir que logs não foram alterados
- Ativado por padrão (90 dias); para retenção maior, envie ao S3

---

## Compliance e Auditoria

| Serviço | Função |
|---------|--------|
| **CloudTrail** | Auditoria de chamadas API |
| **Config** | Conformidade de recursos ao longo do tempo |
| **Security Hub** | Visão centralizada de postura de segurança |
| **Macie** | Detecta dados sensíveis (PII) no S3 |
| **Detective** | Investigação de incidentes de segurança |

---

## Pegadinhas da Prova

1. **KMS keys são regionais** — para cross-region, precisa copiar a chave ou usar multi-region key
2. **SSE-KMS gera audit trail** no CloudTrail — SSE-S3 não
3. **Secrets Manager** é preferível ao Parameter Store para credenciais com rotação
4. **WAF não protege contra DDoS de camada 3/4** — Shield faz isso
5. **CloudTrail não monitora tráfego de rede** — VPC Flow Logs fazem isso
6. **GuardDuty deve ser ativado** — não vem ativo por padrão
7. **ACM gratuito para ALB/CloudFront** — pago para exportar (via Private CA)
