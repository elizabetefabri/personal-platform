# Matriz de Probabilidade — AWS SAA-C03

Baseada na frequência histórica de cobrança na certificação real.

---

## 🔴 Muito Alta Incidência (~50% do tempo)

| Serviço / Tema | Por que é importante |
|----------------|---------------------|
| **IAM** | Toda arquitetura segura passa por permissões |
| **VPC** | Networking é base de tudo na AWS |
| **EC2** | Serviço mais fundamental da AWS |
| **Auto Scaling + ELB** | Sempre testado em cenários de resiliência |
| **S3** | Storage universal, muitas pegadinhas |
| **RDS + Aurora** | Banco relacional gerenciado mais cobrado |
| **DynamoDB** | NoSQL com muitos cenários de uso |
| **Lambda** | Serverless é domínio cobrado extensivamente |
| **Segurança / IAM Roles** | ~30% da prova envolve segurança |
| **Alta Disponibilidade** | Padrões multi-AZ, multi-region |
| **Disaster Recovery** | RTO/RPO — muito cobrado |
| **Well-Architected Framework** | Pilar de excelência operacional |
| **Cost Optimization** | Sempre presente em questões de trade-off |

---

## 🟠 Alta Incidência (~30% do tempo)

| Serviço / Tema | Por que é importante |
|----------------|---------------------|
| **Route 53** | DNS e routing policies são frequentes |
| **CloudFront** | CDN + S3 + segurança = cenários comuns |
| **SQS** | Desacoplamento de aplicações |
| **SNS** | Fan-out patterns |
| **EventBridge** | Event-driven architecture |
| **ECS / EKS / Fargate** | Containers estão crescendo na prova |
| **API Gateway** | Par inseparável do Lambda |
| **CloudWatch** | Monitoramento e alarmes |
| **KMS** | Criptografia de dados |
| **Secrets Manager** | Gerenciamento de credenciais |

---

## 🟡 Média Incidência (~15% do tempo)

| Serviço / Tema | Por que é importante |
|----------------|---------------------|
| **EFS** | Storage compartilhado — questões específicas |
| **FSx** | Windows workloads |
| **Storage Gateway** | Híbrido on-premises + AWS |
| **CloudFormation** | IaC na AWS |
| **Systems Manager** | Gerenciamento de instâncias |
| **Direct Connect** | Conectividade dedicada |
| **VPN Site-to-Site** | Conexão híbrida |
| **ElastiCache** | Cache — Redis vs Memcached |
| **Step Functions** | Orquestração de workflows |

---

## 🟢 Baixa Incidência (~5% do tempo)

| Serviço / Tema | Observação |
|----------------|-----------|
| **Rekognition, Textract** | IA/ML — raramente cobrado |
| **Kinesis Data Firehose** | Big Data — específico |
| **AWS Glue** | ETL — específico |
| **Redshift** | Data Warehouse — específico |
| **AppSync** | GraphQL — raro |
| **Cognito** | Auth pode aparecer em serverless |

---

## Mapa de Decisão Rápida

### Armazenamento

| Necessidade | Serviço |
|------------|---------|
| Arquivos estáticos / objetos | S3 |
| Disco de instância EC2 | EBS |
| Storage compartilhado entre instâncias | EFS |
| Windows + SMB | FSx for Windows |
| Backup on-premises para nuvem | Storage Gateway |

### Banco de Dados

| Necessidade | Serviço |
|------------|---------|
| SQL relacional gerenciado | RDS (MySQL, Postgres, etc.) |
| SQL de alta performance | Aurora |
| NoSQL key-value / document | DynamoDB |
| Cache em memória | ElastiCache (Redis ou Memcached) |
| Data Warehouse | Redshift |

### Serverless

| Necessidade | Serviço |
|------------|---------|
| Processamento de eventos | Lambda |
| API HTTP | API Gateway |
| Fila de mensagens | SQS |
| Pub/Sub | SNS |
| Event-driven | EventBridge |
| Orquestração de steps | Step Functions |

### Containers

| Necessidade | Serviço |
|------------|---------|
| Containers gerenciados simples | ECS |
| Kubernetes gerenciado | EKS |
| Serverless containers | Fargate |

### Alta Disponibilidade

| Necessidade | Solução |
|------------|---------|
| Instâncias redundantes | Multi-AZ |
| Escalonamento automático | Auto Scaling Group |
| Distribuição de carga | ALB / NLB |
| Recuperação global | Multi-Region + Route 53 |

---

## Pegadinhas Mais Comuns

1. **NACLs são stateless** — Security Groups são stateful
2. **S3 não é em uma única AZ** — é regional com durabilidade 99.999999999%
3. **Read Replicas não são para HA** — são para leitura; Multi-AZ é para HA
4. **Lambda tem timeout máximo de 15 minutos**
5. **Spot Instances podem ser terminadas com 2 minutos de aviso**
6. **Reserved Instances reduzem custo mas não garantem capacidade** — use Capacity Reservations
7. **CloudFront serve conteúdo estático e dinâmico** — não é só para S3
8. **DynamoDB On-Demand vs Provisioned** — saber quando usar cada um
9. **SQS Standard pode duplicar mensagens** — FIFO garante exatamente uma entrega
10. **IAM Roles são preferíveis a Access Keys** para aplicações rodando na AWS
