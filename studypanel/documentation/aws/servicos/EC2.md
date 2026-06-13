# EC2 — Elastic Compute Cloud

**Incidência:** 🔴 Muito Alta | **Domínio:** Compute

---

## Conceito

EC2 fornece capacidade computacional escalável na nuvem. É o equivalente a um servidor virtual que você controla.

---

## Tipos de Instâncias

| Família | Uso | Exemplos |
|---------|-----|---------|
| General Purpose | Balanceado | t3, t4g, m6i |
| Compute Optimized | CPU intensivo | c6i, c7g |
| Memory Optimized | RAM intensivo | r6i, x2idn |
| Storage Optimized | I/O intensivo | i3, d3 |
| Accelerated | GPU/FPGA | p4d, inf2 |

---

## Modelos de Preço

| Modelo | Quando Usar | Desconto |
|--------|-------------|---------|
| **On-Demand** | Cargas variáveis, sem compromisso | — |
| **Reserved (1-3 anos)** | Cargas previsíveis, longo prazo | Até 75% |
| **Savings Plans** | Flexível, por uso em $/hora | Até 66% |
| **Spot** | Cargas interruptíveis (batch, CI) | Até 90% |
| **Dedicated Host** | Compliance, licenças por socket | Alto custo |
| **Dedicated Instance** | Isolamento físico | Mais caro |

**Pergunta de prova:** Qual o mais barato para carga interruptível? → **Spot**
**Pergunta de prova:** Qual para banco de dados 24/7 por 3 anos? → **Reserved**

---

## Storage para EC2

| Tipo | Características |
|------|----------------|
| **EBS (gp3, io2)** | Persistente, rede, pode ser desanexado |
| **Instance Store** | Temporário, NVMe local, velocidade máxima |
| **EFS** | Compartilhado entre instâncias, NFS |

**Atenção:** Instance Store **perde dados** quando a instância para.

---

## Auto Scaling Group (ASG)

Define o número mínimo, desejado e máximo de instâncias.

### Políticas de Scaling
| Política | Gatilho |
|---------|---------|
| **Target Tracking** | Mantém métrica no alvo (ex: CPU 50%) |
| **Step Scaling** | Escala em passos por threshold |
| **Scheduled** | Hora pré-definida |
| **Predictive** | ML prevê e escala antecipadamente |

### Launch Template vs Launch Configuration
- **Launch Template** → use sempre (suporta múltiplas versões)
- **Launch Configuration** → legado, deprecado

---

## Elastic Load Balancer (ELB)

| Tipo | Camada | Uso |
|------|--------|-----|
| **ALB** | L7 (HTTP/HTTPS) | Routing por path/host, WebSocket |
| **NLB** | L4 (TCP/UDP) | Ultra baixa latência, IP fixo |
| **CLB** | L4+L7 | Legado, evitar |
| **GWLB** | L3 | Tráfego de appliances de rede |

**Dica de prova:** ALB para HTTP; NLB para TCP/UDP de alta performance.

---

## Placement Groups

| Tipo | Estratégia | Uso |
|------|-----------|-----|
| **Cluster** | Mesmo rack, mesma AZ | HPC, baixa latência |
| **Spread** | Hardware diferente | Crítico, alta disponibilidade |
| **Partition** | Partições lógicas | Hadoop, Kafka, Cassandra |

---

## AMI (Amazon Machine Image)

- Template com OS, aplicações e configurações
- Pode ser regional (copiar entre regiões manualmente)
- Tipos: AWS Public AMI, Community AMI, Custom AMI, Marketplace AMI

---

## Pegadinhas da Prova

1. **Spot Instances** podem ser terminadas com 2 minutos de aviso
2. **Reserved Instances não garantem capacidade** — use Capacity Reservations para isso
3. **EBS é por AZ** — para mover entre AZs, crie um snapshot
4. **Instance Store não persiste** ao parar/terminar a instância
5. **ASG não substitui instâncias imediatamente** — depende do health check
6. **ALB suporta WebSocket e HTTP/2**; NLB suporta TCP/UDP
7. **Hibernate preserva RAM** — instância pausa e retoma do mesmo estado

---

## Caso Real

### Netflix
Usa massivamente Spot Instances para encoding de vídeo. O processo pode ser interrompido e retomado sem perda, economizando >80% no custo de compute.

---

## Lab Relacionado

[Lab 03 — EC2: Lançar, conectar via SSH, EBS](../labs/lab-03-ec2-launch-ssh.md)
