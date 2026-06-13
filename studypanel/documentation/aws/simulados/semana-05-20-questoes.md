# Simulado — Semana 5 (20 Questões)

**Semana:** 5 | **Nível:** Intermediário | **Cobertura:** Semanas 1 a 5

> Responda sem consultar o material. Anote suas respostas e confira no gabarito ao final.

---

## Questões

### 1. IAM
Uma empresa precisa que uma instância EC2 acesse objetos no S3 sem expor credenciais no código. Qual é a abordagem recomendada?

A) Criar um usuário IAM com Access Key e configurar na instância  
B) Criar uma IAM Role com política S3 e anexar à instância EC2  
C) Hardcodar as credenciais do usuário root no código  
D) Usar variáveis de ambiente com as credenciais do admin  

---

### 2. VPC
Uma empresa tem uma aplicação web de 3 camadas (web, app, banco). A camada web deve ser acessível pela internet, a camada app só pela web, e o banco só pela camada app. Qual configuração de VPC atende esse requisito?

A) Três subnets públicas, uma por camada  
B) Uma subnet pública (web) e duas privadas (app e banco)  
C) Três subnets privadas com NAT Gateway  
D) Uma subnet pública com Security Groups restritivos  

---

### 3. EC2 — Pricing
Uma empresa de processamento de imagens precisa executar jobs diariamente por 2 horas, e o processo pode ser interrompido e retomado. Qual opção de EC2 é mais econômica?

A) On-Demand  
B) Reserved Instances (1 ano)  
C) Spot Instances  
D) Dedicated Hosts  

---

### 4. Auto Scaling
Uma aplicação tem picos de uso entre 12h e 14h todos os dias. Qual política de Auto Scaling é mais adequada?

A) Target Tracking (CPU 70%)  
B) Step Scaling  
C) Scheduled Scaling  
D) Predictive Scaling  

---

### 5. S3
Uma empresa quer armazenar logs de auditoria que serão acessados apenas em caso de investigação (raramente). O custo deve ser mínimo. Qual classe S3 é mais adequada?

A) S3 Standard  
B) S3 Standard-IA  
C) S3 Glacier Deep Archive  
D) S3 One Zone-IA  

---

### 6. S3 — Segurança
Uma empresa quer garantir que objetos no S3 só possam ser acessados via CloudFront e nunca diretamente. Qual configuração resolve isso?

A) Bucket ACL privada + CloudFront Signed URLs  
B) CloudFront com OAC (Origin Access Control) + Bucket Policy restritiva  
C) S3 Block Public Access + Route 53  
D) IAM Policy no bucket + CloudFront  

---

### 7. RDS
Uma aplicação de e-commerce tem leituras 10x mais frequentes que escritas. A equipe quer reduzir carga no banco primário. Qual solução é mais adequada?

A) Fazer upgrade para instância maior  
B) Habilitar Multi-AZ  
C) Criar Read Replicas e direcionar consultas a elas  
D) Migrar para DynamoDB  

---

### 8. RDS vs Aurora
Uma startup precisa de banco relacional MySQL com alta disponibilidade, escalabilidade automática de storage e custos baixos. Qual escolha é mais adequada?

A) RDS MySQL com Multi-AZ  
B) Aurora MySQL com Aurora Serverless v2  
C) DynamoDB com partiql  
D) RDS MySQL Standard sem HA  

---

### 9. DynamoDB
Uma aplicação processa 1 milhão de pedidos por dia com picos imprevisíveis. Qual modo de capacidade do DynamoDB é mais recomendado?

A) Provisioned com valores fixos altos  
B) Provisioned com Auto Scaling  
C) On-Demand  
D) Reserved Capacity  

---

### 10. DynamoDB — Índices
Um sistema precisa consultar pedidos por `customer_id` (chave primária) mas também precisa listar todos os pedidos por `status`. Como modelar isso?

A) Criar duas tabelas separadas  
B) Fazer scan completo na tabela e filtrar no código  
C) Criar um GSI com `status` como partition key  
D) Adicionar um LSI com `status` como sort key  

---

### 11. ELB
Uma API REST precisa de SSL termination, routing por path (/api vs /admin) e suporte a WebSocket. Qual tipo de Load Balancer usar?

A) Classic Load Balancer (CLB)  
B) Network Load Balancer (NLB)  
C) Application Load Balancer (ALB)  
D) Gateway Load Balancer (GWLB)  

---

### 12. CloudFront
Um site global serve imagens que mudam raramente. A equipe quer reduzir latência para usuários no mundo todo. Qual solução?

A) Route 53 com Latency Routing direto para S3  
B) CloudFront na frente do S3 com TTL alto  
C) Multi-region S3 Replication  
D) EC2 em múltiplas regiões com Route 53  

---

### 13. IAM — Policy Evaluation
Um usuário IAM tem uma política que concede `s3:PutObject` e pertence a um grupo que tem uma política `Deny s3:PutObject` para o bucket `logs`. O que acontece quando o usuário tenta fazer upload no bucket `logs`?

A) Permitido — Allow prevalece sobre Deny do grupo  
B) Negado — Deny explícito sempre sobrepõe Allow  
C) Permitido — o grupo não tem prioridade sobre o usuário  
D) Depende da ordem de avaliação das políticas  

---

### 14. VPC — Conectividade
Uma empresa tem 3 VPCs: A, B e C. Existe peering entre A-B e B-C. A empresa quer que A acesse recursos em C. O que é necessário?

A) Nada — o tráfego passa por B automaticamente  
B) Criar peering direto entre A e C  
C) Usar Transit Gateway  
D) B ou C — ambas são válidas  

---

### 15. EC2 — Storage
Um banco de dados de alto desempenho em EC2 precisa de armazenamento com IOPS garantidos e baixa latência. Qual volume EBS é mais adequado?

A) gp2 (General Purpose SSD)  
B) gp3 (General Purpose SSD — nova geração)  
C) io2 Block Express (Provisioned IOPS SSD)  
D) st1 (Throughput Optimized HDD)  

---

### 16. S3 — Versionamento
Uma empresa habilita versionamento no S3. Um usuário deleta um arquivo. O que acontece?

A) O arquivo é permanentemente removido  
B) O arquivo é movido para Glacier  
C) Um marcador de deleção é criado; versões anteriores permanecem  
D) O arquivo fica inacessível por 30 dias antes de ser removido  

---

### 17. Alta Disponibilidade
Uma empresa quer garantir que sua aplicação em EC2 continue funcionando mesmo se uma AZ inteira falhar. Qual arquitetura atende isso?

A) EC2 com EBS Multi-Attach  
B) ASG com instâncias em múltiplas AZs + ALB  
C) EC2 com snapshot automático  
D) EC2 com Elastic IP  

---

### 18. Segurança
Uma empresa precisa garantir que todas as chamadas de API na conta AWS sejam registradas para auditoria. Qual serviço usar?

A) CloudWatch Metrics  
B) VPC Flow Logs  
C) AWS CloudTrail  
D) AWS Config  

---

### 19. RDS — Backups
Uma empresa precisa restaurar o banco de dados para um ponto exato de 3 dias atrás (às 14h32). Qual funcionalidade do RDS permite isso?

A) Manual Snapshot  
B) Point-in-Time Recovery (PITR)  
C) Aurora Backtrack  
D) Read Replica promotion  

---

### 20. Cost Optimization
Uma empresa tem instâncias EC2 executando 24/7 há 2 anos e planeja continuar por mais 1 ano. Qual opção reduz mais o custo sem mudar a arquitetura?

A) Spot Instances  
B) On-Demand com downgrade de instância  
C) Reserved Instances (1 ano, all upfront)  
D) Savings Plans  

---

## Gabarito

| # | Resposta | Justificativa |
|---|---------|--------------|
| 1 | B | IAM Role é a forma segura e recomendada para EC2 acessar S3 |
| 2 | B | Arquitetura de 3 camadas clássica com subnet pública e privadas |
| 3 | C | Spot é até 90% mais barato para cargas interruptíveis |
| 4 | C | Picos previsíveis por horário = Scheduled Scaling |
| 5 | C | Glacier Deep Archive = menor custo para acesso rarissimo |
| 6 | B | OAC + Bucket Policy restritiva é o padrão correto |
| 7 | C | Read Replicas escalam leitura; Multi-AZ é para HA |
| 8 | B | Aurora Serverless v2 = auto-scaling, HA nativa, MySQL compatível |
| 9 | C | On-Demand para cargas imprevisíveis sem planejamento prévio |
| 10 | C | GSI permite acesso por atributo diferente da PK |
| 11 | C | ALB = L7, routing por path, SSL termination, WebSocket |
| 12 | B | CloudFront + TTL alto = cache global, menor latência |
| 13 | B | Deny explícito sempre sobrepõe Allow no IAM |
| 14 | D | Peering direto A-C OU Transit Gateway; ambas resolvem |
| 15 | C | io2 Block Express = IOPS garantidos para banco de dados |
| 16 | C | Versionamento cria delete marker; arquivos não são removidos |
| 17 | B | ASG multi-AZ + ALB = alta disponibilidade padrão |
| 18 | C | CloudTrail registra todas as chamadas de API |
| 19 | B | PITR permite restaurar para qualquer segundo no período de retenção |
| 20 | C | Reserved 1 ano all upfront = maior desconto para uso constante |

---

## Análise de Desempenho

Calcule seu percentual: **(acertos / 20) × 100**

| Resultado | Ação |
|-----------|------|
| ≥ 85% | Excelente! Avance |
| 70-84% | Revise os erros e prossiga |
| 50-69% | Revise os serviços com erros antes de avançar |
| < 50% | Revise todo o material das semanas 1-5 |
