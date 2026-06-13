# ECS, EKS e Fargate — Containers na AWS

**Incidência:** 🟠 Alta | **Domínio:** Containers / Compute

---

## Visão Geral

| Serviço | Descrição |
|---------|-----------|
| **ECS** | Orquestrador de containers da própria AWS |
| **EKS** | Kubernetes gerenciado pela AWS |
| **Fargate** | Modo serverless para ECS e EKS (sem gerenciar nós) |
| **ECR** | Registry de imagens Docker da AWS |

---

## ECS — Elastic Container Service

### Componentes
- **Task Definition**: Blueprint do container (imagem, CPU, memória, portas, IAM role)
- **Task**: Instância em execução de uma Task Definition
- **Service**: Garante que N tasks estejam sempre rodando
- **Cluster**: Grupo lógico de tasks e serviços

### Launch Types

| Tipo | Gerencia Infraestrutura | Custo |
|------|------------------------|-------|
| **EC2** | Você gerencia as instâncias | Menor |
| **Fargate** | AWS gerencia (serverless) | Maior |

### IAM no ECS
- **Task Role**: permissões que o container tem (ex: acessar S3)
- **Execution Role**: permissões para o agente ECS (ex: baixar imagem do ECR)

---

## Fargate

Fargate elimina a necessidade de gerenciar servidores EC2.

**Vantagens:**
- Sem gerenciar instâncias, patches ou scaling de nós
- Paga apenas por CPU e memória usados pela task
- Integra com ECS e EKS

**Quando usar:**
- Não quer gerenciar servidores (Platform Engineer)
- Cargas variáveis
- Workloads isolados por segurança

**Quando NÃO usar:**
- Necessidade de controle de hardware específico
- GPU (não suportado)
- Cargas muito intensas onde EC2 é mais barato

---

## EKS — Elastic Kubernetes Service

EKS é o Kubernetes gerenciado pela AWS.

### Componentes
- **Control Plane**: gerenciado pela AWS (master nodes, etcd, scheduler)
- **Worker Nodes**: EC2 ou Fargate onde as pods rodam
- **Node Groups**: grupos de EC2 gerenciados
- **Add-ons**: CoreDNS, kube-proxy, VPC CNI, etc.

### EKS vs ECS

| Critério | ECS | EKS |
|---------|-----|-----|
| Portabilidade | Somente AWS | Kubernetes padrão |
| Complexidade | Menor | Maior |
| Ecossistema | AWS-native | Kubernetes padrão |
| Curva de aprendizado | Menor | Maior |
| Multi-cloud | Não | Sim |

**Regra de prova:** Se a empresa "quer migrar para outro cloud" ou "usa Kubernetes hoje" → **EKS**. Se quer simplicidade na AWS → **ECS**.

---

## ECR — Elastic Container Registry

- Registry privado de imagens Docker
- Integrado com IAM para acesso
- Scan de vulnerabilidades automático
- Replicação entre regiões
- Lifecycle policies para limpeza de imagens antigas

---

## Padrões de Arquitetura

### ECS com ALB
```
Internet
    │
    ▼
ALB (Application Load Balancer)
    │
    ▼
ECS Service (Target Group)
    │
    ├── Task 1 (Container)
    ├── Task 2 (Container)
    └── Task 3 (Container)
```

### Service Discovery
- ECS integra com AWS Cloud Map
- Containers se comunicam via DNS interno
- Ex: `api.local` resolve para tasks do serviço `api`

---

## Pegadinhas da Prova

1. **Fargate não suporta GPU** — use EC2 launch type
2. **ECS Task Role ≠ Execution Role** — propósitos diferentes
3. **EKS control plane** tem custo fixo por hora (~$72/mês)
4. **Fargate** cobra por task iniciada, não por instância reservada
5. **ECS com EC2** é mais barato para cargas constantes e previsíveis
6. **Multi-container task**: containers compartilham network e volumes
7. **ECS Anywhere**: rodar tasks ECS em servidores on-premises

---

## Caso Real

### Nubank
Usa ECS com Fargate para seus microsserviços. Cada serviço é um container independente, escalando automaticamente sem a equipe de infra gerenciar servidores EC2.
